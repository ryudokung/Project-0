package auth

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	GetByPrivyDID(did string) (*User, error)
	GetByWalletAddress(wallet string) (*User, error)
	Create(user *User) error
	UpdateLastLogin(id uuid.UUID) error
	UpdatePrivyDID(id uuid.UUID, did string) error
	UpdateWalletAddress(id uuid.UUID, wallet string) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetByPrivyDID(did string) (*User, error) {
	query := `SELECT id, privy_did, wallet_address, username, credits, last_login, created_at FROM users WHERE privy_did = $1`
	row := r.db.QueryRow(query, did)

	var user User
	var privyDID, walletAddress sql.NullString
	err := row.Scan(&user.ID, &privyDID, &walletAddress, &user.Username, &user.Credits, &user.LastLogin, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning user by did: %w", err)
	}

	user.PrivyDID = privyDID.String
	user.WalletAddress = walletAddress.String

	return &user, nil
}

func (r *postgresRepository) GetByWalletAddress(wallet string) (*User, error) {
	query := `SELECT id, privy_did, wallet_address, username, credits, last_login, created_at FROM users WHERE wallet_address = $1`
	row := r.db.QueryRow(query, wallet)

	var user User
	var privyDID, walletAddress sql.NullString
	err := row.Scan(&user.ID, &privyDID, &walletAddress, &user.Username, &user.Credits, &user.LastLogin, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning user: %w", err)
	}

	user.PrivyDID = privyDID.String
	user.WalletAddress = walletAddress.String

	return &user, nil
}

func (r *postgresRepository) Create(user *User) error {
	query := `INSERT INTO users (id, privy_did, wallet_address, username, credits) VALUES ($1, $2, $3, $4, $5) RETURNING created_at, last_login`
	err := r.db.QueryRow(query, user.ID, user.PrivyDID, user.WalletAddress, user.Username, user.Credits).Scan(&user.CreatedAt, &user.LastLogin)
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

func (r *postgresRepository) UpdatePrivyDID(id uuid.UUID, did string) error {
	query := `UPDATE users SET privy_did = $1 WHERE id = $2`
	_, err := r.db.Exec(query, did, id)
	return err
}

func (r *postgresRepository) UpdateWalletAddress(id uuid.UUID, wallet string) error {
	query := `UPDATE users SET wallet_address = $1 WHERE id = $2`
	_, err := r.db.Exec(query, wallet, id)
	return err
}
