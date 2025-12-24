package exploration

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
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
	var req struct {
		UserID           uuid.UUID  `json:"user_id"`
		SubSectorID      uuid.UUID  `json:"sub_sector_id"`
		PlanetLocationID *uuid.UUID `json:"planet_location_id"`
		VehicleID        uuid.UUID  `json:"vehicle_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expedition, err := h.service.StartExploration(r.Context(), req.UserID, req.SubSectorID, req.PlanetLocationID, req.VehicleID)
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

	// Fetch pilot stats
	pilotStats, err := h.service.gameRepo.GetPilotStats(expedition.UserID)
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
	var req struct {
		ExpeditionID uuid.UUID `json:"expedition_id"`
		VehicleID    uuid.UUID `json:"vehicle_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encounter, err := h.service.StringNewEncounter(r.Context(), req.ExpeditionID, req.VehicleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch updated pilot stats
	expedition, err := h.service.repo.GetExpeditionByID(req.ExpeditionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	pilotStats, err := h.service.gameRepo.GetPilotStats(expedition.UserID)
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
