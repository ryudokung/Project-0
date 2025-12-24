package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/game"
)

type UseCase interface {
	Login(req LoginRequest) (*LoginResponse, error)
	LinkWallet(userID uuid.UUID, walletAddress string) error
}

type authUseCase struct {
	repo      Repository
	gameUC    game.UseCase
	jwtSecret string
}

func NewUseCase(repo Repository, gameUC game.UseCase, jwtSecret string) UseCase {
	return &authUseCase{
		repo:      repo,
		gameUC:    gameUC,
		jwtSecret: jwtSecret,
	}
}

func (u *authUseCase) Login(req LoginRequest) (*LoginResponse, error) {
	fmt.Printf("Login attempt for DID: %s, wallet: %s\n", req.PrivyDID, req.WalletAddress)
	
	if req.PrivyDID == "" && req.WalletAddress == "" {
		return nil, fmt.Errorf("invalid request: missing DID or wallet")
	}

	// 1. Try to find user by Privy DID first (most reliable)
	var user *User
	var err error
	if req.PrivyDID != "" {
		user, err = u.repo.GetByPrivyDID(req.PrivyDID)
		if err != nil {
			return nil, err
		}
	}

	// 2. If not found by DID, try wallet address
	if user == nil && req.WalletAddress != "" {
		user, err = u.repo.GetByWalletAddress(req.WalletAddress)
		if err != nil {
			return nil, err
		}
	}

	if user == nil {
		// Create new user if not exists
		walletAddress := req.WalletAddress
		username := "Pilot-"
		if walletAddress != "" && len(walletAddress) >= 6 {
			username += walletAddress[2:6]
		} else {
			username += uuid.New().String()[:4]
		}

		user = &User{
			ID:            uuid.New(),
			PrivyDID:      req.PrivyDID,
			WalletAddress: walletAddress,
			Username:      username,
			Credits:       0,
		}
		if err := u.repo.Create(user); err != nil {
			return nil, err
		}

		// Initialize Pilot, Gacha Stats, and Starter Gear for new user
		if err := u.gameUC.InitializeNewPlayer(user.ID); err != nil {
			return nil, fmt.Errorf("failed to initialize player: %w", err)
		}
	} else {
		// Update PrivyDID if it's missing (for old users found by wallet)
		if user.PrivyDID == "" && req.PrivyDID != "" {
			if err := u.repo.UpdatePrivyDID(user.ID, req.PrivyDID); err != nil {
				fmt.Printf("Warning: failed to update PrivyDID for user %s: %v\n", user.ID, err)
			} else {
				user.PrivyDID = req.PrivyDID
			}
		}

		// Update WalletAddress if it's missing (for users who logged in with social first)
		if user.WalletAddress == "" && req.WalletAddress != "" {
			if err := u.repo.UpdateWalletAddress(user.ID, req.WalletAddress); err != nil {
				fmt.Printf("Warning: failed to update WalletAddress for user %s: %v\n", user.ID, err)
			} else {
				user.WalletAddress = req.WalletAddress
			}
		}

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

func (u *authUseCase) LinkWallet(userID uuid.UUID, walletAddress string) error {
	if walletAddress == "" {
		return fmt.Errorf("wallet address is required")
	}
	return u.repo.UpdateWalletAddress(userID, walletAddress)
}
