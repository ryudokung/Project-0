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
	GetByCharacterID(charID uuid.UUID) ([]Mech, error)
	Update(ctx context.Context, mech *Mech) error
	UpdateStatus(id uuid.UUID, status MechStatus, tokenID string) error
	UpdateHP(ctx context.Context, id uuid.UUID, newHP int) error

	// Part operations
	CreatePart(part *Part) error
	GetPartsByOwnerID(ownerID uuid.UUID) ([]Part, error)
	GetPartsByMechID(mechID uuid.UUID) ([]Part, error)
	EquipPart(partID uuid.UUID, mechID uuid.UUID) error
	UnequipPart(partID uuid.UUID) error

	// Item operations (DDS)
	CreateItem(ctx context.Context, item *Item) error
	GetItemByID(ctx context.Context, id uuid.UUID) (*Item, error)
	UpdateItem(ctx context.Context, item *Item) error
	GetItemsByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]Item, error)
	UpdateDurability(ctx context.Context, id uuid.UUID, durability int, condition ItemCondition) error
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
		INSERT INTO mechs (id, owner_id, character_id, vehicle_type, class, image_url, stats, rarity, tier, is_void_touched, season, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	_, err = r.db.Exec(query, m.ID, m.OwnerID, m.CharacterID, m.VehicleType, m.Class, m.ImageURL, statsJSON, m.Rarity, m.Tier, m.IsVoidTouched, m.Season, m.Status)
	return err
}

func (r *mechRepository) GetByID(ctx context.Context, id uuid.UUID) (*Mech, error) {
	query := `SELECT id, token_id, owner_id, character_id, vehicle_type, class, image_url, stats, rarity, tier, is_void_touched, season, status, created_at FROM mechs WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var m Mech
	var statsJSON []byte
	var tokenID, imageURL, season sql.NullString
	var charID uuid.NullUUID

	err := row.Scan(&m.ID, &tokenID, &m.OwnerID, &charID, &m.VehicleType, &m.Class, &imageURL, &statsJSON, &m.Rarity, &m.Tier, &m.IsVoidTouched, &season, &m.Status, &m.CreatedAt)
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
	if charID.Valid {
		m.CharacterID = &charID.UUID
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
		SET token_id = $1, owner_id = $2, character_id = $3, vehicle_type = $4, class = $5, image_url = $6, stats = $7, rarity = $8, tier = $9, is_void_touched = $10, season = $11, status = $12
		WHERE id = $13
	`
	_, err := r.db.ExecContext(ctx, query, m.TokenID, m.OwnerID, m.CharacterID, m.VehicleType, m.Class, m.ImageURL, statsJSON, m.Rarity, m.Tier, m.IsVoidTouched, m.Season, m.Status, m.ID)
	return err
}

