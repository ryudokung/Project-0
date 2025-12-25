package exploration

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"github.com/ryudokung/Project-0/backend/internal/vehicle"
)

type NodeType string

const (
	NodeStandard  NodeType = "STANDARD"
	NodeResource  NodeType = "RESOURCE"
	NodeCombat    NodeType = "COMBAT"
	NodeAnomaly   NodeType = "ANOMALY"
	NodeOutpost   NodeType = "OUTPOST"
)

type ApproachType string

const (
	ApproachPassive  ApproachType = "PASSIVE_SCAN"
	ApproachDeep     ApproachType = "DEEP_ANALYSIS"
	ApproachStealth  ApproachType = "STEALTH"
)

type StrategicChoice struct {
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Requirements []string `json:"requirements"` // e.g. "PILOT_AGILITY > 50"
	SuccessChance float64 `json:"success_chance"`
	Rewards      []string `json:"rewards"`
	Risks        []string `json:"risks"`
}

type Node struct {
	ID                     uuid.UUID         `json:"id"`
	ExpeditionID           uuid.UUID         `json:"expedition_id"`
	Name                   string            `json:"name"`
	Type                   NodeType          `json:"type"`
	EnvironmentDescription string            `json:"environment_description"`
	DifficultyMultiplier   float64           `json:"difficulty_multiplier"`
	PositionIndex          int               `json:"position_index"`
	Choices                []StrategicChoice `json:"choices"`
	IsResolved             bool              `json:"is_resolved"`
}

type Session struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	VehicleID     uuid.UUID `json:"vehicle_id"`
	CurrentNodeID uuid.UUID `json:"current_node_id"`
	Status        string    `json:"status"`
}

type Repository interface {
	CreateExpedition(expedition *Expedition) error
	GetExpeditionByID(id uuid.UUID) (*Expedition, error)
	
	// Timeline Nodes
	CreateNodes(nodes []Node) error
	GetNodesByExpeditionID(expeditionID uuid.UUID) ([]Node, error)
	GetNodeByID(id uuid.UUID) (*Node, error)
	UpdateNode(node *Node) error
	
	// Legacy/Other
	SaveEncounter(encounter *Encounter, expeditionID uuid.UUID) error
	GetEncountersByExpeditionID(expeditionID uuid.UUID) ([]Encounter, error)
	GetSessionByUserID(userID uuid.UUID) (*Session, error)
	GetAllSectors() ([]Sector, error)
	GetSubSectorsBySectorID(sectorID uuid.UUID) ([]SubSector, error)
	GetPlanetLocationsBySubSectorID(subSectorID uuid.UUID) ([]PlanetLocation, error)
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
	SuitabilityVehicle int       `json:"suitability_vehicle"`
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
	SuitabilityVehicle int       `json:"suitability_vehicle"`
	CoordinatesX       int       `json:"coordinates_x"`
	CoordinatesY       int       `json:"coordinates_y"`
}

