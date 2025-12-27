package exploration

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/game"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateExpedition(expedition *Expedition) error {
	args := m.Called(expedition)
	return args.Error(0)
}

func (m *MockRepo) GetExpeditionByID(id uuid.UUID) (*Expedition, error) {
	args := m.Called(id)
	return args.Get(0).(*Expedition), args.Error(1)
}

func (m *MockRepo) CreateNodes(nodes []Node) error {
	args := m.Called(nodes)
	return args.Error(0)
}

func (m *MockRepo) GetNodesByExpeditionID(expeditionID uuid.UUID) ([]Node, error) {
	args := m.Called(expeditionID)
	return args.Get(0).([]Node), args.Error(1)
}

func (m *MockRepo) GetNodeByID(id uuid.UUID) (*Node, error) {
	args := m.Called(id)
	return args.Get(0).(*Node), args.Error(1)
}

func (m *MockRepo) UpdateNode(node *Node) error {
	args := m.Called(node)
	return args.Error(0)
}

func (m *MockRepo) SaveEncounter(encounter *Encounter, expeditionID uuid.UUID) error {
	args := m.Called(encounter, expeditionID)
	return args.Error(0)
}

func (m *MockRepo) GetEncountersByExpeditionID(expeditionID uuid.UUID) ([]Encounter, error) {
	args := m.Called(expeditionID)
	return args.Get(0).([]Encounter), args.Error(1)
}

func (m *MockRepo) GetSessionByUserID(userID uuid.UUID) (*Session, error) {
	args := m.Called(userID)
	return args.Get(0).(*Session), args.Error(1)
}

func (m *MockRepo) GetAllSectors() ([]Sector, error) {
	args := m.Called()
	return args.Get(0).([]Sector), args.Error(1)
}

func (m *MockRepo) GetSubSectorsBySectorID(sectorID uuid.UUID) ([]SubSector, error) {
	args := m.Called(sectorID)
	return args.Get(0).([]SubSector), args.Error(1)
}

func (m *MockRepo) GetPlanetLocationsBySubSectorID(subSectorID uuid.UUID) ([]PlanetLocation, error) {
	args := m.Called(subSectorID)
	return args.Get(0).([]PlanetLocation), args.Error(1)
}

func TestGenerateTimeline(t *testing.T) {
	blueprints := &game.BlueprintRegistry{
		Nodes: map[string]game.NodeBlueprint{
			"STANDARD": {ID: "STANDARD", Name: "Quiet Sector", Type: "STANDARD"},
			"OUTPOST":  {ID: "OUTPOST", Name: "Abandoned Outpost", Type: "OUTPOST"},
		},
	}
	s := &Service{blueprints: blueprints}
	expeditionID := uuid.New()
	length := 5
	radarLevel := 1

	nodes := s.GenerateTimeline(expeditionID, length, radarLevel)

	assert.Equal(t, length, len(nodes))
	assert.Equal(t, ZoneOrbital, nodes[0].Zone)
	// Note: In the current implementation, the last node is forced to NodeOutpost
	assert.Equal(t, NodeOutpost, nodes[length-1].Type)

	for _, node := range nodes {
		assert.Equal(t, expeditionID, node.ExpeditionID)
		assert.NotEmpty(t, node.Terrain)
		assert.NotEmpty(t, node.Zone)
	}
}
