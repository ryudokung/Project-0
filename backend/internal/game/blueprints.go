package game

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type NodeBlueprint struct {
	ID          string            `yaml:"id"`
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Type        string            `yaml:"type"`
	MinCP       int               `yaml:"min_cp"`
	Choices     []ChoiceBlueprint `yaml:"choices"`
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

type BlueprintRegistry struct {
	Nodes   map[string]NodeBlueprint
	Enemies map[string]EnemyBlueprint
}

func NewBlueprintRegistry() *BlueprintRegistry {
	return &BlueprintRegistry{
		Nodes:   make(map[string]NodeBlueprint),
		Enemies: make(map[string]EnemyBlueprint),
	}
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
