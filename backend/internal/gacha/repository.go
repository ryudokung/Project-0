package gacha

import (
	"database/sql"
)

type Repository interface {
	SaveHistory(history *GachaHistory) error
}

type gachaRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &gachaRepository{db: db}
}

func (r *gachaRepository) SaveHistory(h *GachaHistory) error {
	query := `
		INSERT INTO gacha_history (
			id, user_id, item_id, item_type, pull_type, rarity, 
			seed, pity_relic_before, pity_relic_after, 
			pity_singularity_before, pity_singularity_after
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.db.Exec(
		query, 
		h.ID, h.UserID, h.ItemID, h.ItemType, h.PullType, h.Rarity,
		h.Seed, h.PityRelicBefore, h.PityRelicAfter,
		h.PitySingularityBefore, h.PitySingularityAfter,
	)
	return err
}
