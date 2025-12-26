package game

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
)

type Repository interface {
	GetPilotStats(charID uuid.UUID) (*PilotStats, error)
	GetActivePilotStats(userID uuid.UUID) (*PilotStats, error)
	UpdatePilotStats(stats *PilotStats) error
	InitializePilot(charID uuid.UUID) error
	InitializeGachaStats(userID uuid.UUID) error
	GetGachaStats(userID uuid.UUID) (*GachaStats, error)
	UpdateGachaStats(stats *GachaStats) error
	ConsumeFreePull(userID uuid.UUID) (bool, error)
	ConsumeResources(charID uuid.UUID, o2, fuel float64) (bool, error)
}

type gameRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &gameRepository{db: db}
}

func (r *gameRepository) GetPilotStats(charID uuid.UUID) (*PilotStats, error) {
	query := `SELECT user_id, character_id, equipped_exosuit_id, resonance_level, resonance_exp, stress, xp, rank, current_o2, current_fuel, scrap_metal, research_data, metadata, updated_at FROM pilot_stats WHERE character_id = $1`
	row := r.db.QueryRow(query, charID)

	var s PilotStats
	var metadataJSON []byte
	err := row.Scan(&s.UserID, &s.CharacterID, &s.EquippedExosuitID, &s.ResonanceLevel, &s.ResonanceExp, &s.Stress, &s.XP, &s.Rank, &s.CurrentO2, &s.CurrentFuel, &s.ScrapMetal, &s.ResearchData, &metadataJSON, &s.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	json.Unmarshal(metadataJSON, &s.Metadata)
	return &s, err
}

func (r *gameRepository) GetActivePilotStats(userID uuid.UUID) (*PilotStats, error) {
	query := `
		SELECT ps.user_id, ps.character_id, ps.equipped_exosuit_id, ps.resonance_level, ps.resonance_exp, ps.stress, ps.xp, ps.rank, ps.current_o2, ps.current_fuel, ps.scrap_metal, ps.research_data, ps.metadata, ps.updated_at 
		FROM pilot_stats ps
		JOIN users u ON ps.character_id = u.active_character_id
		WHERE u.id = $1
	`
	row := r.db.QueryRow(query, userID)

	var s PilotStats
	var metadataJSON []byte
	err := row.Scan(&s.UserID, &s.CharacterID, &s.EquippedExosuitID, &s.ResonanceLevel, &s.ResonanceExp, &s.Stress, &s.XP, &s.Rank, &s.CurrentO2, &s.CurrentFuel, &s.ScrapMetal, &s.ResearchData, &metadataJSON, &s.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	json.Unmarshal(metadataJSON, &s.Metadata)
	return &s, err
}

func (r *gameRepository) UpdatePilotStats(s *PilotStats) error {
	metadataJSON, _ := json.Marshal(s.Metadata)
	query := `
		UPDATE pilot_stats 
		SET equipped_exosuit_id = $1, resonance_level = $2, resonance_exp = $3, stress = $4, xp = $5, rank = $6, current_o2 = $7, current_fuel = $8, scrap_metal = $9, research_data = $10, metadata = $11, updated_at = CURRENT_TIMESTAMP
		WHERE character_id = $12
	`
	_, err := r.db.Exec(query, s.EquippedExosuitID, s.ResonanceLevel, s.ResonanceExp, s.Stress, s.XP, s.Rank, s.CurrentO2, s.CurrentFuel, s.ScrapMetal, s.ResearchData, metadataJSON, s.CharacterID)
	return err
}

func (r *gameRepository) InitializePilot(charID uuid.UUID) error {
	// We need to find the user_id for this character_id
	var userID uuid.UUID
	err := r.db.QueryRow("SELECT user_id FROM characters WHERE id = $1", charID).Scan(&userID)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO pilot_stats (user_id, character_id, resonance_level, resonance_exp, stress, xp, rank, current_o2, current_fuel, scrap_metal, research_data, metadata)
		VALUES ($1, $2, 0, 0, 0, 0, 1, 100.0, 100.0, 0, 0, '{"radar_level": 1, "lab_level": 1, "warp_level": 1}')
		ON CONFLICT (character_id) DO NOTHING
	`
	_, err = r.db.Exec(query, userID, charID)
	return err
}

func (r *gameRepository) InitializeGachaStats(userID uuid.UUID) error {
	query := `
		INSERT INTO gacha_stats (user_id, pity_relic_count, pity_singularity_count, total_pulls)
		VALUES ($1, 0, 0, 0)
		ON CONFLICT (user_id) DO NOTHING
	`
	_, err := r.db.Exec(query, userID)
	return err
}

func (r *gameRepository) GetGachaStats(userID uuid.UUID) (*GachaStats, error) {
	query := `SELECT user_id, pity_relic_count, pity_singularity_count, total_pulls, last_free_pull_at, updated_at FROM gacha_stats WHERE user_id = $1`
	row := r.db.QueryRow(query, userID)

	var s GachaStats
	err := row.Scan(&s.UserID, &s.PityRelicCount, &s.PitySingularityCount, &s.TotalPulls, &s.LastFreePullAt, &s.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &s, err
}

func (r *gameRepository) UpdateGachaStats(s *GachaStats) error {
	query := `
		UPDATE gacha_stats 
		SET pity_relic_count = $1, pity_singularity_count = $2, total_pulls = $3, last_free_pull_at = $4, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $5
	`
	_, err := r.db.Exec(query, s.PityRelicCount, s.PitySingularityCount, s.TotalPulls, s.LastFreePullAt, s.UserID)
	return err
}

func (r *gameRepository) ConsumeFreePull(userID uuid.UUID) (bool, error) {
	query := `
		UPDATE gacha_stats 
		SET last_free_pull_at = CURRENT_TIMESTAMP 
		WHERE user_id = $1 AND (last_free_pull_at IS NULL OR last_free_pull_at < CURRENT_TIMESTAMP - INTERVAL '24 hours')
	`
	res, err := r.db.Exec(query, userID)
	if err != nil {
		return false, err
	}
	rows, _ := res.RowsAffected()
	return rows > 0, nil
}

func (r *gameRepository) ConsumeResources(charID uuid.UUID, o2, fuel float64) (bool, error) {
	query := `
		UPDATE pilot_stats 
		SET current_o2 = current_o2 - $1, current_fuel = current_fuel - $2 
		WHERE character_id = $3 AND current_o2 >= $1 AND current_fuel >= $2
	`
	res, err := r.db.Exec(query, o2, fuel, charID)
	if err != nil {
		return false, err
	}
	rows, _ := res.RowsAffected()
	return rows > 0, nil
}
