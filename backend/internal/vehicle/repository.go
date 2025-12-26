package vehicle

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Repository interface {
	Create(vehicle *Vehicle) error
	GetByID(ctx context.Context, id uuid.UUID) (*Vehicle, error)
	GetByOwnerID(ownerID uuid.UUID) ([]Vehicle, error)
	GetByCharacterID(charID uuid.UUID) ([]Vehicle, error)
	Update(ctx context.Context, vehicle *Vehicle) error
	UpdateStatus(id uuid.UUID, status VehicleStatus, tokenID string) error
	UpdateHP(ctx context.Context, id uuid.UUID, newHP int) error

	// Part operations
	CreatePart(part *Part) error
	GetPartsByOwnerID(ownerID uuid.UUID) ([]Part, error)
	GetPartsByVehicleID(vehicleID uuid.UUID) ([]Part, error)
	EquipPart(partID uuid.UUID, vehicleID uuid.UUID) error
	UnequipPart(partID uuid.UUID) error

	// Item operations (DDS)
	CreateItem(ctx context.Context, item *Item) error
	GetItemByID(ctx context.Context, id uuid.UUID) (*Item, error)
	UpdateItem(ctx context.Context, item *Item) error
	GetItemsByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]Item, error)
	UpdateDurability(ctx context.Context, id uuid.UUID, durability int, condition ItemCondition) error
}

type vehicleRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) Create(v *Vehicle) error {
	statsJSON, err := json.Marshal(v.Stats)
	if err != nil {
		return err
	}
	metadataJSON, _ := json.Marshal(v.Metadata)

	query := `
		INSERT INTO vehicles (id, owner_id, character_id, vehicle_type, class, image_url, stats, cr, suitability_tags, rarity, tier, is_void_touched, season, status, metadata)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`
	_, err = r.db.Exec(query, v.ID, v.OwnerID, v.CharacterID, v.VehicleType, v.Class, v.ImageURL, statsJSON, v.CR, pq.Array(v.SuitabilityTags), v.Rarity, v.Tier, v.IsVoidTouched, v.Season, v.Status, metadataJSON)
	return err
}

