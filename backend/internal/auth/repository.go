package auth

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	GetByWalletAddress(wallet string) (*User, error)
	Create(user *User) error
	UpdateLastLogin(id uuid.UUID) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetByWalletAddress(wallet string) (*User, error) {
	query := `SELECT id, wallet_address, username, credits, last_login, created_at FROM users WHERE wallet_address = $1`
	row := r.db.QueryRow(query, wallet)

	var user User
	err := row.Scan(&user.ID, &user.WalletAddress, &user.Username, &user.Credits, &user.LastLogin, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning user: %w", err)
	}

	return &user, nil
}

func (r *postgresRepository) Create(user *User) error {
	query := `INSERT INTO users (id, wallet_address, username, credits) VALUES ($1, $2, $3, $4) RETURNING created_at, last_login`
	err := r.db.QueryRow(query, user.ID, user.WalletAddress, user.Username, user.Credits).Scan(&user.CreatedAt, &user.LastLogin)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (r *postgresRepository) UpdateLastLogin(id uuid.UUID) error {
	query := `UPDATE users SET last_login = CURRENT_TIMESTAMP WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
