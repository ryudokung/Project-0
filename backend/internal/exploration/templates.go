package exploration

type NodeTemplate struct {
	ID          string
	Name        string
	Description string
	Type        NodeType
	Choices     []StrategicChoice
	RewardPool  []string
	MinCP       int // Minimum Combat Power recommended
}

var NodeTemplates = map[string]NodeTemplate{
	"MINING_OUTPOST": {
		ID:          "MINING_OUTPOST",
		Name:        "Abandoned Mining Outpost",
		Description: "A cluster of rusted drills and storage containers.",
		Type:        NodeResource,
		Choices: []StrategicChoice{
			{
				Label: "Deep Drill",
				Description: "Extract rare minerals from the core.",
				SuccessChance: 0.4,
				Rewards: []string{"Rare Ore", "Scrap Metal"},
				Risks: []string{"Structural Stress"},
				Requirements: []string{"CP > 150"},
			},
			{
				Label: "Surface Scavenge",
				Description: "Quickly gather loose materials.",
				SuccessChance: 0.9,
				Rewards: []string{"Scrap Metal"},
				Requirements: []string{},
			},
		},
		MinCP: 100,
	},
	"SYNDICATE_AMBUSH": {
		ID:          "SYNDICATE_AMBUSH",
		Name:        "Syndicate Ambush Point",
		Description: "Hostile signatures detected behind the asteroid belt.",
		Type:        NodeCombat,
		Choices: []StrategicChoice{
			{
				Label: "Full Assault",
				Description: "Direct confrontation with maximum firepower.",
				SuccessChance: 0.6,
				Risks: []string{"High Damage"},
				Requirements: []string{"CP > 300"},
			},
			{
				Label: "Tactical Flank",
				Description: "Use agility to find a weak spot.",
				SuccessChance: 0.8,
				Risks: []string{"Medium Damage"},
				Requirements: []string{"PILOT_AGILITY > 40", "CP > 200"},
			},
		},
		MinCP: 250,
	},
	"VOID_ANOMALY": {
		ID:          "VOID_ANOMALY",
		Name:        "Shimmering Void Anomaly",
		Description: "Space itself seems to be folding here.",
		Type:        NodeAnomaly,
		Choices: []StrategicChoice{
			{
				Label: "Scientific Study",
				Description: "Analyze the anomaly for data.",
				SuccessChance: 0.7,
				Rewards: []string{"Research Data", "Void Shard"},
				Requirements: []string{"PILOT_INTEL > 50"},
			},
			{
				Label: "Brute Force",
				Description: "Push through the anomaly.",
				SuccessChance: 0.5,
				Risks: []string{"System Glitch"},
				Requirements: []string{"CP > 400"},
			},
		},
		MinCP: 300,
	},
}

// GetTemplateForType returns a random template for a given node type
func GetTemplateForType(t NodeType) NodeTemplate {
	// For now, just return a default or find one that matches
	for _, tmpl := range NodeTemplates {
		if tmpl.Type == t {
			return tmpl
		}
	}
	// Fallback
	return NodeTemplate{
		Name: "Unknown Sector",
		Type: t,
		Choices: []StrategicChoice{
			{Label: "Proceed", Description: "Move forward carefully.", SuccessChance: 1.0},
		},
	}
}
