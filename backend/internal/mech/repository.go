package mech

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	Create(mech *Mech) error
	GetByID(id uuid.UUID) (*Mech, error)
	GetByOwnerID(ownerID uuid.UUID) ([]Mech, error)
	UpdateStatus(id uuid.UUID, status MechStatus, tokenID string) error
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
		INSERT INTO mechs (id, owner_id, vehicle_type, class, image_url, stats, rarity, season, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err = r.db.Exec(query, m.ID, m.OwnerID, m.VehicleType, m.Class, m.ImageURL, statsJSON, m.Rarity, m.Season, m.Status)
	return err
}

func (r *mechRepository) GetByID(id uuid.UUID) (*Mech, error) {
	query := `SELECT id, token_id, owner_id, vehicle_type, class, image_url, stats, rarity, season, status, created_at FROM mechs WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var m Mech
	var statsJSON []byte
	var tokenID sql.NullString

	err := row.Scan(&m.ID, &tokenID, &m.OwnerID, &m.VehicleType, &m.Class, &m.ImageURL, &statsJSON, &m.Rarity, &m.Season, &m.Status, &m.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if tokenID.Valid {
		m.TokenID = &tokenID.String
	}

	if err := json.Unmarshal(statsJSON, &m.Stats); err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *mechRepository) GetByOwnerID(ownerID uuid.UUID) ([]Mech, error) {
	query := `SELECT id, token_id, owner_id, vehicle_type, class, image_url, stats, rarity, season, status, created_at FROM mechs WHERE owner_id = $1`
	rows, err := r.db.Query(query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mechs []Mech
	for rows.Next() {
		var m Mech
		var statsJSON []byte
		var tokenID sql.NullString
		if err := rows.Scan(&m.ID, &tokenID, &m.OwnerID, &m.VehicleType, &m.Class, &m.ImageURL, &statsJSON, &m.Rarity, &m.Season, &m.Status, &m.CreatedAt); err != nil {
			return nil, err
		}
		if tokenID.Valid {
			m.TokenID = &tokenID.String
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
