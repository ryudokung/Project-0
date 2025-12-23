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

	beadStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#383838")).
			Padding(1, 2).
			MarginBottom(1)

	promptStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7571F9")).
			Italic(true)
)

type model struct {
	service *exploration.Service
	thread  *exploration.Thread
	beads   []exploration.Bead
	mechID  uuid.UUID
	err     error
}

func initialModel(db *sql.DB) model {
	repo := exploration.NewRepository(db)
	mechRepo := mech.NewRepository(db)
	gameRepo := game.NewRepository(db)
	service := exploration.NewService(repo, mechRepo, gameRepo)

	// Use the sample thread ID from init.sql
	threadID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	thread, _ := repo.GetThreadByID(threadID)
	beads, _ := repo.GetBeadsByThreadID(threadID)

	// Fetch a sample mech ID
	var mechID uuid.UUID
	db.QueryRow("SELECT id FROM mechs LIMIT 1").Scan(&mechID)

	return model{
		service: service,
		thread:  thread,
		beads:   beads,
		mechID:  mechID,
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
			// String a new bead
			ctx := context.Background()
			bead, err := m.service.StringNewBead(ctx, m.thread.ID, m.mechID)
			if err != nil {
				m.err = err
				return m, nil
			}
			m.beads = append(m.beads, *bead)
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\nPress 'q' to quit.", m.err)
	}

	var s strings.Builder
	s.WriteString(titleStyle.Render("PROJECT-0: NARRATIVE TIMELINE (THREAD & BEADS)"))
	s.WriteString("\n\n")

	if m.thread != nil {
		s.WriteString(fmt.Sprintf("THREAD: %s\n", m.thread.Title))
		s.WriteString(fmt.Sprintf("GOAL: %s\n\n", m.thread.Goal))
	}

	s.WriteString("TIMELINE:\n")
	for i, b := range m.beads {
		content := fmt.Sprintf("[%d] %s (%s)\n%s\n\n", i+1, b.Title, b.Type, b.Description)
		content += promptStyle.Render("Visual DNA: " + b.VisualPrompt)
		s.WriteString(beadStyle.Render(content))
	}

	s.WriteString("\nPress 'n' to string a new bead, 'q' to quit.\n")

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
