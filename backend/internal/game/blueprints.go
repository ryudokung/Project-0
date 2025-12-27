package game

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type NodeBlueprint struct {
	ID            string            `yaml:"id"`
	Name          string            `yaml:"name"`
	Description   string            `yaml:"description"`
	Type          string            `yaml:"type"`
	Zone          string            `yaml:"zone"`
	MinCP         int               `yaml:"min_cp"`
	RequiredTags  []string          `yaml:"required_tags"`
	ForbiddenTags []string          `yaml:"forbidden_tags"`
	ResourceCosts struct {
		Fuel float64 `yaml:"fuel"`
		O2   float64 `yaml:"o2"`
	} `yaml:"resource_costs"`
	Choices []ChoiceBlueprint `yaml:"choices"`
}

type ChoiceBlueprint struct {
	Label         string   `yaml:"label"`
	Description   string   `yaml:"description"`
	SuccessChance float64  `yaml:"success_chance"`
	Rewards       []string `yaml:"rewards"`
	Risks         []string `yaml:"risks"`
	Requirements  []string `yaml:"requirements"`
}

type EnemyBlueprint struct {
	ID     string `yaml:"id"`
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	Class  string `yaml:"class"`
	Rarity string `yaml:"rarity"`
	CR     int    `yaml:"cr"`
	Stats  struct {
		HP      int `yaml:"hp"`
		Attack  int `yaml:"attack"`
		Defense int `yaml:"defense"`
		Speed   int `yaml:"speed"`
	} `yaml:"stats"`
}

type ExpeditionBlueprint struct {
	ID          string `yaml:"id"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Difficulty  int    `yaml:"difficulty"`
	MinLevel    int    `yaml:"min_level"`
	Nodes       []struct {
		ID                     string        `yaml:"id"`
		Type                   string        `yaml:"type"`
		Title                  string        `yaml:"title"`
		Description            string        `yaml:"description"`
		ResourceType           string        `yaml:"resource_type,omitempty"`
		Amount                 int           `yaml:"amount,omitempty"`
		EnemyBlueprint         string        `yaml:"enemy_blueprint,omitempty"`
		EnemyCount             int           `yaml:"enemy_count,omitempty"`
		IsScripted             bool          `yaml:"is_scripted,omitempty"`
		ScriptEvents           []ScriptEvent `yaml:"script_events,omitempty"`
		NextNodes              []string      `yaml:"next_nodes,omitempty"`
		IsEnd                  bool          `yaml:"is_end,omitempty"`
		EnvironmentDescription string        `yaml:"environment_description,omitempty"`
	} `yaml:"nodes"`
}

type BlueprintRegistry struct {
	Nodes       map[string]NodeBlueprint
	Enemies     map[string]EnemyBlueprint
	Expeditions map[string]ExpeditionBlueprint
}

func NewBlueprintRegistry() *BlueprintRegistry {
	return &BlueprintRegistry{
		Nodes:       make(map[string]NodeBlueprint),
		Enemies:     make(map[string]EnemyBlueprint),
		Expeditions: make(map[string]ExpeditionBlueprint),
	}
}

func (r *BlueprintRegistry) LoadExpeditions(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var config struct {
		Expeditions []ExpeditionBlueprint `yaml:"expeditions"`
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	for _, exp := range config.Expeditions {
		r.Expeditions[exp.ID] = exp
	}

	fmt.Printf("Loaded %d expedition blueprints from %s\n", len(config.Expeditions), path)
	return nil
}

func (r *BlueprintRegistry) LoadNodes(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var config struct {
		Nodes []NodeBlueprint `yaml:"nodes"`
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	for _, node := range config.Nodes {
		r.Nodes[node.ID] = node
	}

	fmt.Printf("Loaded %d node blueprints from %s\n", len(config.Nodes), path)
	return nil
}

func (r *BlueprintRegistry) LoadEnemies(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var config struct {
		Enemies []EnemyBlueprint `yaml:"enemies"`
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	for _, enemy := range config.Enemies {
		r.Enemies[enemy.ID] = enemy
	}

	fmt.Printf("Loaded %d enemy blueprints from %s\n", len(config.Enemies), path)
	return nil
}
