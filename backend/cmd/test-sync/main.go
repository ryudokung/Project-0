package main

import (
	"fmt"
	"github.com/ryudokung/Project-0/backend/internal/combat"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"github.com/ryudokung/Project-0/backend/internal/vehicle"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("   TEST: SYNC RATE & ECP MULTIPLIER")
	fmt.Println("========================================")

	// Load Balance Config
	err := combat.LoadBalanceConfig("configs/game_balance.yaml")
	if err != nil {
		fmt.Printf("Failed to load balance config: %v\n", err)
		return
	}

	engine := combat.NewEngine()
	service := combat.NewService(engine)

	// 1. Setup Vehicle
	v := &vehicle.Vehicle{
		Stats: vehicle.VehicleStats{
			HP:      1000,
			Attack:  100,
			Defense: 50,
			Speed:   100,
		},
	}

	// 2. Test with different Sync Levels
	levels := []int{1, 11, 21} // Level 1 (50%), Level 11 (100%), Level 21 (150%)

	for _, lv := range levels {
		pilot := &game.PilotStats{
			SyncLevel:      lv,
			ResonanceLevel: 0,
		}

		stats := service.MapVehicleToUnitStats(v, nil, pilot)
		
		fmt.Printf("\n[Sync Level %d]\n", lv)
		fmt.Printf("Base Attack: %d -> Effective Attack: %d\n", v.Stats.Attack, stats.BaseAttack)
		fmt.Printf("Base Defense: %d -> Effective Defense: %d\n", v.Stats.Defense, stats.TargetDefense)
		
		// Verify multiplier
		expectedMultiplier := combat.GlobalBalance.Progression.BaseSyncRate + (float64(lv-1) * combat.GlobalBalance.Progression.SyncRatePerLevel)
		fmt.Printf("Expected Multiplier: %.2f\n", expectedMultiplier)
	}
}
