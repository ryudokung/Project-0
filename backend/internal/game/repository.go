package game

import (
	"database/sql"
	"github.com/google/uuid"
)

type Repository interface {
	GetPilotStats(userID uuid.UUID) (*PilotStats, error)
	UpdatePilotStats(stats *PilotStats) error
	InitializePilot(userID uuid.UUID) error
	GetGachaStats(userID uuid.UUID) (*GachaStats, error)
	UpdateGachaStats(stats *GachaStats) error
}

type gameRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &gameRepository{db: db}
}

func (r *gameRepository) GetPilotStats(userID uuid.UUID) (*PilotStats, error) {
	query := `SELECT user_id, resonance_level, resonance_exp, xp, rank, current_o2, current_fuel, updated_at FROM pilot_stats WHERE user_id = $1`
	row := r.db.QueryRow(query, userID)

	var s PilotStats
	err := row.Scan(&s.UserID, &s.ResonanceLevel, &s.ResonanceExp, &s.XP, &s.Rank, &s.CurrentO2, &s.CurrentFuel, &s.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &s, err
}

func (r *gameRepository) UpdatePilotStats(s *PilotStats) error {
	query := `
		UPDATE pilot_stats 
		SET resonance_level = $1, resonance_exp = $2, xp = $3, rank = $4, current_o2 = $5, current_fuel = $6, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $7
	`
	_, err := r.db.Exec(query, s.ResonanceLevel, s.ResonanceExp, s.XP, s.Rank, s.CurrentO2, s.CurrentFuel, s.UserID)
	return err
}

func (r *gameRepository) InitializePilot(userID uuid.UUID) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Initialize Pilot Stats
	_, err = tx.Exec(`INSERT INTO pilot_stats (user_id) VALUES ($1) ON CONFLICT (user_id) DO NOTHING`, userID)
	if err != nil {
		return err
	}

	// Initialize Gacha Stats
	_, err = tx.Exec(`INSERT INTO gacha_stats (user_id) VALUES ($1) ON CONFLICT (user_id) DO NOTHING`, userID)
	if err != nil {
		return err
	}

	return tx.Commit()
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
