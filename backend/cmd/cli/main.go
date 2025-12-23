package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ryudokung/Project-0/backend/internal/combat"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"github.com/ryudokung/Project-0/backend/internal/mech"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/project0?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mechRepo := mech.NewRepository(db)
	gameRepo := game.NewRepository(db)
	combatEngine := combat.NewEngine()
	combatService := combat.NewService(combatEngine)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("========================================")
	fmt.Println("   PROJECT-0: TACTICAL NOIR CLI")
	fmt.Println("========================================")

	// 1. Get All Mechs
	// For simplicity in CLI, we'll just fetch all mechs in the DB
	rows, _ := db.Query("SELECT id, class, rarity FROM mechs")
	var mechs []struct {
		ID     uuid.UUID
		Class  string
		Rarity string
	}
	fmt.Println("\nAvailable Mechs in Database:")
	i := 0
	for rows.Next() {
		var m struct {
			ID     uuid.UUID
			Class  string
			Rarity string
		}
		rows.Scan(&m.ID, &m.Class, &m.Rarity)
		mechs = append(mechs, m)
		fmt.Printf("[%d] ID: %s | Class: %s | Rarity: %s\n", i, m.ID.String()[:8], m.Class, m.Rarity)
		i++
	}
	rows.Close()

	if len(mechs) < 2 {
		fmt.Println("Not enough mechs in DB. Please run seed first.")
		return
	}

	// 2. Select Attacker
	fmt.Print("\nSelect Attacker Index: ")
	input, _ := reader.ReadString('\n')
	var attackerIdx int
	fmt.Sscanf(input, "%d", &attackerIdx)

	// 3. Select Defender
	fmt.Print("Select Defender Index: ")
	input, _ = reader.ReadString('\n')
	var defenderIdx int
	fmt.Sscanf(input, "%d", &defenderIdx)

	attackerID := mechs[attackerIdx].ID
	defenderID := mechs[defenderIdx].ID

	// 4. Load Full Data
	ctx := context.Background()
	aMech, _ := mechRepo.GetByID(ctx, attackerID)
	dMech, _ := mechRepo.GetByID(ctx, defenderID)
	aParts, _ := mechRepo.GetPartsByMechID(attackerID)
	dParts, _ := mechRepo.GetPartsByMechID(defenderID)
	aPilot, _ := gameRepo.GetPilotStats(aMech.OwnerID)
	dPilot, _ := gameRepo.GetPilotStats(dMech.OwnerID)

	aStats := combatService.MapMechToUnitStats(aMech, aParts, aPilot)
	dStats := combatService.MapMechToUnitStats(dMech, dParts, dPilot)

	fmt.Printf("\n--- BATTLE START ---\n")
	fmt.Printf("Attacker: %s (HP: %d, ATK: %d)\n", aMech.Class, aStats.HP, aStats.BaseAttack)
	fmt.Printf("Defender: %s (HP: %d, DEF: %d)\n", dMech.Class, dStats.HP, dStats.TargetDefense)

	for aStats.HP > 0 && dStats.HP > 0 {
		fmt.Println("\nChoose Damage Type: [K]inetic, [E]nergy, [X]plosive")
		fmt.Print("> ")
		choice, _ := reader.ReadString('\n')
		choice = strings.ToUpper(strings.TrimSpace(choice))

		var dmgType combat.DamageType
		switch choice {
		case "K":
			dmgType = combat.Kinetic
		case "E":
			dmgType = combat.Energy
		case "X":
			dmgType = combat.Explosive
		default:
			dmgType = combat.Kinetic
		}

		// Attacker hits Defender
		res := combatService.ExecuteAttack(aStats, dStats, dmgType)
		
		fmt.Printf("\n>> ATTACK! Type: %s\n", dmgType)
		if res.IsMiss {
			fmt.Println("MISS!")
		} else {
			critStr := ""
			if res.IsCritical {
				critStr = " CRITICAL!"
			}
			fmt.Printf("DAMAGE: %d%s\n", res.FinalDamage, critStr)
			dStats.HP -= res.FinalDamage
			if res.AppliedEffect != nil {
				fmt.Printf("EFFECT APPLIED: %s (%d turns)\n", res.AppliedEffect.Type, res.AppliedEffect.Duration)
			}
		}

		if dStats.HP <= 0 {
			fmt.Println("\n*** DEFENDER DESTROYED! ***")
			break
		}

		fmt.Printf("Defender HP remaining: %d\n", dStats.HP)

		// Simple Counter Attack (Defender hits Attacker)
		fmt.Println("\n-- Defender Counters! --")
		resCounter := combatService.ExecuteAttack(dStats, aStats, combat.Kinetic)
		if resCounter.IsMiss {
			fmt.Println("Counter MISS!")
		} else {
			fmt.Printf("Counter DAMAGE: %d\n", resCounter.FinalDamage)
			aStats.HP -= resCounter.FinalDamage
		}

		if aStats.HP <= 0 {
			fmt.Println("\n*** ATTACKER DESTROYED! ***")
			break
		}
		fmt.Printf("Attacker HP remaining: %d\n", aStats.HP)
	}

	fmt.Println("\n========================================")
	fmt.Println("           BATTLE OVER")
	fmt.Println("========================================")
}
