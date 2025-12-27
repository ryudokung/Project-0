package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ryudokung/Project-0/backend/internal/combat"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"github.com/ryudokung/Project-0/backend/internal/vehicle"
)

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	titleStyle = lipgloss.NewStyle().
			MarginLeft(2).
			MarginRight(2).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#F25D94")).
			Bold(true)

	statsStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(subtle).
			Padding(1, 2).
			Width(40)

	logStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(subtle).
			Padding(1, 2).
			Height(10).
			Width(84)
)

type model struct {
	db            *sql.DB
	vehicleRepo   vehicle.Repository
	gameRepo      game.Repository
	combatService *combat.Service
	
	attacker      *vehicle.Vehicle
	defender      *vehicle.Vehicle
	aStats        combat.UnitStats
	dStats        combat.UnitStats
	aPilot        *game.PilotStats
	dPilot        *game.PilotStats
	
	aHPBar        progress.Model
	dHPBar        progress.Model
	
	logs          []string
	gameOver      bool
	winner        string
}

func initialModel(db *sql.DB) model {
	vehicleRepo := vehicle.NewRepository(db)
	gameRepo := game.NewRepository(db)
	combatEngine := combat.NewEngine()
	combatService := combat.NewService(combatEngine)

	// Fetch first two vehicles for demo
	rows, _ := db.Query("SELECT id FROM vehicles LIMIT 2")
	var ids []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		rows.Scan(&id)
		ids = append(ids, id)
	}
	rows.Close()

	ctx := context.Background()
	aVehicle, _ := vehicleRepo.GetByID(ctx, ids[0])
	dVehicle, _ := vehicleRepo.GetByID(ctx, ids[1])
	aItems, _ := vehicleRepo.GetItemsByParentItemID(ctx, ids[0])
	dItems, _ := vehicleRepo.GetItemsByParentItemID(ctx, ids[1])
	aPilot, _ := gameRepo.GetActivePilotStats(aVehicle.OwnerID)
	dPilot, _ := gameRepo.GetActivePilotStats(dVehicle.OwnerID)

	aStats := combatService.MapVehicleToUnitStats(aVehicle, aItems, aPilot)
	dStats := combatService.MapVehicleToUnitStats(dVehicle, dItems, dPilot)

	return model{
		db:            db,
		vehicleRepo:   vehicleRepo,
		gameRepo:      gameRepo,
		combatService: combatService,
		attacker:      aVehicle,
		defender:      dVehicle,
		aStats:        aStats,
		dStats:        dStats,
		aPilot:        aPilot,
		dPilot:        dPilot,
		aHPBar:        progress.New(progress.WithDefaultGradient()),
		dHPBar:        progress.New(progress.WithDefaultGradient()),
		logs:          []string{"Battle Started! Tactical Noir Protocol Engaged."},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.gameOver {
			return m, tea.Quit
		}
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1", "2", "3":
			var dmgType combat.DamageType
			switch msg.String() {
			case "1": dmgType = combat.Kinetic
			case "2": dmgType = combat.Energy
			case "3": dmgType = combat.Explosive
			}
			return m.handleAttack(dmgType)
		}
	}
	return m, nil
}

