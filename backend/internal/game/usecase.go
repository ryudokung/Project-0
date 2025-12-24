package game

import (
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
			Stats:   mech.PartStats{BonusDefense: 5},
			VisualDNA: mech.VisualDNA{
				Keywords: []string{"Scavenged Fabric", "Light Wear"},
				Style:    "Nomad",
			},
		},
		{
			ID:      uuid.New(),
			OwnerID: userID,
			Slot:    "SIDEARM",
			Name:    "Rusty Bolt",
			Rarity:  mech.RarityCommon,
			Tier:    1,
			Stats:   mech.PartStats{BonusAttack: 3},
			VisualDNA: mech.VisualDNA{
				Keywords: []string{"Industrial Scrap", "Heavy Wear"},
				Style:    "Scavenger",
			},
		},
		{
			ID:      uuid.New(),
			OwnerID: userID,
			Slot:    "O2_TANK",
			Name:    "Old Lung",
			Rarity:  mech.RarityCommon,
			Tier:    1,
			Stats:   mech.PartStats{BonusHP: 10}, // Using HP as capacity proxy
			VisualDNA: mech.VisualDNA{
				Keywords: []string{"Dented Steel", "Medium Wear"},
				Style:    "Industrial",
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
