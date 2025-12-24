package gacha

import (
	"database/sql"
	"github.com/google/uuid"
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
		INSERT INTO gacha_history (id, user_id, item_id, item_type, pull_type, rarity)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(query, h.ID, h.UserID, h.ItemID, h.ItemType, h.PullType, h.Rarity)
	return err
}
