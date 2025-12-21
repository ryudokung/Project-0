package auth

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `json:"id"`
	WalletAddress string    `json:"wallet_address"`
	Username      string    `json:"username"`
	Credits       float64   `json:"credits"`
	LastLogin     time.Time `json:"last_login"`
	CreatedAt     time.Time `json:"created_at"`
}

type LoginRequest struct {
	WalletAddress string `json:"wallet_address" validate:"required"`
	Signature     string `json:"signature" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
