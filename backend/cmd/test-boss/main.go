package main

import (
	"fmt"
	"github.com/ryudokung/Project-0/backend/internal/combat"
	"github.com/ryudokung/Project-0/backend/internal/game"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("   TEST: BOSS PHASE TRANSITION")
	fmt.Println("   Mission: The Iron Awakening")
	fmt.Println("   Boss: The Gatekeeper")
	fmt.Println("========================================")

	// Load Balance Config
	err := combat.LoadBalanceConfig("configs/game_balance.yaml")
	if err != nil {
		fmt.Printf("Failed to load balance config: %v\n", err)
		return
	}

	engine := combat.NewEngine()
	service := combat.NewService(engine)

	// 1. Setup Player (Mech)
	playerStats := combat.UnitStats{
		HP:                1000, // High HP to survive long enough
		MaxHP:             1000,
		BaseAttack:        100,  // High attack to kill boss mech
		TargetDefense:     10,
		DefenseEfficiency: combat.GlobalBalance.BaseStats.DefaultDefenseEfficiency,
		Accuracy:          100,
		Evasion:           20,
		IsVehicle:         true,
		IsPlayer:          true,
	}

	// 2. Setup Boss (The Gatekeeper Mech)
	bossStats := combat.UnitStats{
		HP:                300, // Lower HP to trigger phase 2 quickly
		MaxHP:             300,
		BaseAttack:        20,  // Lower attack to not kill player too fast
		TargetDefense:     30,
		DefenseEfficiency: 0.6,
		Accuracy:          90,
		Evasion:           10,
		IsVehicle:         true,
		IsPlayer:          false,
	}

	// 3. Setup Scripted Events (from expeditions.yaml)
	scriptEvents := []game.ScriptEvent{
		{
			Trigger:  "player_hp_low",
			Action:   "force_eject",
			Dialogue: "Warning: Hull integrity critical! Ejecting pilot...",
		},
		{
			Trigger:  "boss_phase_2",
			Action:   "spawn_human_pilot",
			Dialogue: "You think a machine is all I am? Witness the power of Resonance!",
		},
	}

	session := &combat.CombatSession{
		PlayerStats:  playerStats,
		EnemyStats:   bossStats,
		IsScripted:   true,
		ScriptEvents: scriptEvents,
	}

	fmt.Println("\n--- PHASE 1: MECH VS MECH ---")
	
	turn := 1
	for {
		fmt.Printf("\n[TURN %d]\n", turn)
		fmt.Printf("Player HP: %d (%s) | Boss HP: %d (%s)\n", 
			session.PlayerStats.HP, 
			func() string { if session.PlayerStats.IsVehicle { return "MECH" }; return "HUMAN" }(),
			session.EnemyStats.HP, 
			func() string { if session.EnemyStats.IsVehicle { return "MECH" }; return "HUMAN" }())

		// Player Attacks
		res := service.ExecuteAttack(session, combat.Kinetic)
		fmt.Printf(">> Player attacks: %d damage (Gauge: %.1f%%)\n", res.FinalDamage, session.PlayerStats.ResonanceGauge)
		
		// Check for Resonance Activation
		if !session.PlayerStats.IsResonanceActive && session.PlayerStats.ResonanceGauge >= 100 {
			service.ActivateResonance(session)
			fmt.Println("!!! PLAYER ACTIVATED RESONANCE MODE !!!")
		}

		// Check for logs (Scripted Events)
		for _, log := range session.Log {
			fmt.Printf("!!! EVENT: %s\n", log)
		}
		session.Log = nil

		if session.EnemyStats.HP <= 0 && !session.EnemyStats.IsVehicle {
			fmt.Println("\n*** BOSS DEFEATED! ***")
			break
		}

		// Boss Counters
		// To counter-attack correctly, we swap roles but keep the session's Player/Enemy context
		// We'll use a temporary session for the counter-attack calculation
		counterSession := &combat.CombatSession{
			PlayerStats:  session.EnemyStats,  // Boss is attacker
			EnemyStats:   session.PlayerStats, // Player is defender
			IsScripted:   true,
			ScriptEvents: scriptEvents,
		}
		resC := service.ExecuteAttack(counterSession, combat.Kinetic)
		
		// Sync back the results
		session.EnemyStats = counterSession.PlayerStats
		session.PlayerStats = counterSession.EnemyStats

		// Build gauge from taking damage too
		session.PlayerStats.ResonanceGauge += float64(resC.FinalDamage) * combat.GlobalBalance.Resonance.GainRateTaken
		if session.PlayerStats.ResonanceGauge > 100 {
			session.PlayerStats.ResonanceGauge = 100
		}

		fmt.Printf("<< Boss counters: %d damage\n", resC.FinalDamage)
		
		for _, log := range counterSession.Log {
			fmt.Printf("!!! EVENT: %s\n", log)
		}

		if session.PlayerStats.HP <= 0 && !session.PlayerStats.IsVehicle {
			fmt.Println("\n*** PLAYER DESTROYED! ***")
			break
		}

		turn++
		if turn > 50 { break } // Safety break
	}

	fmt.Println("\n--- TEST COMPLETE ---")
}
