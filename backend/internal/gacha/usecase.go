package gacha

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"github.com/ryudokung/Project-0/backend/internal/vehicle"
)

type UseCase interface {
	Pull(req GachaPullRequest) (*GachaPullResponse, error)
}

type gachaUseCase struct {
	repo        Repository
	gameRepo    game.Repository
	vehicleRepo vehicle.Repository
}

func NewUseCase(repo Repository, gameRepo game.Repository, vehicleRepo vehicle.Repository) UseCase {
	return &gachaUseCase{
		repo:        repo,
		gameRepo:    gameRepo,
		vehicleRepo: vehicleRepo,
	}
}

func (u *gachaUseCase) Pull(req GachaPullRequest) (*GachaPullResponse, error) {
	stats, err := u.gameRepo.GetGachaStats(req.UserID)
	if err != nil {
		return nil, err
	}
	if stats == nil {
		// Auto-initialize if missing for existing users
		if err := u.gameRepo.InitializeGachaStats(req.UserID); err != nil {
			return nil, fmt.Errorf("failed to initialize gacha stats: %w", err)
		}
		stats, err = u.gameRepo.GetGachaStats(req.UserID)
		if err != nil || stats == nil {
			return nil, fmt.Errorf("gacha stats not found after initialization")
		}
	}

	// Check Daily Signal
	if req.PullType == DailySignal {
		success, err := u.gameRepo.ConsumeFreePull(req.UserID)
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, fmt.Errorf("daily signal already used or not yet available")
		}
		// Refresh stats after consumption to get correct pity counts
		stats, _ = u.gameRepo.GetGachaStats(req.UserID)
		req.Count = 1
	}

	results := make([]GachaResult, 0, req.Count)

	for i := 0; i < req.Count; i++ {
		seed := time.Now().UnixNano()
		rand.Seed(seed)

		pityRelicBefore := stats.PityRelicCount
		pitySingularityBefore := stats.PitySingularityCount

		stats.TotalPulls++
		stats.PityRelicCount++
		stats.PitySingularityCount++

		rarity := u.rollRarity(stats)
		
		// Reset pity if we hit high rarity
		if rarity == vehicle.RaritySingularity {
			stats.PitySingularityCount = 0
			stats.PityRelicCount = 0
		} else if rarity == vehicle.RarityRelic {
			stats.PityRelicCount = 0
		}

		// Create the item (Simplified for now: 50% Vehicle, 50% Part)
		var result GachaResult
		if rand.Float64() < 0.5 {
			v := u.generateRandomVehicle(req.UserID, rarity)
			if err := u.vehicleRepo.Create(v); err != nil {
				return nil, err
			}
			result = GachaResult{
				Item:     v,
				ItemType: "VEHICLE",
				Rarity:   rarity,
			}
		} else {
			p := u.generateRandomPart(req.UserID, rarity)
			if err := u.vehicleRepo.CreatePart(p); err != nil {
				return nil, err
			}
			result = GachaResult{
				Item:     p,
				ItemType: "PART",
				Rarity:   rarity,
			}
		}

		// Save history with enhanced logging
		history := &GachaHistory{
			ID:                    uuid.New(),
			UserID:                req.UserID,
			ItemID:                u.getItemID(result.Item),
			ItemType:              result.ItemType,
			PullType:              req.PullType,
			Rarity:                rarity,
			Seed:                  seed,
			PityRelicBefore:       pityRelicBefore,
			PityRelicAfter:        stats.PityRelicCount,
			PitySingularityBefore: pitySingularityBefore,
			PitySingularityAfter:  stats.PitySingularityCount,
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

func (u *gachaUseCase) rollRarity(stats *game.GachaStats) vehicle.RarityTier {
	// 1. Check Singularity Pity (Hard pity at 80)
	if stats.PitySingularityCount >= 80 {
		return vehicle.RaritySingularity
	}

	// 2. Check Relic Pity (Hard pity at 10)
	if stats.PityRelicCount >= 10 {
		return vehicle.RarityRelic
	}

	// 3. Random Roll
	r := rand.Float64() * 100

	if r < 0.6 { // 0.6% for Singularity (Standard rate)
		return vehicle.RaritySingularity
	}
	if r < 5.1 { // 4.5% for Relic
		return vehicle.RarityRelic
	}
	if r < 15.0 { // 10% for Prototype
		return vehicle.RarityPrototype
	}
	if r < 40.0 { // 25% for Refined
		return vehicle.RarityRefined
	}
	return vehicle.RarityCommon
}

func (u *gachaUseCase) generateRandomVehicle(userID uuid.UUID, rarity vehicle.RarityTier) *vehicle.Vehicle {
	v := &vehicle.Vehicle{
		ID:            uuid.New(),
		OwnerID:       userID,
		VehicleType:   vehicle.TypeMech,
		Class:         vehicle.ClassStriker,
		Rarity:        rarity,
		Tier:          1,
		IsVoidTouched: true,
		Status:        vehicle.StatusPending,
		Stats: vehicle.VehicleStats{
			HP:      100,
			Attack:  10,
			Defense: 10,
			Speed:   10,
		},
		SuitabilityTags: []string{"urban"},
	}
	// Calculate CR
	v.CR = (v.Stats.Attack * 2) + (v.Stats.Defense * 2) + (v.Stats.HP / 10)
	return v
}

func (u *gachaUseCase) generateRandomPart(userID uuid.UUID, rarity vehicle.RarityTier) *vehicle.Part {
	return &vehicle.Part{
		ID:            uuid.New(),
		OwnerID:       userID,
		Slot:          "ARM_L",
		Name:          "Void Striker Arm",
		Rarity:        rarity,
		Tier:          1,
		IsVoidTouched: true,
		Stats: vehicle.PartStats{
			BonusAttack: 5,
		},
	}
}

func (u *gachaUseCase) getItemID(item interface{}) uuid.UUID {
	if v, ok := item.(*vehicle.Vehicle); ok {
		return v.ID
	}
	if p, ok := item.(*vehicle.Part); ok {
		return p.ID
	}
	return uuid.Nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
