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
	// TODO: Verify Privy Token
	// 1. Fetch JWKS from https://auth.privy.io/api/v1/apps/<app-id>/.well-known/jwks.json
	// 2. Verify req.PrivyToken signature and claims (iss, aud, exp)
	
	if req.PrivyToken == "" && req.WalletAddress == "" {
		return nil, fmt.Errorf("invalid request: missing token or wallet")
	}

	// For now, we use the wallet address provided or extract it from token (placeholder)
	walletAddress := req.WalletAddress
	if walletAddress == "" {
		// Placeholder: In reality, extract from verified token
		walletAddress = "0x0000000000000000000000000000000000000000" 
	}

	user, err := u.repo.GetByWalletAddress(walletAddress)
	if err != nil {
		return nil, err
	}

	if user == nil {
		// Create new user if not exists
		user = &User{
			ID:            uuid.New(),
			WalletAddress: walletAddress,
			Username:      "Pilot-" + walletAddress[2:6],
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
