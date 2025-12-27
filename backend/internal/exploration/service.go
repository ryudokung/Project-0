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
	NodeNarrative NodeType = "NARRATIVE"
)

type ZoneType string

const (
	ZoneOrbital  ZoneType = "ORBITAL"
	ZoneSurface  ZoneType = "SURFACE"
	ZoneEVA      ZoneType = "EVA"
	ZoneCorridor ZoneType = "CORRIDOR"
)

type HazardType string

const (
	HazardNone          HazardType = "NONE"
	HazardEMPStorm      HazardType = "EMP_STORM"
	HazardCorrosiveRain HazardType = "CORROSIVE_RAIN"
	HazardSolarFlare    HazardType = "SOLAR_FLARE"
	HazardVoidEcho      HazardType = "VOID_ECHO"
)

type ApproachType string

const (
	ApproachPassive  ApproachType = "PASSIVE_SCAN"
	ApproachDeep     ApproachType = "DEEP_ANALYSIS"
	ApproachStealth  ApproachType = "STEALTH"
)

type TerrainType string

const (
	TerrainIndustrial TerrainType = "INDUSTRIAL"
	TerrainMining     TerrainType = "MINING"
	TerrainCyber      TerrainType = "CYBER"
	TerrainAncient    TerrainType = "ANCIENT"
	TerrainUrban      TerrainType = "URBAN"
	TerrainIslands    TerrainType = "ISLANDS"
	TerrainSky        TerrainType = "SKY"
	TerrainDesert     TerrainType = "DESERT"
	TerrainVoid       TerrainType = "VOID"
	TerrainSpace      TerrainType = "SPACE"
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
	BlueprintID            string            `json:"blueprint_id"`
	Name                   string            `json:"name"`
	Type                   NodeType          `json:"type"`
	Zone                   ZoneType          `json:"zone"`
	Hazard                 HazardType        `json:"hazard"`
	EnvironmentDescription string            `json:"environment_description"`
	DifficultyMultiplier   float64           `json:"difficulty_multiplier"`
	PositionIndex          int               `json:"position_index"`
	Choices                []StrategicChoice `json:"choices"`
	IsResolved             bool              `json:"is_resolved"`
	Terrain                TerrainType       `json:"terrain"`
	DetectionThreshold     int               `json:"detection_threshold"`
	NextNodes              []string          `json:"next_nodes,omitempty"`
	IsScripted             bool              `json:"is_scripted"`
	ScriptEvents           []game.ScriptEvent `json:"script_events,omitempty"`
	IsEnd                  bool              `json:"is_end"`
	EnemyBlueprint         string            `json:"enemy_blueprint,omitempty"`
	EnemyCount             int               `json:"enemy_count,omitempty"`
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
	Status           string     `json:"status"`
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
	blueprints     *game.BlueprintRegistry
}

