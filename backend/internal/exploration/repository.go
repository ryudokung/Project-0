package exploration

import (
	"database/sql"
	"github.com/google/uuid"
)

type explorationRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &explorationRepository{db: db}
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

func (r *explorationRepository) GetThreadByID(id uuid.UUID) (*Thread, error) {
	query := `SELECT id, title, description, goal FROM threads WHERE id = $1`
	row := r.db.QueryRow(query, id)
	var t Thread
	err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Goal)
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
