package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UseCase interface {
	Login(req LoginRequest) (*LoginResponse, error)
}

type authUseCase struct {
	repo      Repository
	jwtSecret string
}

func NewUseCase(repo Repository, jwtSecret string) UseCase {
	return &authUseCase{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (u *authUseCase) Login(req LoginRequest) (*LoginResponse, error) {
	// TODO: Implement actual Web3 signature verification here
	// For now, we assume the signature is valid if it's not empty
	if req.Signature == "" {
		return nil, fmt.Errorf("invalid signature")
	}

	user, err := u.repo.GetByWalletAddress(req.WalletAddress)
	if err != nil {
		return nil, err
	}

	if user == nil {
		// Create new user if not exists
		user = &User{
			ID:            uuid.New(),
			WalletAddress: req.WalletAddress,
			Username:      "Pilot-" + req.WalletAddress[2:6],
			Credits:       0,
		}
		if err := u.repo.Create(user); err != nil {
			return nil, err
		}
	} else {
		if err := u.repo.UpdateLastLogin(user.ID); err != nil {
			return nil, err
		}
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"wallet":  user.WalletAddress,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return nil, fmt.Errorf("error signing token: %w", err)
	}

	return &LoginResponse{
		Token: tokenString,
		User:  *user,
	}, nil
}
