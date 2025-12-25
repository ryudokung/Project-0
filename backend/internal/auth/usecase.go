package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	Login(req LoginRequest) (*LoginResponse, error)
	Signup(req SignupRequest) (*LoginResponse, error)
	ValidateToken(tokenString string) (*User, error)
	GetMe(userID uuid.UUID) (*User, error)
	LinkWallet(userID uuid.UUID, walletAddress string) error

	// Character methods
	CreateCharacter(userID uuid.UUID, req CreateCharacterRequest) (*Character, error)
	GetCharacters(userID uuid.UUID) ([]Character, error)
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
	var user *User
	var err error

	// 1. Traditional Login (Username/Password)
	if req.Username != "" && req.Password != "" {
		user, err = u.repo.GetByUsername(req.Username)
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, fmt.Errorf("user not found")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			return nil, fmt.Errorf("invalid password")
		}
	} else if req.GuestID != "" {
		// 2. Guest Login
		user, err = u.repo.GetByGuestID(req.GuestID)
		if err != nil {
			return nil, err
		}
		if user == nil {
			// Create new Guest
			user = &User{
				ID:       uuid.New(),
				Username: "Guest-" + uuid.New().String()[:4],
				GuestID:  req.GuestID,
				AuthType: "GUEST",
				Credits:  0,
			}
			if err := u.repo.Create(user); err != nil {
				return nil, err
			}
			// Game initialization now happens during Character Creation
		}
	} else {
		// 3. Social/Wallet Login (Privy)
		if req.PrivyDID == "" && req.WalletAddress == "" {
			return nil, fmt.Errorf("invalid request: missing DID or wallet")
		}

		if req.PrivyDID != "" {
			user, err = u.repo.GetByPrivyDID(req.PrivyDID)
			if err != nil {
				return nil, err
			}
		}

		if user == nil && req.WalletAddress != "" {
			user, err = u.repo.GetByWalletAddress(req.WalletAddress)
			if err != nil {
				return nil, err
			}
		}

		if user == nil {
			// Check if we can bind from Guest
			if req.GuestID != "" {
				user, _ = u.repo.GetByGuestID(req.GuestID)
				if user != nil && user.AuthType == "GUEST" {
					// Bind Guest to Social
					if err := u.repo.UpdatePrivyDID(user.ID, req.PrivyDID); err != nil {
						return nil, fmt.Errorf("failed to bind privy did: %w", err)
					}
					user.PrivyDID = req.PrivyDID
					user.AuthType = "SOCIAL"
					if req.WalletAddress != "" {
						if err := u.repo.UpdateWalletAddress(user.ID, req.WalletAddress); err != nil {
							return nil, fmt.Errorf("failed to bind wallet: %w", err)
						}
						user.WalletAddress = req.WalletAddress
					}
				}
			}

			if user == nil {
				// Create new Social User
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
					AuthType:      "SOCIAL",
					Credits:       0,
				}
				if err := u.repo.Create(user); err != nil {
					return nil, err
				}
				// Game initialization now happens during Character Creation
			}
		} else {
			// Update existing user
			if user.PrivyDID == "" && req.PrivyDID != "" {
				if err := u.repo.UpdatePrivyDID(user.ID, req.PrivyDID); err != nil {
					return nil, fmt.Errorf("failed to update privy did: %w", err)
				}
				user.PrivyDID = req.PrivyDID
			}
			if user.WalletAddress == "" && req.WalletAddress != "" {
				if err := u.repo.UpdateWalletAddress(user.ID, req.WalletAddress); err != nil {
					return nil, fmt.Errorf("failed to update wallet: %w", err)
				}
				user.WalletAddress = req.WalletAddress
			}
		}
	}

	if user != nil {
		u.repo.UpdateLastLogin(user.ID)
		
		// Populate ActiveCharacter if exists
		if user.ActiveCharacterID != nil {
			char, _ := u.repo.GetCharacterByID(*user.ActiveCharacterID)
			user.ActiveCharacter = char
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

func (u *authUseCase) Signup(req SignupRequest) (*LoginResponse, error) {
	// 1. Check if username exists
	existing, _ := u.repo.GetByUsername(req.Username)
	if existing != nil {
		return nil, fmt.Errorf("username already taken")
	}

	// 2. Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var user *User
	// 3. Check if upgrading from Guest
	if req.GuestID != "" {
		user, _ = u.repo.GetByGuestID(req.GuestID)
		if user != nil {
			if user.AuthType != "GUEST" {
				return nil, fmt.Errorf("account already upgraded")
			}
			// Upgrade Guest to Traditional
			if err := u.repo.UpdateToTraditional(user.ID, req.Username, req.Email, string(hashed)); err != nil {
				return nil, err
			}
			user.Username = req.Username
			user.Email = req.Email
			user.AuthType = "TRADITIONAL"
		}
	}

	if user == nil {
		// Create new Traditional User
		user = &User{
			ID:           uuid.New(),
			Username:     req.Username,
			Email:        req.Email,
			PasswordHash: string(hashed),
			AuthType:     "TRADITIONAL",
			Credits:      0,
		}
		if err := u.repo.Create(user); err != nil {
			return nil, err
		}
		// Game initialization now happens during Character Creation
	}

	// Populate ActiveCharacter if exists
	if user.ActiveCharacterID != nil {
		char, _ := u.repo.GetCharacterByID(*user.ActiveCharacterID)
		user.ActiveCharacter = char
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: tokenString,
		User:  *user,
	}, nil
}

func (u *authUseCase) ValidateToken(tokenString string) (*User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return nil, fmt.Errorf("user_id not found in token")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id format")
	}

	user, err := u.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	if user != nil && user.ActiveCharacterID != nil {
		char, _ := u.repo.GetCharacterByID(*user.ActiveCharacterID)
		user.ActiveCharacter = char
	}
	return user, nil
}

func (u *authUseCase) GetMe(userID uuid.UUID) (*User, error) {
	user, err := u.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	if user != nil && user.ActiveCharacterID != nil {
		char, _ := u.repo.GetCharacterByID(*user.ActiveCharacterID)
		user.ActiveCharacter = char
	}
	return user, nil
}

func (u *authUseCase) LinkWallet(userID uuid.UUID, walletAddress string) error {
	if walletAddress == "" {
		return fmt.Errorf("wallet address is required")
	}
	return u.repo.UpdateWalletAddress(userID, walletAddress)
}

func (u *authUseCase) CreateCharacter(userID uuid.UUID, req CreateCharacterRequest) (*Character, error) {
	// Anti-Cheat: Check if user already has characters (limit 1 for now)
	existing, _ := u.repo.GetCharactersByUserID(userID)
	if len(existing) > 0 {
		return nil, fmt.Errorf("user already has a character")
	}

	char := &Character{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      req.Name,
		Gender:    req.Gender,
		FaceIndex: req.FaceIndex,
		HairIndex: req.HairIndex,
	}

	if err := u.repo.CreateCharacter(char); err != nil {
		return nil, fmt.Errorf("failed to create character: %w", err)
	}

	// Initialize game stats and starter ship for the character
	if err := u.gameUC.InitializeNewCharacter(userID, char.ID); err != nil {
		return nil, fmt.Errorf("failed to initialize character game data: %w", err)
	}

	// Set as active character
	if err := u.repo.SetActiveCharacter(userID, char.ID); err != nil {
		return nil, fmt.Errorf("failed to set active character: %w", err)
	}

	return char, nil
}

func (u *authUseCase) GetCharacters(userID uuid.UUID) ([]Character, error) {
	return u.repo.GetCharactersByUserID(userID)
}
