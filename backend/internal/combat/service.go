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
		stats.IsVehicle = true
		stats.HP = v.Stats.HP
		stats.MaxHP = v.Stats.HP
		stats.BaseAttack = v.Stats.Attack
		stats.TargetDefense = v.Stats.Defense
		stats.DefenseEfficiency = GlobalBalance.BaseStats.DefaultDefenseEfficiency
		stats.Accuracy = GlobalBalance.BaseStats.DefaultAccuracy
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
	} else {
		stats.IsVehicle = false
	}

	// Apply Neural Resonance Bonus (Newtype effect)
	if pilot != nil {
		stats.IsPlayer = true
		stats.ResonanceLevel = pilot.ResonanceLevel
		stats.ResonanceGauge = pilot.ResonanceGauge
		if active, ok := pilot.Metadata["resonance_active"].(bool); ok && active {
			stats.IsResonanceActive = true
		}

		// Calculate Sync Rate Multiplier
		syncRate := GlobalBalance.Progression.BaseSyncRate + (float64(pilot.SyncLevel-1) * GlobalBalance.Progression.SyncRatePerLevel)
		
		// Apply Sync Rate to Combat Stats (ECP logic)
		stats.BaseAttack = int(float64(stats.BaseAttack) * syncRate)
		stats.TargetDefense = int(float64(stats.TargetDefense) * syncRate)

		// Every level of resonance increases Accuracy and Evasion
		stats.Accuracy += pilot.ResonanceLevel * GlobalBalance.Resonance.BonusAccuracyPerLevel
		stats.Evasion += pilot.ResonanceLevel * GlobalBalance.Resonance.BonusEvasionPerLevel
	}

	return stats
}

type ScriptAction string

const (
	ActionForceEject      ScriptAction = "force_eject"
	ActionSpawnHumanPilot ScriptAction = "spawn_human_pilot"
)

type CombatSession struct {
	PlayerStats   UnitStats
	EnemyStats    UnitStats
	IsScripted    bool
	ScriptEvents  []game.ScriptEvent // We'll need to import game or move the struct
	TurnCount     int
	Log           []string
}

// ExecuteAttack runs a single attack cycle between two vehicles
func (s *Service) ExecuteAttack(session *CombatSession, dmgType DamageType) CombatResult {
	result := s.engine.CalculateDamage(session.PlayerStats, session.EnemyStats, dmgType)
	
	// Update HP
	session.EnemyStats.HP -= result.FinalDamage
	if session.EnemyStats.HP < 0 {
		session.EnemyStats.HP = 0
	}

	// Build Resonance Gauge for Player
	if session.PlayerStats.IsPlayer && !session.PlayerStats.IsResonanceActive {
		// Gain gauge based on damage dealt
		gain := float64(result.FinalDamage) * GlobalBalance.Resonance.GainRateDealt
		session.PlayerStats.ResonanceGauge += gain
		if session.PlayerStats.ResonanceGauge > 100 {
			session.PlayerStats.ResonanceGauge = 100
		}
	}

	// Check for Scripted Triggers
	if session.IsScripted {
		s.handleScriptedEvents(session)
	}

	return result
}

// ActivateResonance triggers Resonance Mode if the gauge is full
func (s *Service) ActivateResonance(session *CombatSession) bool {
	if session.PlayerStats.ResonanceGauge >= 100 {
		session.PlayerStats.IsResonanceActive = true
		session.PlayerStats.ResonanceGauge = 0
		session.Log = append(session.Log, "[SYSTEM] NEURAL RESONANCE SYNCHRONIZED. SCALE SUPPRESSION BYPASSED.")
		return true
	}
	return false
}

func (s *Service) handleScriptedEvents(session *CombatSession) {
	// 1. Player HP Low Trigger
	if session.PlayerStats.IsPlayer && session.PlayerStats.HP < (session.PlayerStats.MaxHP / 4) && session.PlayerStats.IsVehicle {
		for _, event := range session.ScriptEvents {
			if event.Trigger == "player_hp_low" && event.Action == string(ActionForceEject) {
				session.PlayerStats.IsVehicle = false
				session.PlayerStats.HP = GlobalBalance.BaseStats.ForcedSurvivalHP
				session.Log = append(session.Log, "[SCRIPT] "+event.Dialogue)
			}
		}
	}

	// 2. Boss HP Zero Trigger (Phase 2)
	if !session.EnemyStats.IsPlayer && session.EnemyStats.HP <= 0 && session.EnemyStats.IsVehicle {
		for _, event := range session.ScriptEvents {
			if event.Trigger == "boss_phase_2" && event.Action == string(ActionSpawnHumanPilot) {
				// Transform Boss to Human Pilot
				session.EnemyStats.IsVehicle = false
				session.EnemyStats.HP = GlobalBalance.BaseStats.BossPhase2HP
				session.EnemyStats.MaxHP = GlobalBalance.BaseStats.BossPhase2HP
				session.EnemyStats.BaseAttack = GlobalBalance.BaseStats.BossPhase2Attack
				session.EnemyStats.IsResonanceActive = true
				session.EnemyStats.ResonanceLevel = GlobalBalance.BaseStats.BossPhase2ResonanceLevel
				session.Log = append(session.Log, "[SCRIPT] "+event.Dialogue)
			}
		}
	}
}
