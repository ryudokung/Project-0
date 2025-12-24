package auth

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `json:"id"`
	PrivyDID      string    `json:"privy_did"`
	WalletAddress string    `json:"wallet_address"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	PasswordHash  string    `json:"-"`
	GuestID       string    `json:"guest_id"`
	AuthType          string     `json:"auth_type"`
	Credits           float64    `json:"credits"`
	ActiveCharacterID *uuid.UUID `json:"active_character_id"`
	ActiveCharacter   *Character `json:"active_character,omitempty"`
	LastLogin         time.Time  `json:"last_login"`
	CreatedAt         time.Time  `json:"created_at"`
}

type Character struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	FaceIndex int       `json:"face_index"`
	HairIndex int       `json:"hair_index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCharacterRequest struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	FaceIndex int    `json:"face_index"`
	HairIndex int    `json:"hair_index"`
}

type LoginRequest struct {
	PrivyDID      string `json:"privy_did"`
	WalletAddress string `json:"wallet_address"`
	PrivyToken    string `json:"privy_token"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	GuestID       string `json:"guest_id"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	GuestID  string `json:"guest_id"` // Optional: to upgrade a guest account
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
