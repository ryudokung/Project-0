package game

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	useCase UseCase
	repo    Repository
}

func NewHandler(useCase UseCase, repo Repository) *Handler {
	return &Handler{useCase: useCase, repo: repo}
}

func (h *Handler) GetPilotStats(w http.ResponseWriter, r *http.Request) {
	charIDStr := r.URL.Query().Get("character_id")
	if charIDStr == "" {
		http.Error(w, "character_id is required", http.StatusBadRequest)
		return
	}

	charID, err := uuid.Parse(charIDStr)
	if err != nil {
		http.Error(w, "invalid character_id", http.StatusBadRequest)
		return
	}

	stats, err := h.repo.GetPilotStats(charID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if stats == nil {
		http.Error(w, "stats not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
