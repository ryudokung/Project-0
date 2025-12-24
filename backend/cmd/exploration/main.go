package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ryudokung/Project-0/backend/internal/exploration"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"github.com/ryudokung/Project-0/backend/internal/mech"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#F25D94")).
			Padding(0, 1).
			Bold(true)

	encounterStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#383838")).
			Padding(1, 2).
			MarginBottom(1)

	promptStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7571F9")).
			Italic(true)
)

type model struct {
	service    *exploration.Service
	expedition *exploration.Expedition
	encounters []exploration.Encounter
	mechID     uuid.UUID
	err        error
}

func initialModel(db *sql.DB) model {
	repo := exploration.NewRepository(db)
	mechRepo := mech.NewRepository(db)
	gameRepo := game.NewRepository(db)
	service := exploration.NewService(repo, mechRepo, gameRepo)

	// Use the sample expedition ID from init.sql
	expeditionID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	expedition, _ := repo.GetExpeditionByID(expeditionID)
	encounters, _ := repo.GetEncountersByExpeditionID(expeditionID)

	// Fetch a sample mech ID
	var mechID uuid.UUID
	db.QueryRow("SELECT id FROM mechs LIMIT 1").Scan(&mechID)

	return model{
		service:    service,
		expedition: expedition,
		encounters: encounters,
		mechID:     mechID,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "n":
			// String a new encounter
			ctx := context.Background()
			encounter, err := m.service.StringNewEncounter(ctx, m.expedition.ID, m.mechID)
			if err != nil {
				m.err = err
				return m, nil
			}
			m.encounters = append(m.encounters, *encounter)
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\nPress 'q' to quit.", m.err)
	}

	var s strings.Builder
	s.WriteString(titleStyle.Render("PROJECT-0: NARRATIVE TIMELINE (EXPEDITION & ENCOUNTERS)"))
	s.WriteString("\n\n")

	if m.expedition != nil {
		s.WriteString(fmt.Sprintf("EXPEDITION: %s\n", m.expedition.Title))
		s.WriteString(fmt.Sprintf("GOAL: %s\n\n", m.expedition.Goal))
	}

	s.WriteString("TIMELINE:\n")
	for i, b := range m.encounters {
		content := fmt.Sprintf("[%d] %s (%s)\n%s\n\n", i+1, b.Title, b.Type, b.Description)
		content += promptStyle.Render("Visual DNA: " + b.VisualPrompt)
		s.WriteString(encounterStyle.Render(content))
	}

	s.WriteString("\nPress 'n' to string a new encounter, 'q' to quit.\n")

	return s.String()
}

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgresql://user:password@localhost:5432/project0?sslmode=disable"
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
