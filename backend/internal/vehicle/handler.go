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

	vehicle, err := h.useCase.MintStarterVehicle(userID)
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