func NewService(repo Repository, vehicleUseCase vehicle.UseCase, gameRepo game.Repository, blueprints *game.BlueprintRegistry) *Service {
	return &Service{
		repo:           repo,
		vehicleUseCase: vehicleUseCase,
		gameRepo:       gameRepo,
		blueprints:     blueprints,
	}
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
	radarLevel := 1
	if pilot != nil && pilot.Metadata != nil {
		if lv, ok := pilot.Metadata["radar_level"].(float64); ok {
			radarLevel = int(lv)
		}
	}
	nodes := s.GenerateTimeline(expedition.ID, 6, radarLevel)
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

func (s *Service) CreateHandcraftedExpedition(ctx context.Context, userID uuid.UUID, blueprintID string, vehicleID *uuid.UUID) (*Expedition, error) {
	blueprint, ok := s.blueprints.Expeditions[blueprintID]
	if !ok {
		return nil, fmt.Errorf("expedition blueprint %s not found", blueprintID)
	}

	// 1. Create Expedition
	expedition := &Expedition{
		ID:          uuid.New(),
		UserID:      userID,
		VehicleID:   vehicleID,
		Title:       blueprint.Title,
		Description: blueprint.Description,
		Goal:        "Complete the mission objectives.",
		Status:      "ACTIVE",
	}

	if err := s.repo.CreateExpedition(expedition); err != nil {
		return nil, err
	}

	// 2. Create Nodes from Blueprint
	var nodes []Node
	for i, nb := range blueprint.Nodes {
		node := Node{
			ID:                     uuid.New(),
			ExpeditionID:           expedition.ID,
			BlueprintID:            nb.ID,
			Name:                   nb.Title,
			Type:                   NodeType(nb.Type),
			EnvironmentDescription: nb.Description,
			PositionIndex:          i,
			IsResolved:             false,
			NextNodes:              nb.NextNodes,
			IsScripted:             nb.IsScripted,
			ScriptEvents:           nb.ScriptEvents,
			IsEnd:                  nb.IsEnd,
			EnemyBlueprint:         nb.EnemyBlueprint,
			EnemyCount:             nb.EnemyCount,
		}
		nodes = append(nodes, node)
	}

	if err := s.repo.CreateNodes(nodes); err != nil {
		return nil, err
	}

	return expedition, nil
}

func (s *Service) GenerateTimeline(expeditionID uuid.UUID, length int, radarLevel int) []Node {
	nodes := make([]Node, length)
	types := []NodeType{NodeStandard, NodeResource, NodeCombat, NodeAnomaly, NodeNarrative}
	terrains := []TerrainType{TerrainIndustrial, TerrainMining, TerrainCyber, TerrainAncient}
	hazards := []HazardType{HazardNone, HazardEMPStorm, HazardCorrosiveRain, HazardSolarFlare, HazardVoidEcho}

	for i := 0; i < length; i++ {
		nodeType := types[rand.Intn(len(types))]
		zone := ZoneSurface
		
		// Logic for Zone distribution
		if i == 0 {
			zone = ZoneOrbital // Start with Orbital
			nodeType = NodeStandard
		} else if i == length-1 {
			zone = ZoneSurface
			nodeType = NodeOutpost // End with Outpost
		} else {
			// Randomly assign EVA or Corridor for variety
			r := rand.Float64()
			if r < 0.2 {
				zone = ZoneEVA
			} else if r < 0.4 {
				zone = ZoneCorridor
			}
		}

		terrain := terrains[rand.Intn(len(terrains))]
		hazard := HazardNone
		if rand.Float64() < 0.3 { // 30% chance of hazard
			hazard = hazards[1+rand.Intn(len(hazards)-1)]
		}

		// Find matching blueprint
		var blueprint game.NodeBlueprint
		found := false
		for _, b := range s.blueprints.Nodes {
			if b.Type == string(nodeType) {
				blueprint = b
				found = true
				break
			}
		}

		if !found {
			// Fallback
			blueprint = game.NodeBlueprint{
				ID:          "UNKNOWN",
				Name:        "Unknown Sector",
				Description: "A mysterious area of space.",
				Choices: []game.ChoiceBlueprint{
					{Label: "Proceed", Description: "Move forward carefully.", SuccessChance: 1.0},
				},
			}
		}

		// Use zone from blueprint if available
		if blueprint.Zone != "" {
			zone = ZoneType(blueprint.Zone)
		}

		// Map blueprint choices to StrategicChoice
		choices := make([]StrategicChoice, len(blueprint.Choices))
		for j, c := range blueprint.Choices {
			choices[j] = StrategicChoice{
				Label:         c.Label,
				Description:   c.Description,
				SuccessChance: c.SuccessChance,
				Rewards:       c.Rewards,
				Risks:         c.Risks,
				Requirements:  c.Requirements,
			}
		}

		// Radar reduces the detection threshold (making it easier to navigate stealthily)
		baseThreshold := 500 + rand.Intn(500)
		detectionThreshold := int(float64(baseThreshold) / (1.0 + float64(radarLevel-1)*0.2))

		nodes[i] = Node{
			ID:                     uuid.New(),
			ExpeditionID:           expeditionID,
			BlueprintID:            blueprint.ID,
			Name:                   blueprint.Name,
			Type:                   nodeType,
			Zone:                   zone,
			Hazard:                 hazard,
			EnvironmentDescription: blueprint.Description,
			DifficultyMultiplier:   1.0 + (float64(i) * 0.1),
			PositionIndex:          i,
			Choices:                choices,
			IsResolved:             false,
			Terrain:                terrain,
			DetectionThreshold:     detectionThreshold,
		}
	}
	return nodes
}

// CalculateEffectiveCP implements the blueprint formula:
// ECP = (Base_CP * Suitability_Mod) * Resonance_Sync * (1 - Fatigue_Penalty)
// CalculateEffectiveCP implements the blueprint formula:
// ECP = (Vehicle_CP + Exosuit_CP) * Suitability_Mod * Resonance_Sync * (1 - Fatigue_Penalty) * Synergy_Mod
func (s *Service) CalculateEffectiveCP(ctx context.Context, userID uuid.UUID, vehicleID uuid.UUID, terrain TerrainType) (int, error) {
	// 1. Get Pilot Stats
	pilot, err := s.gameRepo.GetActivePilotStats(userID)
	if err != nil {
		return 0, err
	}

	// Handle System NPC or Missing Pilot
	if pilot == nil {
		// Default stats for System NPC
		if vehicleID == uuid.Nil {
			return 50, nil
		}
		vehicleCP, _ := s.vehicleUseCase.GetVehicleCP(ctx, vehicleID)
		return vehicleCP, nil
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

		// Check for Critical Fatigue (from Emergency Retrieval)
		if pilot.Metadata != nil {
			if critical, ok := pilot.Metadata["critical_fatigue"].(bool); ok && critical {
				fatiguePenalty += 0.5 // Additional 50% penalty
				if fatiguePenalty > 0.9 {
					fatiguePenalty = 0.9 // Max penalty 90%
				}
			}
		}

		ecp := 50.0 * (1.0 - fatiguePenalty)
		return int(ecp), nil
	}

	// 3. Get Vehicle Base CP
	vehicleCP, err := s.vehicleUseCase.GetVehicleCP(ctx, vehicleID)
	if err != nil {
		return 0, err
	}

	// 4. Get Vehicle for Metadata and Suitability
	v, err := s.vehicleUseCase.GetVehicleByID(ctx, vehicleID)
	if err != nil {
		return 0, err
	}

	// 5. Get Exosuit CP and Metadata
	exosuitCP := 0
	var exosuitSeries string
	if pilot.EquippedExosuitID != nil {
		// Exosuit is an Item, so we use GetItemByID
		exosuitItem, err := s.vehicleUseCase.GetItemByID(ctx, *pilot.EquippedExosuitID)
		if err != nil {
			fmt.Printf("Error getting exosuit item: %v\n", err)
		} else if exosuitItem != nil {
			// Calculate Exosuit CP: (ATK*2) + (DEF*2) + (HP/10)
			// Note: ItemStats might need to be checked for nil, but struct usually has zero values
			atk := exosuitItem.Stats.Attack + exosuitItem.Stats.BonusAttack
			def := exosuitItem.Stats.Defense + exosuitItem.Stats.BonusDefense
			hp := exosuitItem.Stats.HP + exosuitItem.Stats.BonusHP
			
			exosuitCP = (atk * 2) + (def * 2) + (hp / 10)

			// Get Series ID
			if exosuitItem.SeriesID != nil {
				exosuitSeries = *exosuitItem.SeriesID
			} else if exosuitItem.Metadata != nil {
				// Fallback to metadata
				if meta, ok := exosuitItem.Metadata.(map[string]interface{}); ok {
					if s, ok := meta["Series"].(string); ok {
						exosuitSeries = s
					}
				}
			}
		}
	}

	// 6. Calculate Suitability Modifier
	suitabilityMod := 1.0
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
		// Check for incompatible types
		if (v.VehicleType == vehicle.TypeTank && terrain == TerrainIslands) ||
			(v.VehicleType == vehicle.TypeShip && terrain == TerrainDesert) {
			suitabilityMod = 0.5
		}
	}

	// 7. Calculate Resonance Sync
	// Formula: Min(1.0, Pilot_Resonance / Vehicle_Tier_Requirement)
	tierReq := v.Tier * 20
	resonanceSync := 1.0
	if tierReq > 0 && pilot.ResonanceLevel < tierReq {
		resonanceSync = float64(pilot.ResonanceLevel) / float64(tierReq)
	}
	if resonanceSync < 0.1 {
		resonanceSync = 0.1 // Minimum sync
	}

	// 8. Calculate Fatigue Penalty
	// Formula: Stress / 200 (Max 50% penalty at 100 Stress)
	fatiguePenalty := float64(pilot.Stress) / 200.0
	if fatiguePenalty > 0.5 {
		fatiguePenalty = 0.5
	}

	// Check for Critical Fatigue (from Emergency Retrieval)
	if pilot.Metadata != nil {
		if critical, ok := pilot.Metadata["critical_fatigue"].(bool); ok && critical {
			fatiguePenalty += 0.5 // Additional 50% penalty
			if fatiguePenalty > 0.9 {
				fatiguePenalty = 0.9 // Max penalty 90%
			}
		}
	}

	// 9. Calculate Set Synergy
	synergyMod := 1.0
	vehicleSeries := ""
	if v.Metadata != nil {
		if meta, ok := v.Metadata.(map[string]interface{}); ok {
			if s, ok := meta["Series"].(string); ok {
				vehicleSeries = s
			}
		}
	}

	if vehicleSeries != "" && exosuitSeries != "" && vehicleSeries == exosuitSeries {
		synergyMod = 1.15 // +15% Synergy Bonus
	}

	// 10. Check for Active Skill Buffs (Overclock)
	skillMod := 1.0
	if pilot.Metadata != nil {
		if active, ok := pilot.Metadata["active_skill_overclock"].(bool); ok && active {
			skillMod = 1.3 // +30% ECP
			// Note: The flag should be cleared after use. This logic should be in ResolveNodeChoice or similar.
			// For calculation purposes, we just apply the mod.
		}
	}

	// 11. Final ECP Calculation
	totalBaseCP := float64(vehicleCP + exosuitCP)
	ecp := totalBaseCP * suitabilityMod * resonanceSync * (1.0 - fatiguePenalty) * synergyMod * skillMod

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

	vehicleID := uuid.Nil
	if expedition.VehicleID != nil {
		vehicleID = *expedition.VehicleID
	}

	// 3.1 Fetch Blueprint & Pilot Stats
	blueprint, hasBlueprint := s.blueprints.Nodes[node.BlueprintID]
	stats, err := s.gameRepo.GetActivePilotStats(expedition.UserID)
	if err != nil || stats == nil {
		return nil, fmt.Errorf("failed to fetch pilot stats")
	}

	// 3.2 Resource Check (Pre-resolution)
	fuelCost := 5.0
	o2Cost := 3.0
	if hasBlueprint {
		fuelCost = blueprint.ResourceCosts.Fuel
		o2Cost = blueprint.ResourceCosts.O2
	}

	if stats.CurrentFuel < fuelCost || stats.CurrentO2 < o2Cost {
		// Trigger Emergency Retrieval
		stats.Stress = 100
		if stats.Metadata == nil {
			stats.Metadata = make(map[string]interface{})
		}
		stats.Metadata["emergency_retrieval"] = true
		_ = s.gameRepo.UpdatePilotStats(stats)
		
		// Mark expedition as failed/ended
		expedition.Status = "FAILED"
		// In a real system, we'd update the expedition status in DB
		
		return nil, fmt.Errorf("EMERGENCY RETRIEVAL: Insufficient resources (Fuel: %.1f/%.1f, O2: %.1f/%.1f)", 
			stats.CurrentFuel, fuelCost, stats.CurrentO2, o2Cost)
	}

	// 3.3 Suitability Check (Tags)
	requirementPenalty := 0.0
	if hasBlueprint && vehicleID != uuid.Nil {
		v, _ := s.vehicleUseCase.GetVehicleByID(ctx, vehicleID)
		if v != nil {
			// Check Required Tags
			for _, reqTag := range blueprint.RequiredTags {
				found := false
				for _, vTag := range v.SuitabilityTags {
					if vTag == reqTag {
						found = true
						break
					}
				}
				if !found {
					requirementPenalty -= 0.3
					fmt.Printf("WARNING: Missing required tag %s. Applying penalty.\n", reqTag)
				}
			}
			// Check Forbidden Tags
			for _, forbTag := range blueprint.ForbiddenTags {
				for _, vTag := range v.SuitabilityTags {
					if vTag == forbTag {
						requirementPenalty -= 0.5
						fmt.Printf("WARNING: Forbidden tag %s detected. Applying heavy penalty.\n", forbTag)
						break
					}
				}
			}
		}
	}

	// 3.5 Zone Requirements Check (Legacy/Fallback)
	if node.Zone == ZoneEVA && vehicleID != uuid.Nil {
		// Penalty for bringing a heavy vehicle into tight EVA spaces
		requirementPenalty -= 0.5
		fmt.Printf("WARNING: Heavy vehicle in EVA zone. Applying 50%% penalty.\n")
	}

	if node.Zone == ZoneCorridor {
		if vehicleID == uuid.Nil {
			// Penalty for being on foot in a high-speed corridor
			requirementPenalty -= 0.6
			fmt.Printf("WARNING: On foot in high-speed corridor. Applying 60%% penalty.\n")
		} else {
			v, _ := s.vehicleUseCase.GetVehicleByID(ctx, vehicleID)
			if v != nil && v.VehicleType != vehicle.TypeSpeeder {
				// Penalty for non-speeder vehicle
				requirementPenalty -= 0.4
				fmt.Printf("WARNING: Non-speeder in high-speed corridor. Applying 40%% penalty.\n")
			}
		}
	}

	// 4. Calculate Effective CP (ECP)
	ecp, err := s.CalculateEffectiveCP(ctx, expedition.UserID, vehicleID, node.Terrain)
	if err != nil {
		// Fallback if calculation fails
		ecp = 100
	}

	// Success Chance Adjustment based on ECP vs Node Difficulty
	// Base difficulty is 200 * DifficultyMultiplier
	baseDifficulty := 200.0 * node.DifficultyMultiplier

	// Detection Logic (Alarm Mode)
	isAlarmMode := false
	if vehicleID != uuid.Nil {
		// Calculate Signature (Base on CP for now)
		signature := float64(ecp)
		
		// Bastion Radar reduces effective signature
		radarLevel := 1.0
		modules, err := s.gameRepo.GetBastionModules(expedition.UserID)
		if err == nil {
			for _, m := range modules {
				if m.IsActive && m.ModuleType == "RADAR" {
					radarLevel = float64(m.Level)
				}
			}
		}
		
		effectiveSignature := signature / (1.0 + (radarLevel-1)*0.2)
		
		if int(effectiveSignature) > node.DetectionThreshold {
			isAlarmMode = true
			// Alarm Mode increases difficulty significantly
			baseDifficulty *= 2.0
			fmt.Printf("ALARM MODE TRIGGERED: Signature %v > Threshold %v\n", int(effectiveSignature), node.DetectionThreshold)
		}
	}

	if isAlarmMode {
		// Could add more alarm-specific logic here
	}

	// Success chance formula: 0.5 + (ECP - Difficulty) / 1000
	ecpBonus := (float64(ecp) - baseDifficulty) / 1000.0
	if ecpBonus < -0.3 {
		ecpBonus = -0.3 // Max penalty
	}
	if ecpBonus > 0.4 {
		ecpBonus = 0.4 // Max bonus
	}

	finalSuccessChance := selectedChoice.SuccessChance + ecpBonus + requirementPenalty
	if finalSuccessChance > 0.98 {
		finalSuccessChance = 0.98
	}
	if finalSuccessChance < 0.05 {
		finalSuccessChance = 0.05
	}

	success := rand.Float64() < finalSuccessChance

	// 5. Apply Consequences
	stats, err = s.gameRepo.GetActivePilotStats(expedition.UserID)
	if err == nil && stats != nil {
		// Get Bastion Module Levels (Migrated to Table)
		labLevel := 1.0
		warpLevel := 1.0
		
		modules, err := s.gameRepo.GetBastionModules(expedition.UserID)
		if err == nil {
			for _, m := range modules {
				if m.IsActive {
					switch m.ModuleType {
					case "LAB":
						labLevel = float64(m.Level)
					case "WARP_DRIVE":
						warpLevel = float64(m.Level)
					}
				}
			}
		}

		// Every node resolution increases Stress
		stats.Stress += 5 + rand.Intn(5) // 5-10 stress per node
		
		// Apply Hazard Effects
		switch node.Hazard {
		case HazardVoidEcho:
			stats.Stress += 10 // Extra stress from Void Echoes
		case HazardSolarFlare:
			stats.Stress += 5
			// Could add Heat mechanic here
		case HazardCorrosiveRain:
			if expedition.VehicleID != nil {
				// Apply durability damage to vehicle
				_, _ = s.vehicleUseCase.ApplyDamage(ctx, *expedition.VehicleID, 5)
			}
		case HazardEMPStorm:
			// EMP reduces success chance or affects energy weapons in combat
		}

		if stats.Stress > 100 {
			stats.Stress = 100
		}

		actualFuelCost := fuelCost / (1.0 + (warpLevel-1)*0.1)
		stats.CurrentFuel -= actualFuelCost
		if stats.CurrentFuel < 0 {
			stats.CurrentFuel = 0
		}

		// O2 Consumption
		stats.CurrentO2 -= o2Cost
		if stats.CurrentO2 < 0 {
			stats.CurrentO2 = 0
		}

		// Clear Active Skill Flags (Overclock)
		if stats.Metadata != nil {
			if _, ok := stats.Metadata["active_skill_overclock"]; ok {
				delete(stats.Metadata, "active_skill_overclock")
			}
		}

		// Neural Energy Gain (Phase 4: NE System)
		stats.CurrentNE += 10.0
		if stats.CurrentNE > stats.MaxNE {
			stats.CurrentNE = stats.MaxNE
		}

		// Emergency Retrieval Protocol (Phase 2: Tactical Engine)
		isEmergency := false
		if stats.CurrentFuel <= 0 || stats.CurrentO2 <= 0 {
			isEmergency = true
			stats.Stress = 100
			if stats.Metadata == nil {
				stats.Metadata = make(map[string]interface{})
			}
			stats.Metadata["emergency_retrieval"] = true
			fmt.Printf("CRITICAL: Resources exhausted. Emergency Retrieval initiated.\n")
		}

		if success {
			rewardMod := 1.0 + (labLevel-1)*0.1 // 10% bonus per level
			xpMod := 1.0 + (labLevel-1)*0.15    // 15% bonus per level

			if isEmergency {
				rewardMod *= 0.5 // 50% Penalty for Emergency Retrieval
				xpMod *= 0.5
			}

			for _, reward := range selectedChoice.Rewards {
				switch reward {
				case "Scrap Metal":
					stats.ScrapMetal += int(50.0 * rewardMod)
				case "Research Data":
					stats.ResearchData += int(20.0 * rewardMod)
				case "Rare Ore":
					stats.ScrapMetal += int(150.0 * rewardMod)
				}
			}
			// Always give some XP on success
			stats.XP += int(15.0 * xpMod)
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
	// Find matching blueprint
	var blueprint game.NodeBlueprint
	found := false
	for _, b := range s.blueprints.Nodes {
		if b.Type == string(t) {
			blueprint = b
			found = true
			break
		}
	}

	if !found {
		return []StrategicChoice{
			{Label: "Proceed", Description: "Move forward carefully.", SuccessChance: 1.0},
		}
	}

	choices := make([]StrategicChoice, len(blueprint.Choices))
	for i, c := range blueprint.Choices {
		choices[i] = StrategicChoice{
			Label:         c.Label,
			Description:   c.Description,
			SuccessChance: c.SuccessChance,
			Rewards:       c.Rewards,
			Risks:         c.Risks,
			Requirements:  c.Requirements,
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
	// Skip for system expeditions or if pilot stats are missing
	if expedition.UserID != uuid.Nil && pilot != nil && currentIndex > 0 {
		success, err := s.gameRepo.ConsumeResources(pilot.CharacterID, 15.0, 5.0)
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, fmt.Errorf("insufficient resources to advance (O2: %.1f, Fuel: %.1f)", pilot.CurrentO2, pilot.CurrentFuel)
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
	if encounterType == NodeCombat || encounterType == "boss" {
		if targetNode.EnemyBlueprint != "" {
			// Find enemy by name in blueprints
			for id, b := range s.blueprints.Enemies {
				if b.Name == targetNode.EnemyBlueprint || id == targetNode.EnemyBlueprint {
					uid := uuid.MustParse(id)
					enemyID = &uid
					break
				}
			}
		}

		// Fallback to random enemy if not found or not specified
		if enemyID == nil {
			var enemyIDs []string
			for id := range s.blueprints.Enemies {
				enemyIDs = append(enemyIDs, id)
			}
			
			if len(enemyIDs) > 0 {
				selected := enemyIDs[rand.Intn(len(enemyIDs))]
				uid := uuid.MustParse(selected)
				enemyID = &uid
			}
		}
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

// ActivateSkill handles the usage of Neural Energy (NE) for active skills
func (s *Service) ActivateSkill(ctx context.Context, userID uuid.UUID, skillName string) error {
	// 1. Get Pilot Stats
	stats, err := s.gameRepo.GetActivePilotStats(userID)
	if err != nil {
		return err
	}

	// 2. Check NE Cost and Apply Effect
	cost := 0.0
	switch skillName {
	case "OVERCLOCK":
		cost = 50.0
		if stats.CurrentNE < cost {
			return fmt.Errorf("insufficient neural energy")
		}
		// Set Overclock flag in metadata
		if stats.Metadata == nil {
			stats.Metadata = make(map[string]interface{})
		}
		stats.Metadata["active_skill_overclock"] = true

	case "EMERGENCY_REPAIR":
		cost = 40.0
		if stats.CurrentNE < cost {
			return fmt.Errorf("insufficient neural energy")
		}
		// Restore Vehicle HP (Durability)
		// We need to find the active vehicle first. This is tricky without expedition context.
		// Assuming we repair the currently equipped vehicle or we need expedition ID.
		// For now, let's assume this is called within an expedition context or we find the active vehicle.
		// Simplified: Repair the vehicle in the active session.
		session, err := s.repo.GetSessionByUserID(userID)
		if err != nil || session == nil {
			return fmt.Errorf("no active exploration session found")
		}
		if session.VehicleID != uuid.Nil {
			// Repair 30% of Max Durability
			// We need to fetch the item to know max durability
			item, err := s.vehicleUseCase.GetItemByID(ctx, session.VehicleID)
			if err != nil {
				return fmt.Errorf("vehicle item not found")
			}
			repairAmount := int(float64(item.MaxDurability) * 0.3)
			_, err = s.vehicleUseCase.RepairItem(ctx, session.VehicleID, repairAmount)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("no vehicle to repair")
		}

	default:
		return fmt.Errorf("unknown skill: %s", skillName)
	}

	// 3. Consume NE
	stats.CurrentNE -= cost
	return s.gameRepo.UpdatePilotStats(stats)
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