type Expedition struct {
	ID               uuid.UUID  `json:"id"`
	UserID           uuid.UUID  `json:"user_id"`
	SubSectorID      uuid.UUID  `json:"sub_sector_id"`
	PlanetLocationID *uuid.UUID `json:"planet_location_id"`
	VehicleID        *uuid.UUID `json:"vehicle_id"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	Goal             string     `json:"goal"`
}

type Encounter struct {
	ID           uuid.UUID  `json:"id"`
	Type         NodeType   `json:"type"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	VisualPrompt string     `json:"visual_prompt"`
	EnemyID      *uuid.UUID `json:"enemy_id,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}

type Service struct {
	repo           Repository
	vehicleUseCase vehicle.UseCase
	gameRepo       game.Repository
}

func NewService(repo Repository, vehicleUseCase vehicle.UseCase, gameRepo game.Repository) *Service {
	return &Service{repo: repo, vehicleUseCase: vehicleUseCase, gameRepo: gameRepo}
}

func (s *Service) StartExploration(ctx context.Context, userID uuid.UUID, subSectorID uuid.UUID, planetLocationID *uuid.UUID, vehicleID uuid.UUID) (*Expedition, error) {
	// 0. Verify Vehicle Ownership (if vehicle is provided)
	var vID *uuid.UUID
	if vehicleID != uuid.Nil {
		// Check items table first (Unified System)
		item, err := s.vehicleUseCase.GetItemByID(ctx, vehicleID)
		if err != nil {
			return nil, fmt.Errorf("error fetching vehicle: %v", err)
		}
		
		if item == nil {
			// Fallback to vehicles table
			v, err := s.vehicleUseCase.GetVehicleByID(ctx, vehicleID)
			if err != nil {
				return nil, fmt.Errorf("error fetching vehicle: %v", err)
			}
			if v == nil {
				return nil, fmt.Errorf("vehicle not found")
			}
			if v.OwnerID != userID {
				return nil, fmt.Errorf("unauthorized: you do not own this vehicle")
			}
		} else {
			if item.OwnerID != userID {
				return nil, fmt.Errorf("unauthorized: you do not own this vehicle")
			}
		}
		vID = &vehicleID
	}

	// 1. Create Expedition
	expedition := &Expedition{
		ID:               uuid.New(),
		UserID:           userID,
		SubSectorID:      subSectorID,
		PlanetLocationID: planetLocationID,
		VehicleID:        vID,
		Title:            "The Silent Signal",
		Description:      "Investigating a mysterious signal in the sector.",
		Goal:             "Locate the source of the signal.",
	}

	if err := s.repo.CreateExpedition(expedition); err != nil {
		return nil, err
	}

	// 2. Generate Timeline (5-7 nodes)
	nodes := s.GenerateTimeline(expedition.ID, 6)
	if err := s.repo.CreateNodes(nodes); err != nil {
		return nil, err
	}

	// 3. Generate First Encounter
	_, err := s.GenerateNewEncounter(ctx, expedition.ID, vehicleID)
	if err != nil {
		// Log error but don't fail start (we can generate it later if needed)
		fmt.Printf("Warning: failed to generate first encounter: %v\n", err)
	}

	return expedition, nil
}

func (s *Service) GenerateTimeline(expeditionID uuid.UUID, length int) []Node {
	nodes := make([]Node, length)
	types := []NodeType{NodeStandard, NodeResource, NodeCombat, NodeAnomaly}

	for i := 0; i < length; i++ {
		nodeType := types[rand.Intn(len(types))]
		if i == length-1 {
			nodeType = NodeOutpost // Last node is always an outpost/boss
		}

		nodes[i] = Node{
			ID:                     uuid.New(),
			ExpeditionID:           expeditionID,
			Name:                   fmt.Sprintf("Sector %d", i+1),
			Type:                   nodeType,
			EnvironmentDescription: "Deep space void with flickering lights.",
			DifficultyMultiplier:   1.0 + (float64(i) * 0.1),
			PositionIndex:          i,
			Choices:                s.GenerateChoicesForType(nodeType),
			IsResolved:             false,
		}
	}
	return nodes
}

func (s *Service) ResolveNodeChoice(ctx context.Context, nodeID uuid.UUID, choiceLabel string) (*Node, error) {
	// 1. Fetch Node
	node, err := s.repo.GetNodeByID(nodeID)
	if err != nil {
		return nil, err
	}
	if node.IsResolved {
		return nil, fmt.Errorf("node already resolved")
	}

	// 2. Find Choice
	var selectedChoice *StrategicChoice
	for _, c := range node.Choices {
		if c.Label == choiceLabel {
			selectedChoice = &c
			break
		}
	}
	if selectedChoice == nil {
		return nil, fmt.Errorf("choice not found")
	}

	// 3. Fetch Expedition & Vehicle
	expedition, err := s.repo.GetExpeditionByID(node.ExpeditionID)
	if err != nil {
		return nil, err
	}

	// 4. Calculate Success
	success := rand.Float64() < selectedChoice.SuccessChance
	
	// 5. Apply Consequences
	if !success {
		// Apply Damage if there are risks
		for _, risk := range selectedChoice.Risks {
			if risk == "High Damage" || risk == "Medium Damage" || risk == "Structural Stress" {
				if expedition.VehicleID != nil {
					damage := 10
					if risk == "High Damage" {
						damage = 25
					}
					// Apply damage to the vehicle (as an item)
					_, _ = s.vehicleUseCase.ApplyDamage(ctx, *expedition.VehicleID, damage)
				}
			}
		}
	}

	// 6. Update Node
	node.IsResolved = true
	// We could store the result in metadata if we want
	if err := s.repo.UpdateNode(node); err != nil {
		return nil, err
	}

	return node, nil
}

func (s *Service) GenerateChoicesForType(t NodeType) []StrategicChoice {
	switch t {
	case NodeCombat:
		return []StrategicChoice{
			{Label: "Full Assault", Description: "Direct confrontation with maximum firepower.", SuccessChance: 0.6, Risks: []string{"High Damage"}},
			{Label: "Tactical Flank", Description: "Use agility to find a weak spot.", Requirements: []string{"PILOT_AGILITY > 40"}, SuccessChance: 0.8, Risks: []string{"Medium Damage"}},
		}
	case NodeResource:
		return []StrategicChoice{
			{Label: "Deep Drill", Description: "Extract rare minerals from the core.", SuccessChance: 0.4, Rewards: []string{"Rare Ore"}, Risks: []string{"Structural Stress"}},
			{Label: "Surface Scavenge", Description: "Quickly gather loose materials.", SuccessChance: 0.9, Rewards: []string{"Scrap Metal"}},
		}
	case NodeAnomaly:
		return []StrategicChoice{
			{Label: "Scientific Study", Description: "Analyze the anomaly for data.", Requirements: []string{"PILOT_INTEL > 50"}, SuccessChance: 0.7, Rewards: []string{"Research Data"}},
			{Label: "Brute Force", Description: "Push through the anomaly.", SuccessChance: 0.5, Risks: []string{"System Glitch"}},
		}
	default:
		return []StrategicChoice{
			{Label: "Full Speed", Description: "Travel quickly to the next sector.", SuccessChance: 0.9, Risks: []string{"Fuel Consumption"}},
			{Label: "Eco-Drive", Description: "Conserve energy while moving.", SuccessChance: 1.0},
		}
	}
}

// GenerateNewEncounter generates a new procedural event (Encounter) based on the current Expedition context
func (s *Service) GenerateNewEncounter(ctx context.Context, expeditionID uuid.UUID, vehicleID uuid.UUID) (*Encounter, error) {
	// 1. Fetch Context Data
	expedition, err := s.repo.GetExpeditionByID(expeditionID)
	if err != nil {
		return nil, err
	}

	// Anti-Cheat: Check if last encounter is resolved
	lastEncounters, _ := s.repo.GetEncountersByExpeditionID(expeditionID)
	if len(lastEncounters) > 0 {
		last := lastEncounters[len(lastEncounters)-1]
		if last.Type == NodeCombat && last.EnemyID != nil {
			// Check enemy status (using Item system if possible)
			enemy, _ := s.vehicleUseCase.GetItemByID(ctx, *last.EnemyID)
			if enemy != nil && enemy.Durability > 0 {
				return nil, fmt.Errorf("current combat encounter not resolved")
			}
		}
	}

	// Fetch Vehicle (Unified Item System)
	var item *vehicle.Item
	if vehicleID != uuid.Nil {
		item, err = s.vehicleUseCase.GetItemByID(ctx, vehicleID)
		if err != nil {
			return nil, err
		}

		// Fallback to vehicles table if not found in items
		if item == nil {
			v, _ := s.vehicleUseCase.GetVehicleByID(ctx, vehicleID)
			if v != nil {
				// Map Vehicle to Item structure for compatibility
				item = &vehicle.Item{
					ID:        v.ID,
					Name:      string(v.Class),
					Rarity:    v.Rarity,
					Condition: vehicle.ConditionPristine,
					Stats: vehicle.ItemStats{
						HP:      v.Stats.HP,
						Attack:  v.Stats.Attack,
						Defense: v.Stats.Defense,
						Speed:   v.Stats.Speed,
					},
				}
			}
		}
	}

	// If vehicleID is zero UUID, item will be nil (Pilot Only mode)
	pilot, err := s.gameRepo.GetActivePilotStats(expedition.UserID)
	if err != nil {
		return nil, err
	}

	// Anti-Cheat: Check Resources
	if pilot != nil && pilot.CurrentO2 <= 0 {
		return nil, fmt.Errorf("insufficient oxygen to advance")
	}

	// 2. Determine Encounter Type
	encounterType := NodeCombat
	if pilot != nil {
		success, err := s.gameRepo.ConsumeResources(pilot.CharacterID, 15.0, 5.0)
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, fmt.Errorf("insufficient resources to advance")
		}

		if pilot.CurrentO2 < 30 {
			encounterType = NodeResource
		}
	}

	// 3. Generate Narrative Context
	title := ""
	desc := ""
	env := ""

	switch expedition.Title {
	case "The Silent Signal":
		if encounterType == NodeCombat {
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

	// 4. Generate Visual Prompt (DDS Integrated)
	node := &Node{EnvironmentDescription: env}
	prompt := s.GenerateVisualPrompt(item, node)

	var enemyID *uuid.UUID
	if encounterType == NodeCombat {
		enemies := []string{
			"e0000000-0000-0000-0000-000000000001",
			"e0000000-0000-0000-0000-000000000002",
			"e0000000-0000-0000-0000-000000000003",
		}
		selected := enemies[rand.Intn(len(enemies))]
		uid := uuid.MustParse(selected)
		enemyID = &uid
	}

	encounter := &Encounter{
		ID:           uuid.New(),
		Type:         encounterType,
		Title:        title,
		Description:  desc,
		VisualPrompt: prompt,
		EnemyID:      enemyID,
	}

	// 5. Save to Repository
	if err := s.repo.SaveEncounter(encounter, expeditionID); err != nil {
		return nil, err
	}

	return encounter, nil
}

// GenerateVisualPrompt combines Item DNA and Node Environment for AI Image Generation (DDS Integrated)
func (s *Service) GenerateVisualPrompt(item *vehicle.Item, node *Node) string {
	var dnaKeywords []string

	// 1. Default description for Pilot Only mode
	vehicleDesc := "Pilot in EVA suit"

	if item != nil {
		// 2. Collect DNA Keywords
		dnaKeywords = append(dnaKeywords, item.VisualDNA.Keywords...)

		// 3. Add Condition-based descriptors (Deep Durability System)
		conditionDesc := ""
		switch item.Condition {
		case vehicle.ConditionWorn:
			conditionDesc = "slightly worn"
		case vehicle.ConditionDamaged:
			conditionDesc = "damaged with visible smoke"
		case vehicle.ConditionCritical:
			conditionDesc = "critically damaged, sparking and glitching"
		case vehicle.ConditionBroken:
			conditionDesc = "completely broken and non-functional"
		default:
			conditionDesc = "pristine"
		}

		// 4. Build detailed description
		vehicleDesc = fmt.Sprintf("a %s %s %s with %s features",
			conditionDesc,
			item.Rarity,
			item.Name,
			strings.Join(dnaKeywords, ", "),
		)

		// 5. Add Visual DNA specific effects from DDS
		if item.VisualDNA.SmokeLevel > 0.5 {
			vehicleDesc += ", heavy smoke billowing"
		}
		if item.VisualDNA.SparksEnabled {
			vehicleDesc += ", electrical sparks flying"
		}
		if item.VisualDNA.GlitchIntensity > 0.5 {
			vehicleDesc += ", holographic glitch effects"
		}
	}

	prompt := fmt.Sprintf("Tactical Noir style, %s, standing in a %s environment, cinematic lighting, high detail",
		vehicleDesc,
		node.EnvironmentDescription,
	)

	return prompt
}
