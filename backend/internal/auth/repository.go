package auth

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	GetByPrivyDID(did string) (*User, error)
	GetByWalletAddress(wallet string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByGuestID(guestID string) (*User, error)
	GetByID(id uuid.UUID) (*User, error)
	Create(user *User) error
	UpdateLastLogin(id uuid.UUID) error
	UpdatePrivyDID(id uuid.UUID, did string) error
	UpdateWalletAddress(id uuid.UUID, wallet string) error
	UpdateToTraditional(id uuid.UUID, username, email, passwordHash string) error

	// Character methods
	CreateCharacter(char *Character) error
	SetActiveCharacter(userID, charID uuid.UUID) error
	GetCharactersByUserID(userID uuid.UUID) ([]Character, error)
	GetCharacterByID(id uuid.UUID) (*Character, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetByPrivyDID(did string) (*User, error) {
	query := `SELECT id, privy_did, wallet_address, username, email, password_hash, guest_id, auth_type, credits, active_character_id, last_login, created_at FROM users WHERE privy_did = $1`
	row := r.db.QueryRow(query, did)

	var user User
	var privyDID, walletAddress, email, passwordHash, guestID, authType sql.NullString
	var activeCharID uuid.NullUUID
	err := row.Scan(&user.ID, &privyDID, &walletAddress, &user.Username, &email, &passwordHash, &guestID, &authType, &user.Credits, &activeCharID, &user.LastLogin, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning user by did: %w", err)
	}

	user.PrivyDID = privyDID.String
	user.WalletAddress = walletAddress.String
	user.Email = email.String
	user.PasswordHash = passwordHash.String
	user.GuestID = guestID.String
	user.AuthType = authType.String
	if activeCharID.Valid {
		user.ActiveCharacterID = &activeCharID.UUID
	}

	return &user, nil
}

func (r *postgresRepository) GetByWalletAddress(wallet string) (*User, error) {
	query := `SELECT id, privy_did, wallet_address, username, email, password_hash, guest_id, auth_type, credits, active_character_id, last_login, created_at FROM users WHERE wallet_address = $1`
	row := r.db.QueryRow(query, wallet)

	var user User
	var privyDID, walletAddress, email, passwordHash, guestID, authType sql.NullString
	var activeCharID uuid.NullUUID
	err := row.Scan(&user.ID, &privyDID, &walletAddress, &user.Username, &email, &passwordHash, &guestID, &authType, &user.Credits, &activeCharID, &user.LastLogin, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning user by wallet: %w", err)
	}

	user.PrivyDID = privyDID.String
	user.WalletAddress = walletAddress.String
	user.Email = email.String
	user.PasswordHash = passwordHash.String
	user.GuestID = guestID.String
	user.AuthType = authType.String
	if activeCharID.Valid {
		user.ActiveCharacterID = &activeCharID.UUID
	}

	return &user, nil
}

func (r *postgresRepository) GetByUsername(username string) (*User, error) {
	query := `SELECT id, privy_did, wallet_address, username, email, password_hash, guest_id, auth_type, credits, active_character_id, last_login, created_at FROM users WHERE username = $1`
	row := r.db.QueryRow(query, username)

	var user User
	var privyDID, walletAddress, email, passwordHash, guestID, authType sql.NullString
	var activeCharID uuid.NullUUID
	err := row.Scan(&user.ID, &privyDID, &walletAddress, &user.Username, &email, &passwordHash, &guestID, &authType, &user.Credits, &activeCharID, &user.LastLogin, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning user by username: %w", err)
	}

	user.PrivyDID = privyDID.String
	user.WalletAddress = walletAddress.String
	user.Email = email.String
	user.PasswordHash = passwordHash.String
	user.GuestID = guestID.String
	user.AuthType = authType.String
	if activeCharID.Valid {
		user.ActiveCharacterID = &activeCharID.UUID
	}

	return &user, nil
}

func (r *postgresRepository) GetByGuestID(guestID string) (*User, error) {
	query := `SELECT id, privy_did, wallet_address, username, email, password_hash, guest_id, auth_type, credits, active_character_id, last_login, created_at FROM users WHERE guest_id = $1`
	row := r.db.QueryRow(query, guestID)

	var user User
	var privyDID, walletAddress, email, passwordHash, gID, authType sql.NullString
	var activeCharID uuid.NullUUID
	err := row.Scan(&user.ID, &privyDID, &walletAddress, &user.Username, &email, &passwordHash, &gID, &authType, &user.Credits, &activeCharID, &user.LastLogin, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning user by guest_id: %w", err)
	}

	user.PrivyDID = privyDID.String
	user.WalletAddress = walletAddress.String
	user.Email = email.String
	user.PasswordHash = passwordHash.String
	user.GuestID = gID.String
	user.AuthType = authType.String
	if activeCharID.Valid {
		user.ActiveCharacterID = &activeCharID.UUID
	}

	return &user, nil
}

func (r *postgresRepository) GetByID(id uuid.UUID) (*User, error) {
	query := `SELECT id, privy_did, wallet_address, username, email, password_hash, guest_id, auth_type, credits, active_character_id, last_login, created_at FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var user User
	var privyDID, walletAddress, email, passwordHash, guestID, authType sql.NullString
	var activeCharID uuid.NullUUID
	err := row.Scan(&user.ID, &privyDID, &walletAddress, &user.Username, &email, &passwordHash, &guestID, &authType, &user.Credits, &activeCharID, &user.LastLogin, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning user by id: %w", err)
	}

	user.PrivyDID = privyDID.String
	user.WalletAddress = walletAddress.String
	user.Email = email.String
	user.PasswordHash = passwordHash.String
	user.GuestID = guestID.String
	user.AuthType = authType.String
	if activeCharID.Valid {
		user.ActiveCharacterID = &activeCharID.UUID
	}

	return &user, nil
}

func (r *postgresRepository) Create(user *User) error {
	query := `INSERT INTO users (id, privy_did, wallet_address, username, email, password_hash, guest_id, auth_type, credits) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING created_at, last_login`
	err := r.db.QueryRow(
		query,
		user.ID,
		toNullString(user.PrivyDID),
		toNullString(user.WalletAddress),
		user.Username,
		toNullString(user.Email),
		toNullString(user.PasswordHash),
		toNullString(user.GuestID),
		user.AuthType,
		user.Credits,
	).Scan(&user.CreatedAt, &user.LastLogin)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func (r *postgresRepository) UpdateLastLogin(id uuid.UUID) error {
	query := `UPDATE users SET last_login = CURRENT_TIMESTAMP WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *postgresRepository) UpdatePrivyDID(id uuid.UUID, did string) error {
	query := `UPDATE users SET privy_did = $1, auth_type = 'SOCIAL' WHERE id = $2`
	_, err := r.db.Exec(query, toNullString(did), id)
	return err
}

func (r *postgresRepository) UpdateWalletAddress(id uuid.UUID, wallet string) error {
	query := `UPDATE users SET wallet_address = $1 WHERE id = $2`
	_, err := r.db.Exec(query, toNullString(wallet), id)
	return err
}

func (r *postgresRepository) UpdateToTraditional(id uuid.UUID, username, email, passwordHash string) error {
	query := `UPDATE users SET username = $1, email = $2, password_hash = $3, auth_type = 'TRADITIONAL' WHERE id = $4`
	_, err := r.db.Exec(query, username, toNullString(email), toNullString(passwordHash), id)
	return err
}

func (r *postgresRepository) CreateCharacter(char *Character) error {
	query := `INSERT INTO characters (id, user_id, name, gender, face_index, hair_index) VALUES ($1, $2, $3, $4, $5, $6) RETURNING created_at, updated_at`
	return r.db.QueryRow(query, char.ID, char.UserID, char.Name, char.Gender, char.FaceIndex, char.HairIndex).Scan(&char.CreatedAt, &char.UpdatedAt)
}

func (r *postgresRepository) SetActiveCharacter(userID, charID uuid.UUID) error {
	query := `UPDATE users SET active_character_id = $1 WHERE id = $2`
	_, err := r.db.Exec(query, charID, userID)
	return err
}

func (r *postgresRepository) GetCharactersByUserID(userID uuid.UUID) ([]Character, error) {
	query := `SELECT id, user_id, name, gender, face_index, hair_index, created_at, updated_at FROM characters WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chars []Character
	for rows.Next() {
		var char Character
		if err := rows.Scan(&char.ID, &char.UserID, &char.Name, &char.Gender, &char.FaceIndex, &char.HairIndex, &char.CreatedAt, &char.UpdatedAt); err != nil {
			return nil, err
		}
		chars = append(chars, char)
	}
	return chars, nil
}

func (r *postgresRepository) GetCharacterByID(id uuid.UUID) (*Character, error) {
	query := `SELECT id, user_id, name, gender, face_index, hair_index, created_at, updated_at FROM characters WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var char Character
	if err := row.Scan(&char.ID, &char.UserID, &char.Name, &char.Gender, &char.FaceIndex, &char.HairIndex, &char.CreatedAt, &char.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &char, nil
}