func (r *vehicleRepository) GetByID(ctx context.Context, id uuid.UUID) (*Vehicle, error) {
	query := `SELECT id, token_id, owner_id, character_id, vehicle_type, class, image_url, stats, cr, suitability_tags, rarity, tier, is_void_touched, season, status, metadata, created_at FROM vehicles WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var v Vehicle
	var statsJSON, metadataJSON []byte
	var suitabilityTags []string
	var tokenID, imageURL, season sql.NullString
	var charID uuid.NullUUID

	err := row.Scan(&v.ID, &tokenID, &v.OwnerID, &charID, &v.VehicleType, &v.Class, &imageURL, &statsJSON, &v.CR, pq.Array(&suitabilityTags), &v.Rarity, &v.Tier, &v.IsVoidTouched, &season, &v.Status, &metadataJSON, &v.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if tokenID.Valid {
		v.TokenID = &tokenID.String
	}
	if imageURL.Valid {
		v.ImageURL = &imageURL.String
	}
	if season.Valid {
		v.Season = &season.String
	}
	if charID.Valid {
		v.CharacterID = &charID.UUID
	}
	v.SuitabilityTags = suitabilityTags

	if err := json.Unmarshal(statsJSON, &v.Stats); err != nil {
		return nil, err
	}
	_ = json.Unmarshal(metadataJSON, &v.Metadata)

	return &v, nil
}

func (r *vehicleRepository) Update(ctx context.Context, v *Vehicle) error {
	statsJSON, _ := json.Marshal(v.Stats)
	metadataJSON, _ := json.Marshal(v.Metadata)
	query := `
		UPDATE vehicles 
		SET token_id = $1, owner_id = $2, character_id = $3, vehicle_type = $4, class = $5, image_url = $6, stats = $7, cr = $8, suitability_tags = $9, rarity = $10, tier = $11, is_void_touched = $12, season = $13, status = $14, metadata = $15
		WHERE id = $16
	`
	_, err := r.db.ExecContext(ctx, query, v.TokenID, v.OwnerID, v.CharacterID, v.VehicleType, v.Class, v.ImageURL, statsJSON, v.CR, pq.Array(v.SuitabilityTags), v.Rarity, v.Tier, v.IsVoidTouched, v.Season, v.Status, metadataJSON, v.ID)
	return err
}

func (r *vehicleRepository) GetByOwnerID(ownerID uuid.UUID) ([]Vehicle, error) {
	query := `SELECT id, token_id, owner_id, character_id, vehicle_type, class, image_url, stats, cr, suitability_tags, rarity, tier, is_void_touched, season, status, metadata, created_at FROM vehicles WHERE owner_id = $1`
	rows, err := r.db.Query(query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vehicles := []Vehicle{}
	for rows.Next() {
		var v Vehicle
		var statsJSON, metadataJSON []byte
		var suitabilityTags []string
		var tokenID, imageURL, season sql.NullString
		var charID uuid.NullUUID
		if err := rows.Scan(&v.ID, &tokenID, &v.OwnerID, &charID, &v.VehicleType, &v.Class, &imageURL, &statsJSON, &v.CR, pq.Array(&suitabilityTags), &v.Rarity, &v.Tier, &v.IsVoidTouched, &season, &v.Status, &metadataJSON, &v.CreatedAt); err != nil {
			return nil, err
		}
		if tokenID.Valid {
			v.TokenID = &tokenID.String
		}
		if imageURL.Valid {
			v.ImageURL = &imageURL.String
		}
		if season.Valid {
			v.Season = &season.String
		}
		if charID.Valid {
			v.CharacterID = &charID.UUID
		}
		v.SuitabilityTags = suitabilityTags
		json.Unmarshal(statsJSON, &v.Stats)
		json.Unmarshal(metadataJSON, &v.Metadata)
		vehicles = append(vehicles, v)
	}
	return vehicles, nil
}

func (r *vehicleRepository) GetByCharacterID(charID uuid.UUID) ([]Vehicle, error) {
	query := `SELECT id, token_id, owner_id, character_id, vehicle_type, class, image_url, stats, cr, suitability_tags, rarity, tier, is_void_touched, season, status, metadata, created_at FROM vehicles WHERE character_id = $1`
	rows, err := r.db.Query(query, charID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vehicles := []Vehicle{}
	for rows.Next() {
		var v Vehicle
		var statsJSON, metadataJSON []byte
		var suitabilityTags []string
		var tokenID, imageURL, season sql.NullString
		var cID uuid.NullUUID
		if err := rows.Scan(&v.ID, &tokenID, &v.OwnerID, &cID, &v.VehicleType, &v.Class, &imageURL, &statsJSON, &v.CR, pq.Array(&suitabilityTags), &v.Rarity, &v.Tier, &v.IsVoidTouched, &season, &v.Status, &metadataJSON, &v.CreatedAt); err != nil {
			return nil, err
		}
		if tokenID.Valid {
			v.TokenID = &tokenID.String
		}
		if imageURL.Valid {
			v.ImageURL = &imageURL.String
		}
		if season.Valid {
			v.Season = &season.String
		}
		if cID.Valid {
			v.CharacterID = &cID.UUID
		}
		v.SuitabilityTags = suitabilityTags
		json.Unmarshal(statsJSON, &v.Stats)
		json.Unmarshal(metadataJSON, &v.Metadata)
		vehicles = append(vehicles, v)
	}
	return vehicles, nil
}

func (r *vehicleRepository) UpdateStatus(id uuid.UUID, status VehicleStatus, tokenID string) error {
	query := `UPDATE vehicles SET status = $1, token_id = $2 WHERE id = $3`
	_, err := r.db.Exec(query, status, tokenID, id)
	return err
}

func (r *vehicleRepository) UpdateHP(ctx context.Context, id uuid.UUID, newHP int) error {
	query := `UPDATE vehicles SET stats = jsonb_set(stats, '{hp}', $1::text::jsonb) WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, newHP, id)
	return err
}

