package game

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/mech"
)

type UseCase interface {
	InitializeNewPlayer(userID uuid.UUID) error
}

type gameUseCase struct {
	repo     Repository
	mechRepo mech.Repository
}

func NewUseCase(repo Repository, mechRepo mech.Repository) UseCase {
	return &gameUseCase{
		repo:     repo,
		mechRepo: mechRepo,
	}
}

func (u *gameUseCase) InitializeNewPlayer(userID uuid.UUID) error {
	// 1. Initialize Stats
	if err := u.repo.InitializePilot(userID); err != nil {
		return err
	}

	// 2. Assign Starter Gear
	starterGear := []mech.Part{
		{
			ID:      uuid.New(),
			OwnerID: userID,
			Slot:    "PILOT_SUIT",
			Name:    "Nomad-01",
			Rarity:  mech.RarityCommon,
			Tier:    1,
			Stats:   map[string]interface{}{"defense": 5},
			VisualDNA: map[string]interface{}{
				"base": "Scavenged Fabric",
				"wear": "Light",
			},
		},
		{
			ID:      uuid.New(),
			OwnerID: userID,
			Slot:    "SIDEARM",
			Name:    "Rusty Bolt",
			Rarity:  mech.RarityCommon,
			Tier:    1,
			Stats:   map[string]interface{}{"attack": 3},
			VisualDNA: map[string]interface{}{
				"base": "Industrial Scrap",
				"wear": "Heavy",
			},
		},
		{
			ID:      uuid.New(),
			OwnerID: userID,
			Slot:    "O2_TANK",
			Name:    "Old Lung",
			Rarity:  mech.RarityCommon,
			Tier:    1,
			Stats:   map[string]interface{}{"o2_capacity": 50},
			VisualDNA: map[string]interface{}{
				"base": "Dented Steel",
				"wear": "Medium",
			},
		},
	}

	for _, p := range starterGear {
		if err := u.mechRepo.CreatePart(&p); err != nil {
			return err
		}
	}

	return nil
}
