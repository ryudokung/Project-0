package exploration

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/auth/constants"
	"github.com/ryudokung/Project-0/backend/internal/game"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	expeditionIDStr := r.URL.Query().Get("expedition_id")
	expeditionID, err := uuid.Parse(expeditionIDStr)
	if err != nil {
		http.Error(w, "Invalid expedition_id", http.StatusBadRequest)
		return
	}

	nodes, err := h.service.repo.GetNodesByExpeditionID(expeditionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nodes)
}

func (h *Handler) ResolveNode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		NodeID uuid.UUID `json:"node_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	node, err := h.service.ResolveNode(r.Context(), req.NodeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(node)
}

func (h *Handler) ResolveChoice(w http.ResponseWriter, r *http.Request) {
	var req struct {
		NodeID uuid.UUID `json:"node_id"`
		Choice string    `json:"choice"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	node, err := h.service.ResolveNodeChoice(r.Context(), req.NodeID, req.Choice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch updated stats to return to frontend
	expedition, _ := h.service.repo.GetExpeditionByID(node.ExpeditionID)
	stats, _ := h.service.gameRepo.GetActivePilotStats(expedition.UserID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"node":        node,
		"pilot_stats": stats,
	})
}

func (h *Handler) GetUniverseMap(w http.ResponseWriter, r *http.Request) {
	sectors, err := h.service.repo.GetAllSectors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type SubSectorResponse struct {
		SubSector
		Locations []PlanetLocation `json:"locations,omitempty"`
	}

	type SectorResponse struct {
		Sector
		SubSectors []SubSectorResponse `json:"subSectors"`
	}

	var response []SectorResponse
	for _, s := range sectors {
		subSectors, err := h.service.repo.GetSubSectorsBySectorID(s.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var subSectorResponses []SubSectorResponse
		for _, ss := range subSectors {
			var locations []PlanetLocation
			if ss.Type == "PLANET" {
				locations, err = h.service.repo.GetPlanetLocationsBySubSectorID(ss.ID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			subSectorResponses = append(subSectorResponses, SubSectorResponse{
				SubSector: ss,
				Locations: locations,
			})
		}

		response = append(response, SectorResponse{
			Sector:     s,
			SubSectors: subSectorResponses,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) StartExploration(w http.ResponseWriter, r *http.Request) {
	// 0. Get User from Context (Security)
	userID, ok := r.Context().Value(constants.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req struct {
		SubSectorID      uuid.UUID  `json:"sub_sector_id"`
		PlanetLocationID *uuid.UUID `json:"planet_location_id"`
		VehicleID        *uuid.UUID `json:"vehicle_id"`
		BlueprintID      string     `json:"blueprint_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vID := uuid.Nil
	if req.VehicleID != nil {
		vID = *req.VehicleID
	}

	var expedition *Expedition
	var err error

	if req.BlueprintID != "" {
		expedition, err = h.service.CreateHandcraftedExpedition(r.Context(), userID, req.BlueprintID, &vID)
	} else {
		expedition, err = h.service.StartExploration(r.Context(), userID, req.SubSectorID, req.PlanetLocationID, vID)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch encounters for the expedition
	encounters, err := h.service.repo.GetEncountersByExpeditionID(expedition.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if encounters == nil {
		encounters = []Encounter{}
	}

	// Fetch pilot stats
	pilotStats, err := h.service.gameRepo.GetActivePilotStats(expedition.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Expedition *Expedition      `json:"expedition"`
		Encounters []Encounter      `json:"encounters"`
		PilotStats *game.PilotStats `json:"pilot_stats"`
	}{
		Expedition: expedition,
		Encounters: encounters,
		PilotStats: pilotStats,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) AdvanceTimeline(w http.ResponseWriter, r *http.Request) {
	// 0. Get User from Context (Security)
	userID, ok := r.Context().Value(constants.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req struct {
		ExpeditionID uuid.UUID `json:"expedition_id"`
		VehicleID    uuid.UUID `json:"vehicle_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 1. Verify Expedition Ownership (Anti-Cheat)
	expedition, err := h.service.repo.GetExpeditionByID(req.ExpeditionID)
	if err != nil {
		http.Error(w, "Expedition not found", http.StatusNotFound)
		return
	}
	if expedition.UserID != userID {
		http.Error(w, "You do not own this expedition", http.StatusForbidden)
		return
	}

	encounter, err := h.service.GenerateNewEncounter(r.Context(), req.ExpeditionID, req.VehicleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch updated pilot stats
	expedition, err = h.service.repo.GetExpeditionByID(req.ExpeditionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	pilotStats, err := h.service.gameRepo.GetActivePilotStats(expedition.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Encounter  *Encounter       `json:"encounter"`
		PilotStats *game.PilotStats `json:"pilot_stats"`
	}{
		Encounter:  encounter,
		PilotStats: pilotStats,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
