package vehicle

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/auth/constants"
)

type Handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) MintStarter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Try to get userID from context (Security)
	userID, ok := r.Context().Value(constants.UserIDKey).(uuid.UUID)
	if !ok {
		// Fallback to query param for testing
		userIDStr := r.URL.Query().Get("user_id")
		var err error
		userID, err = uuid.Parse(userIDStr)
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusBadRequest)
			return
		}
	}

	vehicle, err := h.useCase.InitializeStarterPack(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicle)
}

func (h *Handler) ListVehicles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Try to get userID from context (Security)
	userID, ok := r.Context().Value(constants.UserIDKey).(uuid.UUID)
	if !ok {
		// Fallback to query param for testing if not using middleware
		userIDStr := r.URL.Query().Get("user_id")
		if userIDStr != "" {
			var err error
			userID, err = uuid.Parse(userIDStr)
			if err != nil {
				http.Error(w, "Invalid user_id", http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, "Unauthorized: missing user_id", http.StatusUnauthorized)
			return
		}
	}

	charIDStr := r.URL.Query().Get("character_id")
	if charIDStr != "" {
		charID, err := uuid.Parse(charIDStr)
		if err != nil {
			http.Error(w, "Invalid character_id", http.StatusBadRequest)
			return
		}
		vehicles, err := h.useCase.GetCharacterVehicles(charID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vehicles)
		return
	}

	vehicles, err := h.useCase.GetUserVehicles(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicles)
}

func (h *Handler) GetVehicleCP(w http.ResponseWriter, r *http.Request) {
	vehicleIDStr := r.URL.Query().Get("id")
	vehicleID, err := uuid.Parse(vehicleIDStr)
	if err != nil {
		http.Error(w, "Invalid vehicle id", http.StatusBadRequest)
		return
	}

	cp, err := h.useCase.GetVehicleCP(r.Context(), vehicleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"cp": cp})
}

func (h *Handler) EquipItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ItemID    uuid.UUID `json:"item_id"`
		VehicleID uuid.UUID `json:"vehicle_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.useCase.EquipItem(r.Context(), req.ItemID, req.VehicleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (h *Handler) UnequipItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ItemID uuid.UUID `json:"item_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.useCase.UnequipItem(r.Context(), req.ItemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// Item & DDS Endpoints

func (h *Handler) ListItems(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(constants.UserIDKey).(uuid.UUID)
	if !ok {
		userIDStr := r.URL.Query().Get("user_id")
		var err error
		userID, err = uuid.Parse(userIDStr)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}

	items, err := h.useCase.GetItems(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) ApplyDamage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ItemID uuid.UUID `json:"item_id"`
		Damage int       `json:"damage"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	item, err := h.useCase.ApplyDamage(r.Context(), req.ItemID, req.Damage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) Repair(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ItemID uuid.UUID `json:"item_id"`
		Amount int       `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	item, err := h.useCase.RepairItem(r.Context(), req.ItemID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}
