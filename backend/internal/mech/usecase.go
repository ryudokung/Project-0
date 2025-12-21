package mech

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type UseCase interface {
	MintStarterMech(userID uuid.UUID) (*Mech, error)
	GetUserMechs(userID uuid.UUID) ([]Mech, error)
}

type mechUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &mechUseCase{repo: repo}
}

func (u *mechUseCase) MintStarterMech(userID uuid.UUID) (*Mech, error) {
	// Randomly pick a type and class for starter
	types := []VehicleType{TypeMech, TypeTank, TypeShip}
	classes := []VehicleClass{ClassStriker, ClassGuardian, ClassScout}
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	m := &Mech{
		ID:          uuid.New(),
		OwnerID:     userID,
		VehicleType: types[r.Intn(len(types))],
		Class:       classes[r.Intn(len(classes))],
		Rarity:      RarityCommon,
		Season:      "Genesis",
		Status:      StatusMinted, // For starter, we mark as minted immediately in DB
		Stats: MechStats{
			HP:      50 + r.Intn(50),
			Attack:  10 + r.Intn(10),
			Defense: 10 + r.Intn(10),
			Speed:   5 + r.Intn(15),
		},
		ImageURL: "https://api.dicebear.com/7.x/bottts/svg?seed=" + uuid.New().String(), // Placeholder AI image
	}

	if err := u.repo.Create(m); err != nil {
		return nil, err
	}

	return m, nil
}

func (u *mechUseCase) GetUserMechs(userID uuid.UUID) ([]Mech, error) {
	return u.repo.GetByOwnerID(userID)
}