func (m model) handleAttack(dmgType combat.DamageType) (model, tea.Cmd) {
	// Attacker hits Defender
	res := m.combatService.ExecuteAttack(m.aStats, m.dStats, dmgType)
	logMsg := fmt.Sprintf("Attacker used %s: ", dmgType)
	if res.IsMiss {
		logMsg += "MISS!"
	} else {
		m.dStats.HP -= res.FinalDamage
		logMsg += fmt.Sprintf("HIT for %d damage!", res.FinalDamage)
		if res.IsCritical {
			logMsg += " CRITICAL!"
		}
		if res.AppliedEffect != nil {
			logMsg += fmt.Sprintf(" [%s Applied]", res.AppliedEffect.Type)
		}
	}
	m.logs = append(m.logs, logMsg)

	if m.dStats.HP <= 0 {
		m.dStats.HP = 0
		m.gameOver = true
		m.winner = "ATTACKER"
		m.logs = append(m.logs, ">>> DEFENDER DESTROYED! <<<")
		return m, nil
	}

	// Counter Attack
	resC := m.combatService.ExecuteAttack(m.dStats, m.aStats, combat.Kinetic)
	logMsgC := "Defender Counters: "
	if resC.IsMiss {
		logMsgC += "MISS!"
	} else {
		m.aStats.HP -= resC.FinalDamage
		logMsgC += fmt.Sprintf("HIT for %d damage!", resC.FinalDamage)
	}
	m.logs = append(m.logs, logMsgC)

	if m.aStats.HP <= 0 {
		m.aStats.HP = 0
		m.gameOver = true
		m.winner = "DEFENDER"
		m.logs = append(m.logs, ">>> ATTACKER DESTROYED! <<<")
	}

	// Keep only last 8 logs
	if len(m.logs) > 8 {
		m.logs = m.logs[len(m.logs)-8:]
	}

	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render(" PROJECT-0: TACTICAL NOIR TUI ") + "\n\n"

	// Attacker Stats
	aInfo := fmt.Sprintf("ATTACKER: %s\n", m.attacker.Class)
	aInfo += fmt.Sprintf("HP: %d/%d\n", m.aStats.HP, m.aStats.MaxHP)
	aInfo += m.aHPBar.ViewAs(float64(m.aStats.HP) / float64(m.aStats.MaxHP)) + "\n\n"
	aInfo += fmt.Sprintf("ATK: %d | DEF: %d\n", m.aStats.BaseAttack, m.aStats.TargetDefense)
	aInfo += fmt.Sprintf("SPD: %d | EVA: %d\n", m.aStats.Speed, m.aStats.Evasion)
	if m.aPilot != nil {
		aInfo += fmt.Sprintf("Resonance: Lvl %d\n", m.aPilot.ResonanceLevel)
		aInfo += fmt.Sprintf("O2: %.1f%% | Fuel: %.1f%%", m.aPilot.CurrentO2, m.aPilot.CurrentFuel)
	}

	// Defender Stats
	dInfo := fmt.Sprintf("DEFENDER: %s\n", m.defender.Class)
	dInfo += fmt.Sprintf("HP: %d/%d\n", m.dStats.HP, m.dStats.MaxHP)
	dInfo += m.dHPBar.ViewAs(float64(m.dStats.HP) / float64(m.dStats.MaxHP)) + "\n\n"
	dInfo += fmt.Sprintf("ATK: %d | DEF: %d\n", m.dStats.BaseAttack, m.dStats.TargetDefense)
	dInfo += fmt.Sprintf("SPD: %d | EVA: %d\n", m.dStats.Speed, m.dStats.Evasion)
	if m.dPilot != nil {
		dInfo += fmt.Sprintf("Resonance: Lvl %d\n", m.dPilot.ResonanceLevel)
		dInfo += fmt.Sprintf("O2: %.1f%% | Fuel: %.1f%%", m.dPilot.CurrentO2, m.dPilot.CurrentFuel)
	}

	cols := lipgloss.JoinHorizontal(lipgloss.Top, statsStyle.Render(aInfo), statsStyle.Render(dInfo))
	s += cols + "\n\n"

	// Logs
	logContent := strings.Join(m.logs, "\n")
	s += logStyle.Render(logContent) + "\n\n"

	if m.gameOver {
		s += lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true).Render(fmt.Sprintf(" BATTLE OVER - WINNER: %s (Press any key to quit)", m.winner))
	} else {
		s += "Commands: [1] Kinetic [2] Energy [3] Explosive | [Q] Quit"
	}

	return s
}

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

	p := tea.NewProgram(initialModel(db))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