func (r *mechRepository) GetByOwnerID(ownerID uuid.UUID) ([]Mech, error) {
	query := `SELECT id, token_id, owner_id, character_id, vehicle_type, class, image_url, stats, rarity, tier, is_void_touched, season, status, created_at FROM mechs WHERE owner_id = $1`
	rows, err := r.db.Query(query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mechs := []Mech{}
	for rows.Next() {
		var m Mech
		var statsJSON []byte
		var tokenID, imageURL, season sql.NullString
		var charID uuid.NullUUID
		if err := rows.Scan(&m.ID, &tokenID, &m.OwnerID, &charID, &m.VehicleType, &m.Class, &imageURL, &statsJSON, &m.Rarity, &m.Tier, &m.IsVoidTouched, &season, &m.Status, &m.CreatedAt); err != nil {
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
		if charID.Valid {
			m.CharacterID = &charID.UUID
		}
		json.Unmarshal(statsJSON, &m.Stats)
		mechs = append(mechs, m)
	}
	return mechs, nil
}

func (r *mechRepository) GetByCharacterID(charID uuid.UUID) ([]Mech, error) {
	query := `SELECT id, token_id, owner_id, character_id, vehicle_type, class, image_url, stats, rarity, tier, is_void_touched, season, status, created_at FROM mechs WHERE character_id = $1`
	rows, err := r.db.Query(query, charID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mechs := []Mech{}
	for rows.Next() {
		var m Mech
		var statsJSON []byte
		var tokenID, imageURL, season sql.NullString
		var cID uuid.NullUUID
		if err := rows.Scan(&m.ID, &tokenID, &m.OwnerID, &cID, &m.VehicleType, &m.Class, &imageURL, &statsJSON, &m.Rarity, &m.Tier, &m.IsVoidTouched, &season, &m.Status, &m.CreatedAt); err != nil {
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
		if cID.Valid {
			m.CharacterID = &cID.UUID
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

func (r *mechRepository) UpdateHP(ctx context.Context, id uuid.UUID, newHP int) error {
	query := `UPDATE mechs SET stats = jsonb_set(stats, '{hp}', $1::text::jsonb) WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, newHP, id)
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

// Item implementations (DDS)

func (r *mechRepository) CreateItem(ctx context.Context, i *Item) error {
	statsJSON, _ := json.Marshal(i.Stats)
	dnaJSON, _ := json.Marshal(i.VisualDNA)
	metaJSON, _ := json.Marshal(i.Metadata)

	query := `
		INSERT INTO items (
			id, owner_id, character_id, name, item_type, rarity, tier, slot, 
			is_nft, token_id, durability, max_durability, condition, 
			stats, visual_dna, metadata, is_equipped, parent_item_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
	`
	_, err := r.db.ExecContext(ctx, query,
		i.ID, i.OwnerID, i.CharacterID, i.Name, i.ItemType, i.Rarity, i.Tier, i.Slot,
		i.IsNFT, i.TokenID, i.Durability, i.MaxDurability, i.Condition,
		statsJSON, dnaJSON, metaJSON, i.IsEquipped, i.ParentItemID,
	)
	return err
}

func (r *mechRepository) GetItemByID(ctx context.Context, id uuid.UUID) (*Item, error) {
	query := `
		SELECT 
			id, owner_id, character_id, name, item_type, rarity, tier, slot, 
			is_nft, token_id, durability, max_durability, condition, 
			stats, visual_dna, metadata, is_equipped, parent_item_id, created_at, updated_at
		FROM items WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var i Item
	var statsJSON, dnaJSON, metaJSON []byte
	var tokenID sql.NullString
	var charID, parentID uuid.NullUUID

	err := row.Scan(
		&i.ID, &i.OwnerID, &charID, &i.Name, &i.ItemType, &i.Rarity, &i.Tier, &i.Slot,
		&i.IsNFT, &tokenID, &i.Durability, &i.MaxDurability, &i.Condition,
		&statsJSON, &dnaJSON, &metaJSON, &i.IsEquipped, &parentID, &i.CreatedAt, &i.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if tokenID.Valid {
		i.TokenID = &tokenID.String
	}
	if charID.Valid {
		i.CharacterID = &charID.UUID
	}
	if parentID.Valid {
		i.ParentItemID = &parentID.UUID
	}

	json.Unmarshal(statsJSON, &i.Stats)
	json.Unmarshal(dnaJSON, &i.VisualDNA)
	json.Unmarshal(metaJSON, &i.Metadata)

	return &i, nil
}

func (r *mechRepository) UpdateItem(ctx context.Context, i *Item) error {
	statsJSON, _ := json.Marshal(i.Stats)
	dnaJSON, _ := json.Marshal(i.VisualDNA)
	metaJSON, _ := json.Marshal(i.Metadata)

	query := `
		UPDATE items SET 
			owner_id = $1, character_id = $2, name = $3, item_type = $4, rarity = $5, 
			tier = $6, slot = $7, is_nft = $8, token_id = $9, durability = $10, 
			max_durability = $11, condition = $12, stats = $13, visual_dna = $14, 
			metadata = $15, is_equipped = $16, parent_item_id = $17
		WHERE id = $18
	`
	_, err := r.db.ExecContext(ctx, query,
		i.OwnerID, i.CharacterID, i.Name, i.ItemType, i.Rarity,
		i.Tier, i.Slot, i.IsNFT, i.TokenID, i.Durability,
		i.MaxDurability, i.Condition, statsJSON, dnaJSON,
		metaJSON, i.IsEquipped, i.ParentItemID, i.ID,
	)
	return err
}

func (r *mechRepository) GetItemsByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]Item, error) {
	query := `
		SELECT 
			id, owner_id, character_id, name, item_type, rarity, tier, slot, 
			is_nft, token_id, durability, max_durability, condition, 
			stats, visual_dna, metadata, is_equipped, parent_item_id, created_at, updated_at
		FROM items WHERE owner_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var i Item
		var statsJSON, dnaJSON, metaJSON []byte
		var tokenID sql.NullString
		var charID, parentID uuid.NullUUID

		err := rows.Scan(
			&i.ID, &i.OwnerID, &charID, &i.Name, &i.ItemType, &i.Rarity, &i.Tier, &i.Slot,
			&i.IsNFT, &tokenID, &i.Durability, &i.MaxDurability, &i.Condition,
			&statsJSON, &dnaJSON, &metaJSON, &i.IsEquipped, &parentID, &i.CreatedAt, &i.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if tokenID.Valid {
			i.TokenID = &tokenID.String
		}
		if charID.Valid {
			i.CharacterID = &charID.UUID
		}
		if parentID.Valid {
			i.ParentItemID = &parentID.UUID
		}

		json.Unmarshal(statsJSON, &i.Stats)
		json.Unmarshal(dnaJSON, &i.VisualDNA)
		json.Unmarshal(metaJSON, &i.Metadata)
		items = append(items, i)
	}
	return items, nil
}

func (r *mechRepository) UpdateDurability(ctx context.Context, id uuid.UUID, durability int, condition ItemCondition) error {
	query := `UPDATE items SET durability = $1, condition = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, durability, condition, id)
	return err
}
