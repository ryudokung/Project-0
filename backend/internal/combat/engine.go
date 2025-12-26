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
	Shields          int     `json:"shields"` // Added for Energy damage bonus
	BaseAttack       int     `json:"base_attack"`
	TargetDefense    int     `json:"target_defense"`
	DefenseEfficiency float64 `json:"defense_efficiency"`
	Accuracy         int     `json:"accuracy"`
	Evasion          int     `json:"evasion"`
	Speed            int     `json:"speed"`
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
