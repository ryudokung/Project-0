package combat

import (
	"github.com/ryudokung/Project-0/backend/internal/vehicle"
	"github.com/ryudokung/Project-0/backend/internal/game"
)

type Service struct {
	engine *Engine
}

func NewService(engine *Engine) *Service {
	return &Service{engine: engine}
}

// MapVehicleToUnitStats converts a Vehicle and its equipped items into UnitStats for the combat engine
func (s *Service) MapVehicleToUnitStats(v *vehicle.Vehicle, items []vehicle.Item, pilot *game.PilotStats) UnitStats {
	// Default stats for Pilot Only mode
	stats := UnitStats{
		HP:                100,
		MaxHP:             100,
		BaseAttack:        10,
		TargetDefense:     5,
		DefenseEfficiency: 0.3,
		Accuracy:          70,
		Evasion:           10,
		Speed:             50,
	}

	if v != nil {
		stats.HP = v.Stats.HP
		stats.MaxHP = v.Stats.HP
		stats.BaseAttack = v.Stats.Attack
		stats.TargetDefense = v.Stats.Defense
		stats.DefenseEfficiency = 0.5 // Default from Bible
		stats.Accuracy = 80           // Base accuracy
		stats.Evasion = v.Stats.Speed / 10
		stats.Speed = v.Stats.Speed

		// Apply Item Bonuses
		for _, i := range items {
			// Only count equipped items (though the query should filter this)
			if i.IsEquipped {
				stats.HP += i.Stats.BonusHP
				stats.MaxHP += i.Stats.BonusHP
				stats.BaseAttack += i.Stats.BonusAttack
				stats.TargetDefense += i.Stats.BonusDefense
			}
		}
	}

	// Apply Neural Resonance Bonus (Newtype effect)
	if pilot != nil {
		// Every level of resonance increases Accuracy and Evasion by 2%
		stats.Accuracy += pilot.ResonanceLevel * 2
		stats.Evasion += pilot.ResonanceLevel * 2
	}

	return stats
}

// ExecuteAttack runs a single attack cycle between two vehicles
func (s *Service) ExecuteAttack(attackerStats UnitStats, defenderStats UnitStats, dmgType DamageType) CombatResult {
	return s.engine.CalculateDamage(attackerStats, defenderStats, dmgType)
}
