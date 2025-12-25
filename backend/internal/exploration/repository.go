package exploration

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type explorationRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &explorationRepository{db: db}
}

func (r *explorationRepository) CreateNodes(nodes []Node) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO nodes (id, expedition_id, name, type, environment_description, difficulty_multiplier, position_index, choices, is_resolved) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	for _, n := range nodes {
		choicesJSON, err := json.Marshal(n.Choices)
		if err != nil {
			return err
		}
		_, err = tx.Exec(query, n.ID, n.ExpeditionID, n.Name, n.Type, n.EnvironmentDescription, n.DifficultyMultiplier, n.PositionIndex, choicesJSON, n.IsResolved)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *explorationRepository) GetNodesByExpeditionID(expeditionID uuid.UUID) ([]Node, error) {
	query := `SELECT id, expedition_id, name, type, environment_description, difficulty_multiplier, position_index, choices, is_resolved 
	          FROM nodes WHERE expedition_id = $1 ORDER BY position_index ASC`
	rows, err := r.db.Query(query, expeditionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []Node
	for rows.Next() {
		var n Node
		var choicesJSON []byte
		if err := rows.Scan(&n.ID, &n.ExpeditionID, &n.Name, &n.Type, &n.EnvironmentDescription, &n.DifficultyMultiplier, &n.PositionIndex, &choicesJSON, &n.IsResolved); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(choicesJSON, &n.Choices); err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}
	return nodes, nil
}

func (r *explorationRepository) GetNodeByID(id uuid.UUID) (*Node, error) {
	query := `SELECT id, expedition_id, name, type, environment_description, difficulty_multiplier, position_index, choices, is_resolved 
	          FROM nodes WHERE id = $1`
	var n Node
	var choicesJSON []byte
	err := r.db.QueryRow(query, id).Scan(&n.ID, &n.ExpeditionID, &n.Name, &n.Type, &n.EnvironmentDescription, &n.DifficultyMultiplier, &n.PositionIndex, &choicesJSON, &n.IsResolved)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(choicesJSON, &n.Choices); err != nil {
		return nil, err
	}
	return &n, nil
}

func (r *explorationRepository) UpdateNode(n *Node) error {
	choicesJSON, err := json.Marshal(n.Choices)
	if err != nil {
		return err
	}
	query := `UPDATE nodes SET is_resolved = $1, choices = $2 WHERE id = $3`
	_, err = r.db.Exec(query, n.IsResolved, choicesJSON, n.ID)
	return err
}

func (r *explorationRepository) CreateExpedition(e *Expedition) error {
	query := `INSERT INTO expeditions (id, user_id, sub_sector_id, planet_location_id, vehicle_id, title, description, goal) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(query, e.ID, e.UserID, e.SubSectorID, e.PlanetLocationID, e.VehicleID, e.Title, e.Description, e.Goal)
	return err
}

func (r *explorationRepository) GetExpeditionByID(id uuid.UUID) (*Expedition, error) {
	query := `SELECT id, user_id, sub_sector_id, planet_location_id, vehicle_id, title, description, goal FROM expeditions WHERE id = $1`
	var e Expedition
	err := r.db.QueryRow(query, id).Scan(&e.ID, &e.UserID, &e.SubSectorID, &e.PlanetLocationID, &e.VehicleID, &e.Title, &e.Description, &e.Goal)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *explorationRepository) SaveEncounter(e *Encounter, expeditionID uuid.UUID) error {
	query := `INSERT INTO encounters (id, expedition_id, type, title, description, visual_prompt, enemy_id) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, e.ID, expeditionID, e.Type, e.Title, e.Description, e.VisualPrompt, e.EnemyID)
	return err
}

