package combat

import (
	"testing"
)

func TestCalculateDamage(t *testing.T) {
	attacker := UnitStats{
		BaseAttack:       100,
		Accuracy:         80,
		DefenseEfficiency: 0.5,
	}

	defender := UnitStats{
		TargetDefense:    50,
		DefenseEfficiency: 0.5,
		Evasion:          10,
	}

	// Test Kinetic vs Kinetic (1.0x)
	// Expected: (100 * 1.0) - (50 * 0.5) = 100 - 25 = 75
	result := CalculateDamage(attacker, defender, Kinetic)
	
	if !result.IsMiss {
		if !result.IsCritical && result.FinalDamage != 75 {
			t.Errorf("Expected 75 damage, got %d", result.FinalDamage)
		}
		if result.IsCritical && result.FinalDamage != 112 { // 75 * 1.5 = 112.5 -> 112
			t.Errorf("Expected 112 critical damage, got %d", result.FinalDamage)
		}
	}
}

func TestStatusEffects(t *testing.T) {
	attacker := UnitStats{BaseAttack: 100, Accuracy: 100}
	defender := UnitStats{TargetDefense: 0, DefenseEfficiency: 0}

	// We might need to run multiple times to catch the 20% chance
	foundEffect := false
	for i := 0; i < 100; i++ {
		result := CalculateDamage(attacker, defender, Energy)
		if result.AppliedEffect != nil && result.AppliedEffect.Type == Overheat {
			foundEffect = true
			break
		}
	}

	if !foundEffect {
		t.Errorf("Should have applied Overheat effect at least once in 100 tries")
	}
}

func TestTypeMultipliers(t *testing.T) {
	if TypeMultipliers[Kinetic][Energy] != 1.5 {
		t.Errorf("Kinetic vs Energy should be 1.5x")
	}
	if TypeMultipliers[Energy][Explosive] != 1.5 {
		t.Errorf("Energy vs Explosive should be 1.5x")
	}
	if TypeMultipliers[Explosive][Kinetic] != 1.5 {
		t.Errorf("Explosive vs Kinetic should be 1.5x")
	}
}
