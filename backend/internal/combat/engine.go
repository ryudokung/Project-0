package combat

import (
	"math/rand"
	"time"
)

type DamageType string

const (
	Kinetic   DamageType = "KINETIC"
	Energy    DamageType = "ENERGY"
	Explosive DamageType = "EXPLOSIVE"
	Void      DamageType = "VOID"
)

type UnitStats struct {
	HP               int     `json:"hp"`
	MaxHP            int     `json:"max_hp"`
	Shields          int     `json:"shields"`
	BaseAttack       int     `json:"base_attack"`
	TargetDefense    int     `json:"target_defense"`
	DefenseEfficiency float64 `json:"defense_efficiency"`
	Accuracy         int     `json:"accuracy"`
	Evasion          int     `json:"evasion"`
	Speed            int     `json:"speed"`
	ResonanceLevel   int     `json:"resonance_level"` // 0 = Normal, >0 = Resonant
	ResonanceGauge   float64 `json:"resonance_gauge"` // 0-100
	IsResonanceActive bool    `json:"is_resonance_active"`
	IsVehicle        bool    `json:"is_vehicle"`      // To handle Scale Suppression
	IsPlayer         bool    `json:"is_player"`       // To handle scripted events
}

type CombatResult struct {
	FinalDamage  int            `json:"final_damage"`
	IsCritical   bool           `json:"is_critical"`
	IsMiss       bool           `json:"is_miss"`
	AppliedEffect *StatusEffect `json:"applied_effect,omitempty"`
}

type StatusEffectType string

const (
	Overheat    StatusEffectType = "OVERHEAT"
	ArmorBreach StatusEffectType = "ARMOR_BREACH"
	EngineStall StatusEffectType = "ENGINE_STALL"
)

var TypeMultipliers = map[DamageType]map[DamageType]float64{
	Kinetic: {
		Kinetic:   1.0,
		Energy:    1.5,
		Explosive: 0.5,
	},
	Energy: {
		Kinetic:   0.5,
		Energy:    1.0,
		Explosive: 1.5,
	},
	Explosive: {
		Kinetic:   1.5,
		Energy:    0.5,
		Explosive: 1.0,
	},
}

type StatusEffect struct {
	Type     StatusEffectType `json:"type"`
	Duration int              `json:"duration"` // in turns
}

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) CalculateDamage(attacker UnitStats, defender UnitStats, dmgType DamageType) CombatResult {
	rand.Seed(time.Now().UnixNano())

	// 1. Check for Miss
	missChance := defender.Evasion - (attacker.Accuracy / 10)
	if missChance > 0 && rand.Intn(100) < missChance {
		return CombatResult{FinalDamage: 0, IsMiss: true}
	}

	// 2. Calculate Base Damage & Apply Damage Matrix
	baseDmg := float64(attacker.BaseAttack)
	
	// Scale Suppression & Resonance Logic
	if !attacker.IsVehicle && defender.IsVehicle {
		// Human attacking Vehicle
		if attacker.IsResonanceActive {
			// Resonance bypasses suppression and adds bonus
			resonanceBonus := 1.0 + (float64(attacker.ResonanceLevel) * GlobalBalance.Resonance.ResonanceDamageMultiplier)
			baseDmg *= resonanceBonus
		} else {
			// Normal human vs Vehicle: 90% damage reduction
			baseDmg *= GlobalBalance.ScaleSuppression.HumanVsMechDamageReduction
		}
	} else if attacker.IsVehicle && !defender.IsVehicle {
		// Vehicle attacking Human
		if defender.IsResonanceActive {
			// Resonant human can partially dodge/deflect vehicle-scale damage
			if attacker.IsResonanceActive {
				// If both are resonant, the mech's power is harder to dodge
				baseDmg *= GlobalBalance.ScaleSuppression.ResonantHumanVsMechDamageMultiplier 
			} else {
				baseDmg *= GlobalBalance.ScaleSuppression.ResonantHumanDeflectionRate
			}
		} else {
			// Vehicle vs normal human: 300% damage (Overkill)
			baseDmg *= GlobalBalance.ScaleSuppression.MechVsHumanDamageMultiplier
		}
	}

	// Defense Calculation
	defense := float64(defender.TargetDefense) * defender.DefenseEfficiency

	// Damage Matrix Logic
	switch dmgType {
	case Energy:
		// +20% Damage vs Shields
		if defender.Shields > 0 {
			baseDmg *= 1.2
		}
	case Void:
		// Ignores 30% Defense
		defense *= 0.7
	case Kinetic:
		// Standard (High Base - usually handled by higher base stats on Kinetic weapons)
	}

	finalDmg := baseDmg - defense
	if finalDmg < 1 {
		finalDmg = 1 // Minimum 1 damage
	}

	// 3. Check for Critical Hit
	critChance := 5 + (attacker.Accuracy / 100)
	isCritical := rand.Intn(100) < critChance
	if isCritical {
		finalDmg *= 1.5
	}

	// 4. Determine Status Effect (Simplified chance)
	var effect *StatusEffect
	if !isCritical && rand.Intn(100) < 20 { // 20% chance on normal hit
		switch dmgType {
		case Energy:
			effect = &StatusEffect{Type: Overheat, Duration: 2}
		case Kinetic:
			effect = &StatusEffect{Type: ArmorBreach, Duration: 99} // Permanent
		case Explosive:
			effect = &StatusEffect{Type: EngineStall, Duration: 1}
		}
	}

	return CombatResult{
		FinalDamage:   int(finalDmg),
		IsCritical:    isCritical,
		IsMiss:        false,
		AppliedEffect: effect,
	}
}