func (r *explorationRepository) GetEncountersByExpeditionID(expeditionID uuid.UUID) ([]Encounter, error) {
	query := `SELECT id, type, title, description, visual_prompt, enemy_id, created_at FROM encounters WHERE expedition_id = $1 ORDER BY created_at ASC`
	rows, err := r.db.Query(query, expeditionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var encounters []Encounter
	for rows.Next() {
		var e Encounter
		if err := rows.Scan(&e.ID, &e.Type, &e.Title, &e.Description, &e.VisualPrompt, &e.EnemyID, &e.CreatedAt); err != nil {
			return nil, err
		}
		encounters = append(encounters, e)
	}
	return encounters, nil
}

func (r *explorationRepository) GetSessionByUserID(userID uuid.UUID) (*Session, error) {
	query := `SELECT id, user_id, vehicle_id, current_node_id, status FROM exploration_sessions WHERE user_id = $1`
	var s Session
	err := r.db.QueryRow(query, userID).Scan(&s.ID, &s.UserID, &s.VehicleID, &s.CurrentNodeID, &s.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *explorationRepository) GetAllSectors() ([]Sector, error) {
	query := `SELECT id, name, description, difficulty, coordinates_x, coordinates_y, color FROM sectors WHERE is_active = TRUE`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sectors []Sector
	for rows.Next() {
		var s Sector
		if err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Difficulty, &s.CoordinatesX, &s.CoordinatesY, &s.Color); err != nil {
			return nil, err
		}
		sectors = append(sectors, s)
	}
	return sectors, nil
}

func (r *explorationRepository) GetSubSectorsBySectorID(sectorID uuid.UUID) ([]SubSector, error) {
	query := `SELECT id, sector_id, type, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y FROM sub_sectors WHERE sector_id = $1`
	rows, err := r.db.Query(query, sectorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subSectors []SubSector
	for rows.Next() {
		var ss SubSector
		if err := rows.Scan(&ss.ID, &ss.SectorID, &ss.Type, &ss.Name, &ss.Description, pq.Array(&ss.Rewards), pq.Array(&ss.Requirements), pq.Array(&ss.AllowedModes), &ss.RequiresAtmosphere, &ss.SuitabilityPilot, &ss.SuitabilityVehicle, &ss.CoordinatesX, &ss.CoordinatesY); err != nil {
			return nil, err
		}
		subSectors = append(subSectors, ss)
	}
	return subSectors, nil
}

func (r *explorationRepository) GetPlanetLocationsBySubSectorID(subSectorID uuid.UUID) ([]PlanetLocation, error) {
	query := `SELECT id, sub_sector_id, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_vehicle, coordinates_x, coordinates_y FROM planet_locations WHERE sub_sector_id = $1`
	rows, err := r.db.Query(query, subSectorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []PlanetLocation
	for rows.Next() {
		var pl PlanetLocation
		if err := rows.Scan(&pl.ID, &pl.SubSectorID, &pl.Name, &pl.Description, pq.Array(&pl.Rewards), pq.Array(&pl.Requirements), pq.Array(&pl.AllowedModes), &pl.RequiresAtmosphere, &pl.SuitabilityPilot, &pl.SuitabilityVehicle, &pl.CoordinatesX, &pl.CoordinatesY); err != nil {
			return nil, err
		}
		locations = append(locations, pl)
	}
	return locations, nil
}

func (r *explorationRepository) GetNodesByStarID(starID uuid.UUID) ([]Node, error) {
	query := `SELECT id, name, type, environment_description, difficulty_multiplier, position_index, choices, is_resolved FROM nodes WHERE star_id = $1 ORDER BY position_index ASC`
	rows, err := r.db.Query(query, starID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []Node
	for rows.Next() {
		var n Node
		var choicesJSON []byte
		if err := rows.Scan(&n.ID, &n.Name, &n.Type, &n.EnvironmentDescription, &n.DifficultyMultiplier, &n.PositionIndex, &choicesJSON, &n.IsResolved); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(choicesJSON, &n.Choices); err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}
	return nodes, nil
}

func (r *explorationRepository) CreateSession(s *Session) error {
	query := `INSERT INTO exploration_sessions (id, user_id, vehicle_id, current_node_id, status) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, s.ID, s.UserID, s.VehicleID, s.CurrentNodeID, s.Status)
	return err
}

func (r *explorationRepository) UpdateSession(s *Session) error {
	query := `UPDATE exploration_sessions SET current_node_id = $1, status = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3`
	_, err := r.db.Exec(query, s.CurrentNodeID, s.Status, s.ID)
	return err
}
