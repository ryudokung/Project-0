package exploration

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"github.com/ryudokung/Project-0/backend/internal/mech"
)

type NodeType string

const (
	NodeCombat    NodeType = "COMBAT"
	NodeResource  NodeType = "RESOURCE"
	NodeNarrative NodeType = "NARRATIVE"
	NodeRest      NodeType = "REST"
	NodeBoss      NodeType = "BOSS"
)

type Node struct {
	ID                     uuid.UUID `json:"id"`
	StarID                 uuid.UUID `json:"star_id"`
	Name                   string    `json:"name"`
	Type                   NodeType  `json:"type"`
	EnvironmentDescription string    `json:"environment_description"`
	DifficultyMultiplier   float64   `json:"difficulty_multiplier"`
	PositionIndex          int       `json:"position_index"`
}

type Session struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	MechID        uuid.UUID `json:"mech_id"`
	CurrentNodeID uuid.UUID `json:"current_node_id"`
	Status        string    `json:"status"`
}

type Repository interface {
	GetNodesByStarID(starID uuid.UUID) ([]Node, error)
	CreateSession(session *Session) error
	UpdateSession(session *Session) error
	GetSessionByUserID(userID uuid.UUID) (*Session, error)

	// Universe Map
	GetAllSectors() ([]Sector, error)
	GetSubSectorsBySectorID(sectorID uuid.UUID) ([]SubSector, error)
	GetPlanetLocationsBySubSectorID(subSectorID uuid.UUID) ([]PlanetLocation, error)

	// Thread & Bead operations
	CreateThread(thread *Thread) error
	GetThreadByID(id uuid.UUID) (*Thread, error)
	SaveBead(bead *Bead, threadID uuid.UUID) error
	GetBeadsByThreadID(threadID uuid.UUID) ([]Bead, error)
}

type Sector struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Difficulty   string    `json:"difficulty"`
	CoordinatesX int       `json:"coordinates_x"`
	CoordinatesY int       `json:"coordinates_y"`
	Color        string    `json:"color"`
}

type SubSector struct {
	ID                 uuid.UUID `json:"id"`
	SectorID           uuid.UUID `json:"sector_id"`
	Type               string    `json:"type"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Rewards            []string  `json:"rewards"`
	Requirements       []string  `json:"requirements"`
	AllowedModes       []string  `json:"allowed_modes"`
	RequiresAtmosphere bool      `json:"requires_atmosphere"`
	SuitabilityPilot   int       `json:"suitability_pilot"`
	SuitabilityMech    int       `json:"suitability_mech"`
	CoordinatesX       int       `json:"coordinates_x"`
	CoordinatesY       int       `json:"coordinates_y"`
}

type PlanetLocation struct {
	ID                 uuid.UUID `json:"id"`
	SubSectorID        uuid.UUID `json:"sub_sector_id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Rewards            []string  `json:"rewards"`
	Requirements       []string  `json:"requirements"`
	AllowedModes       []string  `json:"allowed_modes"`
	RequiresAtmosphere bool      `json:"requires_atmosphere"`
	SuitabilityPilot   int       `json:"suitability_pilot"`
	SuitabilityMech    int       `json:"suitability_mech"`
	CoordinatesX       int       `json:"coordinates_x"`
	CoordinatesY       int       `json:"coordinates_y"`
}

type Thread struct {
	ID               uuid.UUID  `json:"id"`
	UserID           uuid.UUID  `json:"user_id"`
	SubSectorID      uuid.UUID  `json:"sub_sector_id"`
	PlanetLocationID *uuid.UUID `json:"planet_location_id"`
	VehicleID        uuid.UUID  `json:"vehicle_id"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	Goal             string     `json:"goal"`
}

type Bead struct {
	ID          uuid.UUID `json:"id"`
	Type        NodeType  `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	VisualPrompt string    `json:"visual_prompt"`
}

type Service struct {
	repo     Repository
	mechRepo mech.Repository
	gameRepo game.Repository
}

func NewService(repo Repository, mechRepo mech.Repository, gameRepo game.Repository) *Service {
	return &Service{repo: repo, mechRepo: mechRepo, gameRepo: gameRepo}
}

