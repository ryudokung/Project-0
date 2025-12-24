package exploration

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type explorationRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &explorationRepository{db: db}
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
	query := `SELECT id, sector_id, type, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_mech, coordinates_x, coordinates_y FROM sub_sectors WHERE sector_id = $1`
	rows, err := r.db.Query(query, sectorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subSectors []SubSector
	for rows.Next() {
		var ss SubSector
		if err := rows.Scan(&ss.ID, &ss.SectorID, &ss.Type, &ss.Name, &ss.Description, pq.Array(&ss.Rewards), pq.Array(&ss.Requirements), pq.Array(&ss.AllowedModes), &ss.RequiresAtmosphere, &ss.SuitabilityPilot, &ss.SuitabilityMech, &ss.CoordinatesX, &ss.CoordinatesY); err != nil {
			return nil, err
		}
		subSectors = append(subSectors, ss)
	}
	return subSectors, nil
}

func (r *explorationRepository) GetPlanetLocationsBySubSectorID(subSectorID uuid.UUID) ([]PlanetLocation, error) {
	query := `SELECT id, sub_sector_id, name, description, rewards, requirements, allowed_modes, requires_atmosphere, suitability_pilot, suitability_mech, coordinates_x, coordinates_y FROM planet_locations WHERE sub_sector_id = $1`
	rows, err := r.db.Query(query, subSectorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []PlanetLocation
	for rows.Next() {
		var pl PlanetLocation
		if err := rows.Scan(&pl.ID, &pl.SubSectorID, &pl.Name, &pl.Description, pq.Array(&pl.Rewards), pq.Array(&pl.Requirements), pq.Array(&pl.AllowedModes), &pl.RequiresAtmosphere, &pl.SuitabilityPilot, &pl.SuitabilityMech, &pl.CoordinatesX, &pl.CoordinatesY); err != nil {
			return nil, err
		}
		locations = append(locations, pl)
	}
	return locations, nil
}

func (r *explorationRepository) GetNodesByStarID(starID uuid.UUID) ([]Node, error) {
	query := `SELECT id, star_id, name, type, environment_description, difficulty_multiplier, position_index FROM nodes WHERE star_id = $1 ORDER BY position_index ASC`
	rows, err := r.db.Query(query, starID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []Node
	for rows.Next() {
		var n Node
		if err := rows.Scan(&n.ID, &n.StarID, &n.Name, &n.Type, &n.EnvironmentDescription, &n.DifficultyMultiplier, &n.PositionIndex); err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}
	return nodes, nil
}

func (r *explorationRepository) CreateSession(s *Session) error {
	query := `INSERT INTO exploration_sessions (id, user_id, mech_id, current_node_id, status) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, s.ID, s.UserID, s.MechID, s.CurrentNodeID, s.Status)
	return err
}

func (r *explorationRepository) UpdateSession(s *Session) error {
	query := `UPDATE exploration_sessions SET current_node_id = $1, status = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3`
	_, err := r.db.Exec(query, s.CurrentNodeID, s.Status, s.ID)
	return err
}

func (r *explorationRepository) GetSessionByUserID(userID uuid.UUID) (*Session, error) {
	query := `SELECT id, user_id, mech_id, current_node_id, status FROM exploration_sessions WHERE user_id = $1 AND status = 'ACTIVE'`
	row := r.db.QueryRow(query, userID)

	var s Session
	err := row.Scan(&s.ID, &s.UserID, &s.MechID, &s.CurrentNodeID, &s.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &s, err
}

func (r *explorationRepository) CreateThread(t *Thread) error {
	query := `INSERT INTO threads (id, user_id, sub_sector_id, planet_location_id, vehicle_id, title, description, goal) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(query, t.ID, t.UserID, t.SubSectorID, t.PlanetLocationID, t.VehicleID, t.Title, t.Description, t.Goal)
	return err
}

func (r *explorationRepository) GetThreadByID(id uuid.UUID) (*Thread, error) {
	query := `SELECT id, user_id, sub_sector_id, planet_location_id, vehicle_id, title, description, goal FROM threads WHERE id = $1`
	row := r.db.QueryRow(query, id)
	var t Thread
	err := row.Scan(&t.ID, &t.UserID, &t.SubSectorID, &t.PlanetLocationID, &t.VehicleID, &t.Title, &t.Description, &t.Goal)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &t, err
}

func (r *explorationRepository) SaveBead(b *Bead, threadID uuid.UUID) error {
	query := `INSERT INTO beads (id, thread_id, type, title, description, visual_prompt) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, b.ID, threadID, b.Type, b.Title, b.Description, b.VisualPrompt)
	return err
}

func (r *explorationRepository) GetBeadsByThreadID(threadID uuid.UUID) ([]Bead, error) {
	query := `SELECT id, type, title, description, visual_prompt FROM beads WHERE thread_id = $1 ORDER BY created_at ASC`
	rows, err := r.db.Query(query, threadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beads []Bead
	for rows.Next() {
		var b Bead
		if err := rows.Scan(&b.ID, &b.Type, &b.Title, &b.Description, &b.VisualPrompt); err != nil {
			return nil, err
		}
		beads = append(beads, b)
	}
	return beads, nil
}
