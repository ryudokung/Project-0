package combat

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/auth/constants"
	"github.com/ryudokung/Project-0/backend/internal/vehicle"
	"github.com/ryudokung/Project-0/backend/internal/game"
)

type Handler struct {
	service     *Service
	vehicleRepo vehicle.Repository
	gameRepo    game.Repository
}

func NewHandler(service *Service, vehicleRepo vehicle.Repository, gameRepo game.Repository) *Handler {
	return &Handler{
		service:     service,
		vehicleRepo: vehicleRepo,
		gameRepo:    gameRepo,
	}
}

type BattleRequest struct {
	AttackerVehicleID string `json:"attacker_vehicle_id"`
	DefenderVehicleID string `json:"defender_vehicle_id"`
	DamageType        string `json:"damage_type"`
}

func (h *Handler) SimulateAttack(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 0. Get User from Context (Security)
	userID, ok := r.Context().Value(constants.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req BattleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	attackerUUID, err := uuid.Parse(req.AttackerVehicleID)
	if err != nil {
		http.Error(w, "Invalid attacker ID format", http.StatusBadRequest)
		return
	}
	defenderUUID, err := uuid.Parse(req.DefenderVehicleID)
	if err != nil {
		http.Error(w, "Invalid defender ID format", http.StatusBadRequest)
		return
	}

	// 1. Fetch Vehicles
	attacker, err := h.vehicleRepo.GetByID(r.Context(), attackerUUID)
	if err != nil {
		http.Error(w, "Error fetching attacker: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if attacker == nil {
		http.Error(w, "Attacker vehicle not found", http.StatusNotFound)
		return
	}

	// Ownership Check (Anti-Cheat)
	if attacker.OwnerID != userID {
		http.Error(w, "You do not own this vehicle", http.StatusForbidden)
		return
	}

	defender, err := h.vehicleRepo.GetByID(r.Context(), defenderUUID)
	if err != nil {
		http.Error(w, "Error fetching defender: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if defender == nil {
		http.Error(w, "Defender vehicle not found", http.StatusNotFound)
		return
	}

	// Check if defender is already dead (Anti-Cheat/Race Condition)
	if defender.Stats.HP <= 0 {
		http.Error(w, "Target is already destroyed", http.StatusBadRequest)
		return
	}

	// 2. Fetch Parts (Equipped)
	attackerParts, _ := h.vehicleRepo.GetPartsByVehicleID(attackerUUID)
	defenderParts, _ := h.vehicleRepo.GetPartsByVehicleID(defenderUUID)

	// 3. Fetch Pilot Stats (for Resonance bonus)
	attackerPilot, _ := h.gameRepo.GetActivePilotStats(attacker.OwnerID)
	defenderPilot, _ := h.gameRepo.GetActivePilotStats(defender.OwnerID)

	// 4. Map to Combat Stats
	attackerStats := h.service.MapVehicleToUnitStats(attacker, attackerParts, attackerPilot)
	defenderStats := h.service.MapVehicleToUnitStats(defender, defenderParts, defenderPilot)

	// 5. Execute Attack
	result := h.service.ExecuteAttack(attackerStats, defenderStats, DamageType(req.DamageType))

	// 6. Persist State (Anti-Cheat)
	newHP := defender.Stats.HP - result.FinalDamage
	if newHP < 0 {
		newHP = 0
	}
	
	err = h.vehicleRepo.UpdateHP(r.Context(), defenderUUID, newHP)
	if err != nil {
		http.Error(w, "Failed to update defender HP", http.StatusInternalServerError)
		return
	}

	// Update local stats for response
	defenderStats.HP = newHP

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"attacker_stats": attackerStats,
		"defender_stats": defenderStats,
		"result":         result,
	})
}
