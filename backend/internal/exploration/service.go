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

type TerrainType string

const (
	TerrainUrban   TerrainType = "URBAN"
	TerrainIslands TerrainType = "ISLANDS"
	TerrainSky     TerrainType = "SKY"
	TerrainDesert  TerrainType = "DESERT"
	TerrainVoid    TerrainType = "VOID"
	TerrainSpace   TerrainType = "SPACE"
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
	Terrain                TerrainType       `json:"terrain"`
	DetectionThreshold     int               `json:"detection_threshold"`
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
	ID                 uuid.UUID   `json:"id"`
	SectorID           uuid.UUID   `json:"sector_id"`
	Type               string      `json:"type"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	Rewards            []string    `json:"rewards"`
	Requirements       []string    `json:"requirements"`
	AllowedModes       []string    `json:"allowed_modes"`
	RequiresAtmosphere bool        `json:"requires_atmosphere"`
	Terrain            TerrainType `json:"terrain"`
	DetectionThreshold int         `json:"detection_threshold"`
	SuitabilityPilot   int         `json:"suitability_pilot"`
	SuitabilityVehicle int         `json:"suitability_vehicle"`
	CoordinatesX       int         `json:"coordinates_x"`
	CoordinatesY       int         `json:"coordinates_y"`
}

type PlanetLocation struct {
	ID                 uuid.UUID   `json:"id"`
	SubSectorID        uuid.UUID   `json:"sub_sector_id"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	Rewards            []string    `json:"rewards"`
	Requirements       []string    `json:"requirements"`
	AllowedModes       []string    `json:"allowed_modes"`
	RequiresAtmosphere bool        `json:"requires_atmosphere"`
	Terrain            TerrainType `json:"terrain"`
	DetectionThreshold int         `json:"detection_threshold"`
	SuitabilityPilot   int         `json:"suitability_pilot"`
	SuitabilityVehicle int         `json:"suitability_vehicle"`
	CoordinatesX       int         `json:"coordinates_x"`
	CoordinatesY       int         `json:"coordinates_y"`
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
	ID                 uuid.UUID   `json:"id"`
	Type               NodeType    `json:"type"`
	Title              string      `json:"title"`
	Description        string      `json:"description"`
	VisualPrompt       string      `json:"visual_prompt"`
	EnemyID            *uuid.UUID  `json:"enemy_id,omitempty"`
	Terrain            TerrainType `json:"terrain"`
	DetectionThreshold int         `json:"detection_threshold"`
	CreatedAt          time.Time   `json:"created_at"`
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

	// 1.5 Ensure Pilot has resources to start
	pilot, _ := s.gameRepo.GetActivePilotStats(userID)
	if pilot != nil && (pilot.CurrentO2 < 20 || pilot.CurrentFuel < 10) {
		pilot.CurrentO2 = 100.0
		pilot.CurrentFuel = 100.0
		_ = s.gameRepo.UpdatePilotStats(pilot)
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
	terrains := []TerrainType{TerrainUrban, TerrainIslands, TerrainSky, TerrainDesert, TerrainVoid, TerrainSpace}

	for i := 0; i < length; i++ {
		nodeType := types[rand.Intn(len(types))]
		if i == length-1 {
			nodeType = NodeOutpost // Last node is always an outpost/boss
		}

		terrain := terrains[rand.Intn(len(terrains))]
		tmpl := GetTemplateForType(nodeType)

		nodes[i] = Node{
			ID:                     uuid.New(),
			ExpeditionID:           expeditionID,
			Name:                   tmpl.Name,
			Type:                   nodeType,
			EnvironmentDescription: tmpl.Description,
			DifficultyMultiplier:   1.0 + (float64(i) * 0.1),
			PositionIndex:          i,
			Choices:                tmpl.Choices,
			IsResolved:             false,
			Terrain:                terrain,
			DetectionThreshold:     500 + rand.Intn(500), // 500-1000
		}
	}
	return nodes
}

// CalculateEffectiveCP implements the blueprint formula:
// ECP = (Base_CP * Suitability_Mod) * Resonance_Sync * (1 - Fatigue_Penalty)
func (s *Service) CalculateEffectiveCP(ctx context.Context, userID uuid.UUID, vehicleID uuid.UUID, terrain TerrainType) (int, error) {
	// 1. Get Pilot Stats
	pilot, err := s.gameRepo.GetActivePilotStats(userID)
	if err != nil {
		return 0, err
	}

	// 2. Handle Pilot Only Mode
	if vehicleID == uuid.Nil {
		// Base Pilot CP is 50
		// Suitability is always 1.0 for pilot
		// Resonance Sync is always 1.0 for pilot
		fatiguePenalty := float64(pilot.Stress) / 200.0
		if fatiguePenalty > 0.5 {
			fatiguePenalty = 0.5
		}
		ecp := 50.0 * (1.0 - fatiguePenalty)
		return int(ecp), nil
	}

	// 3. Get Base CP
	baseCP, err := s.vehicleUseCase.GetVehicleCP(ctx, vehicleID)
	if err != nil {
		return 0, err
	}

	// 4. Get Vehicle for Suitability
	v, err := s.vehicleUseCase.GetVehicleByID(ctx, vehicleID)
	if err != nil {
		return 0, err
	}

	// 5. Calculate Suitability Modifier
	suitabilityMod := 1.0
	// Simple mapping for now: check if terrain is in suitability_tags
	isSuitable := false
	for _, tag := range v.SuitabilityTags {
		if string(terrain) == tag {
			isSuitable = true
			break
		}
	}

	if isSuitable {
		suitabilityMod = 1.2
	} else {
		// Check for incompatible types (e.g. Tank in Islands)
		if (v.VehicleType == TypeTank && terrain == TerrainIslands) ||
			(v.VehicleType == TypeShip && terrain == TerrainDesert) {
			suitabilityMod = 0.5
		}
	}

	// 6. Calculate Resonance Sync
	// Formula: Min(1.0, Pilot_Resonance / Vehicle_Tier_Requirement)
	// Assuming Tier 1 needs 20, Tier 2 needs 40, etc.
	tierReq := v.Tier * 20
	resonanceSync := 1.0
	if tierReq > 0 && pilot.ResonanceLevel < tierReq {
		resonanceSync = float64(pilot.ResonanceLevel) / float64(tierReq)
	}
	if resonanceSync < 0.1 {
		resonanceSync = 0.1 // Minimum sync
	}

	// 7. Calculate Fatigue Penalty
	// Formula: Stress / 200 (Max 50% penalty at 100 Stress)
	fatiguePenalty := float64(pilot.Stress) / 200.0
	if fatiguePenalty > 0.5 {
		fatiguePenalty = 0.5
	}

	// 8. Final ECP Calculation
	ecp := float64(baseCP) * suitabilityMod * resonanceSync * (1.0 - fatiguePenalty)

	return int(ecp), nil
}

func (s *Service) ResolveNode(ctx context.Context, nodeID uuid.UUID) (*Node, error) {
	node, err := s.repo.GetNodeByID(nodeID)
	if err != nil {
		return nil, err
	}
	node.IsResolved = true
	if err := s.repo.UpdateNode(node); err != nil {
		return nil, err
	}
	return node, nil
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

	// 4. Calculate Effective CP (ECP)
	vehicleID := uuid.Nil
	if expedition.VehicleID != nil {
		vehicleID = *expedition.VehicleID
	}

	ecp, err := s.CalculateEffectiveCP(ctx, expedition.UserID, vehicleID, node.Terrain)
	if err != nil {
		// Fallback if calculation fails
		ecp = 100
	}

	// Success Chance Adjustment based on ECP vs Node Difficulty
	// Base difficulty is 200 * DifficultyMultiplier
	baseDifficulty := 200.0 * node.DifficultyMultiplier

	// Success chance formula: 0.5 + (ECP - Difficulty) / 1000
	ecpBonus := (float64(ecp) - baseDifficulty) / 1000.0
	if ecpBonus < -0.3 {
		ecpBonus = -0.3 // Max penalty
	}
	if ecpBonus > 0.4 {
		ecpBonus = 0.4 // Max bonus
	}

	finalSuccessChance := selectedChoice.SuccessChance + ecpBonus
	if finalSuccessChance > 0.98 {
		finalSuccessChance = 0.98
	}
	if finalSuccessChance < 0.05 {
		finalSuccessChance = 0.05
	}

	success := rand.Float64() < finalSuccessChance

	// 5. Apply Consequences
	stats, err := s.gameRepo.GetActivePilotStats(expedition.UserID)
	if err == nil && stats != nil {
		// Every node resolution increases Stress
		stats.Stress += 5 + rand.Intn(5) // 5-10 stress per node
		if stats.Stress > 100 {
			stats.Stress = 100
		}

		if success {
			// Apply Rewards
			for _, reward := range selectedChoice.Rewards {
				switch reward {
				case "Scrap Metal":
					stats.ScrapMetal += 50
				case "Research Data":
					stats.ResearchData += 20
				case "Rare Ore":
					stats.ScrapMetal += 150
				}
			}
			// Always give some XP on success
			stats.XP += 15
		} else {
			// Apply Damage and Stress on failure
			stats.Stress += 10 // Extra stress on failure
			if stats.Stress > 100 {
				stats.Stress = 100
			}

			for _, risk := range selectedChoice.Risks {
				if risk == "High Damage" || risk == "Medium Damage" || risk == "Structural Stress" {
					if expedition.VehicleID != nil {
						damage := 15
						if risk == "High Damage" {
							damage = 30
						}
						// Apply damage to the vehicle (as an item)
						_, _ = s.vehicleUseCase.ApplyDamage(ctx, *expedition.VehicleID, damage)
					}
				}
			}
		}
		_ = s.gameRepo.UpdatePilotStats(stats)
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
	choices := []StrategicChoice{}
	switch t {
	case NodeCombat:
		choices = []StrategicChoice{
			{
				Label: "Full Assault", 
				Description: "Direct confrontation with maximum firepower.", 
				SuccessChance: 0.6, 
				Risks: []string{"High Damage"},
				Requirements: []string{},
				Rewards: []string{},
			},
			{
				Label: "Tactical Flank", 
				Description: "Use agility to find a weak spot.", 
				Requirements: []string{"PILOT_AGILITY > 40"}, 
				SuccessChance: 0.8, 
				Risks: []string{"Medium Damage"},
				Rewards: []string{},
			},
		}
	case NodeResource:
		choices = []StrategicChoice{
			{
				Label: "Deep Drill", 
				Description: "Extract rare minerals from the core.", 
				SuccessChance: 0.4, 
				Rewards: []string{"Rare Ore"}, 
				Risks: []string{"Structural Stress"},
				Requirements: []string{},
			},
			{
				Label: "Surface Scavenge", 
				Description: "Quickly gather loose materials.", 
				SuccessChance: 0.9, 
				Rewards: []string{"Scrap Metal"},
				Risks: []string{},
				Requirements: []string{},
			},
		}
	case NodeAnomaly:
		choices = []StrategicChoice{
			{
				Label: "Scientific Study", 
				Description: "Analyze the anomaly for data.", 
				Requirements: []string{"PILOT_INTEL > 50"}, 
				SuccessChance: 0.7, 
				Rewards: []string{"Research Data"},
				Risks: []string{},
			},
			{
				Label: "Brute Force", 
				Description: "Push through the anomaly.", 
				SuccessChance: 0.5, 
				Risks: []string{"System Glitch"},
				Requirements: []string{},
				Rewards: []string{},
			},
		}
	default:
		choices = []StrategicChoice{
			{
				Label: "Full Speed", 
				Description: "Travel quickly to the next sector.", 
				SuccessChance: 0.9, 
				Risks: []string{"Fuel Consumption"},
				Requirements: []string{},
				Rewards: []string{},
			},
			{
				Label: "Eco-Drive", 
				Description: "Conserve energy while moving.", 
				SuccessChance: 1.0,
				Risks: []string{},
				Requirements: []string{},
				Rewards: []string{},
			},
		}
	}
	return choices
}

// GenerateNewEncounter generates a new procedural event (Encounter) based on the current Expedition context
func (s *Service) GenerateNewEncounter(ctx context.Context, expeditionID uuid.UUID, vehicleID uuid.UUID) (*Encounter, error) {
	// 1. Fetch Context Data
	expedition, err := s.repo.GetExpeditionByID(expeditionID)
	if err != nil {
		return nil, err
	}

	// Get all nodes for this expedition
	nodes, err := s.repo.GetNodesByExpeditionID(expeditionID)
	if err != nil {
		return nil, err
	}

	// Get existing encounters to find current position
	existingEncounters, _ := s.repo.GetEncountersByExpeditionID(expeditionID)
	currentIndex := len(existingEncounters)

	if currentIndex >= len(nodes) {
		return nil, fmt.Errorf("expedition completed")
	}

	targetNode := nodes[currentIndex]

	// Anti-Cheat: Check if last encounter is resolved
	if currentIndex > 0 {
		last := existingEncounters[currentIndex-1]
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

	// 2. Determine Encounter Type from Timeline Node
	encounterType := targetNode.Type

	// Consume resources (only if advancing, not for the first encounter)
	if pilot != nil && currentIndex > 0 {
		success, err := s.gameRepo.ConsumeResources(pilot.CharacterID, 15.0, 5.0)
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, fmt.Errorf("insufficient resources to advance")
		}
	}

	// 3. Generate Narrative Context based on Node Type
	title := targetNode.Name
	desc := targetNode.EnvironmentDescription

	switch encounterType {
	case NodeCombat:
		title = "Hostile Signature Detected"
		desc = "An unknown unit is approaching with weapons hot."
	case NodeResource:
		title = "Resource Cluster"
		desc = "Scanners indicate a high concentration of valuable materials."
	case NodeAnomaly:
		title = "Spatial Anomaly"
		desc = "The fabric of space seems distorted in this region."
	case NodeOutpost:
		title = "Abandoned Outpost"
		desc = "A derelict structure floats silently in the void."
	}

	// 4. Generate Visual Prompt (DDS Integrated)
	prompt := s.GenerateVisualPrompt(item, &targetNode)

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
		ID:                 uuid.New(),
		Type:               encounterType,
		Title:              title,
		Description:        desc,
		VisualPrompt:       prompt,
		EnemyID:            enemyID,
		Terrain:            targetNode.Terrain,
		DetectionThreshold: targetNode.DetectionThreshold,
		CreatedAt:          time.Now(),
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