func (s *Service) StartExploration(ctx context.Context, userID uuid.UUID, subSectorID uuid.UUID, planetLocationID *uuid.UUID, vehicleID uuid.UUID) (*Thread, error) {
	// 1. Create Thread
	thread := &Thread{
		ID:               uuid.New(),
		UserID:           userID,
		SubSectorID:      subSectorID,
		PlanetLocationID: planetLocationID,
		VehicleID:        vehicleID,
		Title:            "The Silent Signal",
		Description:      "Investigating a mysterious signal in the sector.",
		Goal:             "Locate the source of the signal.",
	}

	if err := s.repo.CreateThread(thread); err != nil {
		return nil, err
	}

	// 2. Generate First Bead
	_, err := s.StringNewBead(ctx, thread.ID, vehicleID)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

// StringNewBead generates a new procedural event (Bead) based on the current Thread context
func (s *Service) StringNewBead(ctx context.Context, threadID uuid.UUID, mechID uuid.UUID) (*Bead, error) {
	// 1. Fetch Context Data
	thread, err := s.repo.GetThreadByID(threadID)
	if err != nil {
		return nil, err
	}
	m, err := s.mechRepo.GetByID(ctx, mechID)
	if err != nil {
		return nil, err
	}
	parts, err := s.mechRepo.GetPartsByMechID(mechID)
	if err != nil {
		return nil, err
	}
	pilot, err := s.gameRepo.GetPilotStats(m.OwnerID)
	if err != nil {
		return nil, err
	}

	// 2. Determine Bead Type based on Thread and Pilot Stats
	beadType := NodeCombat
	if pilot != nil {
		// Consume Resources
		pilot.CurrentO2 -= 15.0
		pilot.CurrentFuel -= 5.0
		if pilot.CurrentO2 < 0 {
			pilot.CurrentO2 = 0
		}
		if pilot.CurrentFuel < 0 {
			pilot.CurrentFuel = 0
		}
		
		if err := s.gameRepo.UpdatePilotStats(pilot); err != nil {
			return nil, err
		}

		if pilot.CurrentO2 < 30 {
			beadType = NodeResource
		}
	}

	// 3. Generate Narrative Context
	title := ""
	desc := ""
	env := ""

	switch thread.Title {
	case "The Silent Signal":
		if beadType == NodeCombat {
			title = "Scavenger Ambush"
			desc = "A group of Iron Syndicate scavengers spotted your repair signal."
			env = "Electromagnetic Storm, Rusted Satellite Debris"
		} else {
			title = "Signal Echo"
			desc = "You found an old data log while scanning for parts."
			env = "Quiet Void, Flickering Radar Screen"
		}
	default:
		title = "Unknown Encounter"
		desc = "Something emerges from the dark void."
		env = "Deep Space, Neon Fog"
	}

	// 4. Generate Visual Prompt
	node := &Node{EnvironmentDescription: env}
	prompt := s.GenerateVisualPrompt(m, parts, node)

	bead := &Bead{
		ID:           uuid.New(),
		Type:         beadType,
		Title:        title,
		Description:  desc,
		VisualPrompt: prompt,
	}

	// 5. Save to Repository
	if err := s.repo.SaveBead(bead, threadID); err != nil {
		return nil, err
	}

	return bead, nil
}

// GenerateVisualPrompt combines Mech DNA and Node Environment for AI Image Generation
func (s *Service) GenerateVisualPrompt(m *mech.Mech, parts []mech.Part, node *Node) string {
	var dnaKeywords []string
	
	// 1. Collect Mech DNA
	for _, p := range parts {
		for _, k := range p.VisualDNA.Keywords {
			dnaKeywords = append(dnaKeywords, k)
		}
	}

	// 2. Combine with Node Environment
	prompt := fmt.Sprintf("Tactical Noir style, a %s mech with %s features, standing in a %s environment, cinematic lighting, high detail",
		m.Class,
		strings.Join(dnaKeywords, ", "),
		node.EnvironmentDescription,
	)

	return prompt
}
