package combat

import (
	"github.com/ryudokung/Project-0/backend/internal/mech"
	"github.com/ryudokung/Project-0/backend/internal/game"
)

type Service struct {
	engine *Engine
}

func NewService(engine *Engine) *Service {
	return &Service{engine: engine}
}

// MapMechToUnitStats converts a Mech and its equipped parts into UnitStats for the combat engine
func (s *Service) MapMechToUnitStats(m *mech.Mech, parts []mech.Part, pilot *game.PilotStats) UnitStats {
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

	if m != nil {
		stats.HP = m.Stats.HP
		stats.MaxHP = m.Stats.HP
		stats.BaseAttack = m.Stats.Attack
		stats.TargetDefense = m.Stats.Defense
		stats.DefenseEfficiency = 0.5 // Default from Bible
		stats.Accuracy = 80           // Base accuracy
		stats.Evasion = m.Stats.Speed / 10
		stats.Speed = m.Stats.Speed

		// Apply Part Bonuses
		for _, p := range parts {
			stats.HP += p.Stats.BonusHP
			stats.MaxHP += p.Stats.BonusHP
			stats.BaseAttack += p.Stats.BonusAttack
			stats.TargetDefense += p.Stats.BonusDefense
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

// ExecuteAttack runs a single attack cycle between two mechs
func (s *Service) ExecuteAttack(attackerStats UnitStats, defenderStats UnitStats, dmgType DamageType) CombatResult {
	return s.engine.CalculateDamage(attackerStats, defenderStats, dmgType)
}
