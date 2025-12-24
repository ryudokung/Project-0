package mech

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
)

type Repository interface {
	Create(mech *Mech) error
	GetByID(ctx context.Context, id uuid.UUID) (*Mech, error)
	GetByOwnerID(ownerID uuid.UUID) ([]Mech, error)
	Update(ctx context.Context, mech *Mech) error
	UpdateStatus(id uuid.UUID, status MechStatus, tokenID string) error

	// Part operations
	CreatePart(part *Part) error
	GetPartsByOwnerID(ownerID uuid.UUID) ([]Part, error)
	GetPartsByMechID(mechID uuid.UUID) ([]Part, error)
	EquipPart(partID uuid.UUID, mechID uuid.UUID) error
	UnequipPart(partID uuid.UUID) error
}

type mechRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &mechRepository{db: db}
}

func (r *mechRepository) Create(m *Mech) error {
	statsJSON, err := json.Marshal(m.Stats)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO mechs (id, owner_id, vehicle_type, class, image_url, stats, rarity, tier, is_void_touched, season, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err = r.db.Exec(query, m.ID, m.OwnerID, m.VehicleType, m.Class, m.ImageURL, statsJSON, m.Rarity, m.Tier, m.IsVoidTouched, m.Season, m.Status)
	return err
}

func (r *mechRepository) GetByID(ctx context.Context, id uuid.UUID) (*Mech, error) {
	query := `SELECT id, token_id, owner_id, vehicle_type, class, image_url, stats, rarity, tier, is_void_touched, season, status, created_at FROM mechs WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var m Mech
	var statsJSON []byte
	var tokenID, imageURL, season sql.NullString

	err := row.Scan(&m.ID, &tokenID, &m.OwnerID, &m.VehicleType, &m.Class, &imageURL, &statsJSON, &m.Rarity, &m.Tier, &m.IsVoidTouched, &season, &m.Status, &m.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if tokenID.Valid {
		m.TokenID = &tokenID.String
	}
	if imageURL.Valid {
		m.ImageURL = &imageURL.String
	}
	if season.Valid {
		m.Season = &season.String
	}

	if err := json.Unmarshal(statsJSON, &m.Stats); err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *mechRepository) Update(ctx context.Context, m *Mech) error {
	statsJSON, _ := json.Marshal(m.Stats)
	query := `
		UPDATE mechs 
		SET token_id = $1, owner_id = $2, vehicle_type = $3, class = $4, image_url = $5, stats = $6, rarity = $7, tier = $8, is_void_touched = $9, season = $10, status = $11
		WHERE id = $12
	`
	_, err := r.db.ExecContext(ctx, query, m.TokenID, m.OwnerID, m.VehicleType, m.Class, m.ImageURL, statsJSON, m.Rarity, m.Tier, m.IsVoidTouched, m.Season, m.Status, m.ID)
	return err
}

func (r *mechRepository) GetByOwnerID(ownerID uuid.UUID) ([]Mech, error) {
	query := `SELECT id, token_id, owner_id, vehicle_type, class, image_url, stats, rarity, tier, is_void_touched, season, status, created_at FROM mechs WHERE owner_id = $1`
	rows, err := r.db.Query(query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mechs []Mech
	for rows.Next() {
		var m Mech
		var statsJSON []byte
		var tokenID, imageURL, season sql.NullString
		if err := rows.Scan(&m.ID, &tokenID, &m.OwnerID, &m.VehicleType, &m.Class, &imageURL, &statsJSON, &m.Rarity, &m.Tier, &m.IsVoidTouched, &season, &m.Status, &m.CreatedAt); err != nil {
			return nil, err
		}
		if tokenID.Valid {
			m.TokenID = &tokenID.String
		}
		if imageURL.Valid {
			m.ImageURL = &imageURL.String
		}
		if season.Valid {
			m.Season = &season.String
		}
		json.Unmarshal(statsJSON, &m.Stats)
		mechs = append(mechs, m)
	}
	return mechs, nil
}

func (r *mechRepository) UpdateStatus(id uuid.UUID, status MechStatus, tokenID string) error {
	query := `UPDATE mechs SET status = $1, token_id = $2 WHERE id = $3`
	_, err := r.db.Exec(query, status, tokenID, id)
	return err
}

func (r *mechRepository) CreatePart(p *Part) error {
	statsJSON, _ := json.Marshal(p.Stats)
	dnaJSON, _ := json.Marshal(p.VisualDNA)

	query := `
		INSERT INTO parts (id, owner_id, mech_id, slot, name, rarity, tier, is_void_touched, is_minted, stats, visual_dna)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.db.Exec(query, p.ID, p.OwnerID, p.MechID, p.Slot, p.Name, p.Rarity, p.Tier, p.IsVoidTouched, p.IsMinted, statsJSON, dnaJSON)
	return err
}

func (r *mechRepository) GetPartsByOwnerID(ownerID uuid.UUID) ([]Part, error) {
	query := `SELECT id, owner_id, mech_id, slot, name, rarity, tier, is_void_touched, is_minted, stats, visual_dna, created_at FROM parts WHERE owner_id = $1`
	rows, err := r.db.Query(query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanParts(rows)
}

func (r *mechRepository) GetPartsByMechID(mechID uuid.UUID) ([]Part, error) {
	query := `SELECT id, owner_id, mech_id, slot, name, rarity, tier, is_void_touched, is_minted, stats, visual_dna, created_at FROM parts WHERE mech_id = $1`
	rows, err := r.db.Query(query, mechID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanParts(rows)
}

func (r *mechRepository) scanParts(rows *sql.Rows) ([]Part, error) {
	var parts []Part
	for rows.Next() {
		var p Part
		var statsJSON, dnaJSON []byte
		if err := rows.Scan(&p.ID, &p.OwnerID, &p.MechID, &p.Slot, &p.Name, &p.Rarity, &p.Tier, &p.IsVoidTouched, &p.IsMinted, &statsJSON, &dnaJSON, &p.CreatedAt); err != nil {
			return nil, err
		}
		json.Unmarshal(statsJSON, &p.Stats)
		json.Unmarshal(dnaJSON, &p.VisualDNA)
		parts = append(parts, p)
	}
	return parts, nil
}

func (r *mechRepository) EquipPart(partID uuid.UUID, mechID uuid.UUID) error {
	query := `UPDATE parts SET mech_id = $1 WHERE id = $2`
	_, err := r.db.Exec(query, mechID, partID)
	return err
}

func (r *mechRepository) UnequipPart(partID uuid.UUID) error {
	query := `UPDATE parts SET mech_id = NULL WHERE id = $1`
	_, err := r.db.Exec(query, partID)
	return err
}
