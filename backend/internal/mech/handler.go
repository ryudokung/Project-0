package mech

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

	mech, err := h.useCase.MintStarterMech(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mech)
}

func (h *Handler) ListMechs(w http.ResponseWriter, r *http.Request) {
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
		mechs, err := h.useCase.GetCharacterMechs(charID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mechs)
		return
	}

	mechs, err := h.useCase.GetUserMechs(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mechs)
}
