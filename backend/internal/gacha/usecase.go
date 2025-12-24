package gacha

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"github.com/ryudokung/Project-0/backend/internal/mech"
)

type UseCase interface {
	Pull(req GachaPullRequest) (*GachaPullResponse, error)
}

type gachaUseCase struct {
	repo      Repository
	gameRepo  game.Repository
	mechRepo  mech.Repository
}

func NewUseCase(repo Repository, gameRepo game.Repository, mechRepo mech.Repository) UseCase {
	return &gachaUseCase{
		repo:     repo,
		gameRepo: gameRepo,
		mechRepo: mechRepo,
	}
}

func (u *gachaUseCase) Pull(req GachaPullRequest) (*GachaPullResponse, error) {
	stats, err := u.gameRepo.GetGachaStats(req.UserID)
	if err != nil {
		return nil, err
	}
	if stats == nil {
		return nil, fmt.Errorf("gacha stats not found for user")
	}

	results := make([]GachaResult, 0, req.Count)

	for i := 0; i < req.Count; i++ {
		stats.TotalPulls++
		stats.PityRelicCount++
		stats.PitySingularityCount++

		rarity := u.rollRarity(stats)
		
		// Reset pity if we hit high rarity
		if rarity == mech.RaritySingularity {
			stats.PitySingularityCount = 0
			stats.PityRelicCount = 0
		} else if rarity == mech.RarityRelic {
			stats.PityRelicCount = 0
		}

		// Create the item (Simplified for now: 50% Mech, 50% Part)
		var result GachaResult
		if rand.Float64() < 0.5 {
			m := u.generateRandomMech(req.UserID, rarity)
			if err := u.mechRepo.Create(m); err != nil {
				return nil, err
			}
			result = GachaResult{
				Item:     m,
				ItemType: "MECH",
				Rarity:   rarity,
			}
		} else {
			p := u.generateRandomPart(req.UserID, rarity)
			if err := u.mechRepo.CreatePart(p); err != nil {
				return nil, err
			}
			result = GachaResult{
				Item:     p,
				ItemType: "PART",
				Rarity:   rarity,
			}
		}

		// Save history
		history := &GachaHistory{
			ID:       uuid.New(),
			UserID:   req.UserID,
			ItemID:   u.getItemID(result.Item),
			ItemType: result.ItemType,
			PullType: req.PullType,
			Rarity:   rarity,
		}
		u.repo.SaveHistory(history)

		results = append(results, result)
	}

	// Update stats in DB
	if err := u.gameRepo.UpdateGachaStats(stats); err != nil {
		return nil, err
	}

	return &GachaPullResponse{Results: results}, nil
}

func (u *gachaUseCase) rollRarity(stats *game.GachaStats) mech.RarityTier {
	// 1. Check Singularity Pity (Hard pity at 80)
	if stats.PitySingularityCount >= 80 {
		return mech.RaritySingularity
	}

	// 2. Check Relic Pity (Hard pity at 10)
	if stats.PityRelicCount >= 10 {
		return mech.RarityRelic
	}

	// 3. Random Roll
	r := rand.Float64() * 100

	if r < 0.6 { // 0.6% for Singularity (Standard rate)
		return mech.RaritySingularity
	}
	if r < 5.1 { // 4.5% for Relic
		return mech.RarityRelic
	}
	if r < 15.0 { // 10% for Prototype
		return mech.RarityPrototype
	}
	if r < 40.0 { // 25% for Refined
		return mech.RarityRefined
	}
	return mech.RarityCommon
}

func (u *gachaUseCase) generateRandomMech(userID uuid.UUID, rarity mech.RarityTier) *mech.Mech {
	return &mech.Mech{
		ID:            uuid.New(),
		OwnerID:       userID,
		VehicleType:   mech.TypeMech,
		Class:         mech.ClassStriker,
		Rarity:        rarity,
		Tier:          1,
		IsVoidTouched: true,
		Status:        mech.StatusPending,
		Stats: mech.MechStats{
			HP:      100,
			Attack:  10,
			Defense: 10,
			Speed:   10,
		},
	}
}

func (u *gachaUseCase) generateRandomPart(userID uuid.UUID, rarity mech.RarityTier) *mech.Part {
	return &mech.Part{
		ID:            uuid.New(),
		OwnerID:       userID,
		Slot:          "ARM_L",
		Name:          "Void Striker Arm",
		Rarity:        rarity,
		Tier:          1,
		IsVoidTouched: true,
		Stats: mech.PartStats{
			BonusAttack: 5,
		},
	}
}

func (u *gachaUseCase) getItemID(item interface{}) uuid.UUID {
	if m, ok := item.(*mech.Mech); ok {
		return m.ID
	}
	if p, ok := item.(*mech.Part); ok {
		return p.ID
	}
	return uuid.Nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
