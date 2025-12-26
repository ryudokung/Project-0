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

func (s *Service) GenerateTimeline(expeditionID uuid.UUID, length int, radarLevel int) []Node {
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

		// Radar reduces the detection threshold (making it easier to navigate stealthily)
		baseThreshold := 500 + rand.Intn(500)
		detectionThreshold := int(float64(baseThreshold) / (1.0 + float64(radarLevel-1)*0.2))

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
		if stats.Stress > 100 {
			stats.Stress = 100
		}

		// Fuel Consumption (Warp Drive reduces this)
		fuelCost := 5.0 / (1.0 + (warpLevel-1)*0.1) // 10% reduction per level
		stats.CurrentFuel -= fuelCost
		if stats.CurrentFuel < 0 {
			stats.CurrentFuel = 0
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
			stats.Stress += 50
			if stats.Stress > 100 {
				stats.Stress = 100
			}
			if stats.Metadata == nil {
				stats.Metadata = make(map[string]interface{})
			}
			stats.Metadata["critical_fatigue"] = true
			// In a full implementation, this would trigger an immediate end to the expedition.
		}

		if success {
			// Apply Rewards (Lab increases these)
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