func (r *vehicleRepository) CreatePart(p *Part) error {
	statsJSON, _ := json.Marshal(p.Stats)
	dnaJSON, _ := json.Marshal(p.VisualDNA)

	query := `
		INSERT INTO parts (id, owner_id, vehicle_id, slot, name, rarity, tier, is_void_touched, is_minted, stats, visual_dna)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.db.Exec(query, p.ID, p.OwnerID, p.VehicleID, p.Slot, p.Name, p.Rarity, p.Tier, p.IsVoidTouched, p.IsMinted, statsJSON, dnaJSON)
	return err
}

func (r *vehicleRepository) GetPartsByOwnerID(ownerID uuid.UUID) ([]Part, error) {
	query := `SELECT id, owner_id, vehicle_id, slot, name, rarity, tier, is_void_touched, is_minted, stats, visual_dna, created_at FROM parts WHERE owner_id = $1`
	rows, err := r.db.Query(query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanParts(rows)
}

func (r *vehicleRepository) GetPartsByVehicleID(vehicleID uuid.UUID) ([]Part, error) {
	query := `SELECT id, owner_id, vehicle_id, slot, name, rarity, tier, is_void_touched, is_minted, stats, visual_dna, created_at FROM parts WHERE vehicle_id = $1`
	rows, err := r.db.Query(query, vehicleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanParts(rows)
}

func (r *vehicleRepository) scanParts(rows *sql.Rows) ([]Part, error) {
	var parts []Part
	for rows.Next() {
		var p Part
		var statsJSON, dnaJSON []byte
		if err := rows.Scan(&p.ID, &p.OwnerID, &p.VehicleID, &p.Slot, &p.Name, &p.Rarity, &p.Tier, &p.IsVoidTouched, &p.IsMinted, &statsJSON, &dnaJSON, &p.CreatedAt); err != nil {
			return nil, err
		}
		json.Unmarshal(statsJSON, &p.Stats)
		json.Unmarshal(dnaJSON, &p.VisualDNA)
		parts = append(parts, p)
	}
	return parts, nil
}

func (r *vehicleRepository) EquipPart(partID uuid.UUID, vehicleID uuid.UUID) error {
	query := `UPDATE parts SET vehicle_id = $1 WHERE id = $2`
	_, err := r.db.Exec(query, vehicleID, partID)
	return err
}

func (r *vehicleRepository) UnequipPart(partID uuid.UUID) error {
	query := `UPDATE parts SET vehicle_id = NULL WHERE id = $1`
	_, err := r.db.Exec(query, partID)
	return err
}

// Item implementations (DDS)

func (r *vehicleRepository) CreateItem(ctx context.Context, i *Item) error {
	statsJSON, _ := json.Marshal(i.Stats)
	dnaJSON, _ := json.Marshal(i.VisualDNA)
	metaJSON, _ := json.Marshal(i.Metadata)

	query := `
		INSERT INTO items (
			id, owner_id, character_id, name, item_type, rarity, tier, slot, 
			damage_type, series_id,
			is_nft, token_id, durability, max_durability, condition, 
			stats, visual_dna, metadata, is_equipped, parent_item_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
	`
	_, err := r.db.ExecContext(ctx, query,
		i.ID, i.OwnerID, i.CharacterID, i.Name, i.ItemType, i.Rarity, i.Tier, i.Slot,
		i.DamageType, i.SeriesID,
		i.IsNFT, i.TokenID, i.Durability, i.MaxDurability, i.Condition,
		statsJSON, dnaJSON, metaJSON, i.IsEquipped, i.ParentItemID,
	)
	return err
}

func (r *vehicleRepository) GetItemByID(ctx context.Context, id uuid.UUID) (*Item, error) {
	query := `
		SELECT 
			id, owner_id, character_id, name, item_type, rarity, tier, slot, 
			damage_type, series_id,
			is_nft, token_id, durability, max_durability, condition, 
			stats, visual_dna, metadata, is_equipped, parent_item_id, created_at, updated_at
		FROM items WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var i Item
	var statsJSON, dnaJSON, metaJSON []byte
	var tokenID, damageType, seriesID sql.NullString
	var charID, parentID uuid.NullUUID

	err := row.Scan(
		&i.ID, &i.OwnerID, &charID, &i.Name, &i.ItemType, &i.Rarity, &i.Tier, &i.Slot,
		&damageType, &seriesID,
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
	if damageType.Valid {
		i.DamageType = &damageType.String
	}
	if seriesID.Valid {
		i.SeriesID = &seriesID.String
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

func (r *vehicleRepository) UpdateItem(ctx context.Context, i *Item) error {
	statsJSON, _ := json.Marshal(i.Stats)
	dnaJSON, _ := json.Marshal(i.VisualDNA)
	metaJSON, _ := json.Marshal(i.Metadata)

	query := `
		UPDATE items SET 
			owner_id = $1, character_id = $2, name = $3, item_type = $4, rarity = $5, 
			tier = $6, slot = $7, damage_type = $8, series_id = $9, is_nft = $10, token_id = $11, durability = $12, 
			max_durability = $13, condition = $14, stats = $15, visual_dna = $16, 
			metadata = $17, is_equipped = $18, parent_item_id = $19
		WHERE id = $20
	`
	_, err := r.db.ExecContext(ctx, query,
		i.OwnerID, i.CharacterID, i.Name, i.ItemType, i.Rarity,
		i.Tier, i.Slot, i.DamageType, i.SeriesID, i.IsNFT, i.TokenID, i.Durability,
		i.MaxDurability, i.Condition, statsJSON, dnaJSON,
		metaJSON, i.IsEquipped, i.ParentItemID, i.ID,
	)
	return err
}

func (r *vehicleRepository) GetItemsByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]Item, error) {
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

func (r *vehicleRepository) UpdateDurability(ctx context.Context, id uuid.UUID, durability int, condition ItemCondition) error {
	query := `UPDATE items SET durability = $1, condition = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, durability, condition, id)
	return err
}
