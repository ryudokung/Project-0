package game

import (
	"database/sql"
	"github.com/google/uuid"
)

type Repository interface {
	GetPilotStats(userID uuid.UUID) (*PilotStats, error)
	UpdatePilotStats(stats *PilotStats) error
	InitializePilot(userID uuid.UUID) error
}

type gameRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &gameRepository{db: db}
}

func (r *gameRepository) GetPilotStats(userID uuid.UUID) (*PilotStats, error) {
	query := `SELECT user_id, resonance_level, resonance_exp, current_o2, current_fuel, updated_at FROM pilot_stats WHERE user_id = $1`
	row := r.db.QueryRow(query, userID)

	var s PilotStats
	err := row.Scan(&s.UserID, &s.ResonanceLevel, &s.ResonanceExp, &s.CurrentO2, &s.CurrentFuel, &s.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &s, err
}

func (r *gameRepository) UpdatePilotStats(s *PilotStats) error {
	query := `
		UPDATE pilot_stats 
		SET resonance_level = $1, resonance_exp = $2, current_o2 = $3, current_fuel = $4, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $5
	`
	_, err := r.db.Exec(query, s.ResonanceLevel, s.ResonanceExp, s.CurrentO2, s.CurrentFuel, s.UserID)
	return err
}

func (r *gameRepository) InitializePilot(userID uuid.UUID) error {
	query := `INSERT INTO pilot_stats (user_id) VALUES ($1) ON CONFLICT (user_id) DO NOTHING`
	_, err := r.db.Exec(query, userID)
	return err
}
