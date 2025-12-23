package combat

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/mech"
	"github.com/ryudokung/Project-0/backend/internal/game"
)

type Handler struct {
	service    *Service
	mechRepo   mech.Repository
	gameRepo   game.Repository
}

func NewHandler(service *Service, mechRepo mech.Repository, gameRepo game.Repository) *Handler {
	return &Handler{
		service:  service,
		mechRepo: mechRepo,
		gameRepo: gameRepo,
	}
}

type BattleRequest struct {
	AttackerMechID string `json:"attacker_mech_id"`
	DefenderMechID string `json:"defender_mech_id"`
	DamageType     string `json:"damage_type"`
}

func (h *Handler) SimulateAttack(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req BattleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	attackerUUID, err := uuid.Parse(req.AttackerMechID)
	if err != nil {
		http.Error(w, "Invalid attacker ID format", http.StatusBadRequest)
		return
	}
	defenderUUID, err := uuid.Parse(req.DefenderMechID)
	if err != nil {
		http.Error(w, "Invalid defender ID format", http.StatusBadRequest)
		return
	}

	// 1. Fetch Mechs
	attacker, err := h.mechRepo.GetByID(r.Context(), attackerUUID)
	if err != nil {
		http.Error(w, "Error fetching attacker: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if attacker == nil {
		http.Error(w, "Attacker mech not found", http.StatusNotFound)
		return
	}

	defender, err := h.mechRepo.GetByID(r.Context(), defenderUUID)
	if err != nil {
		http.Error(w, "Error fetching defender: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if defender == nil {
		http.Error(w, "Defender mech not found", http.StatusNotFound)
		return
	}

	// 2. Fetch Parts (Equipped)
	attackerParts, _ := h.mechRepo.GetPartsByMechID(attackerUUID)
	defenderParts, _ := h.mechRepo.GetPartsByMechID(defenderUUID)

	// 3. Fetch Pilot Stats (for Resonance bonus)
	attackerPilot, _ := h.gameRepo.GetPilotStats(attacker.OwnerID)
	defenderPilot, _ := h.gameRepo.GetPilotStats(defender.OwnerID)

	// 4. Map to Combat Stats
	attackerStats := h.service.MapMechToUnitStats(attacker, attackerParts, attackerPilot)
	defenderStats := h.service.MapMechToUnitStats(defender, defenderParts, defenderPilot)

	// 5. Execute Attack
	result := h.service.ExecuteAttack(attackerStats, defenderStats, DamageType(req.DamageType))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"attacker_stats": attackerStats,
		"defender_stats": defenderStats,
		"result":         result,
	})
}
