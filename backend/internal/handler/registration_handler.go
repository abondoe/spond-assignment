package handler

import (
	"abondoe/spond-assignment/internal/models/dto"
	"abondoe/spond-assignment/internal/service"
	"encoding/json"
	"errors"
	"net/http"
)

type RegistrationHandler struct {
	registrationService service.RegistrationService
}

func NewRegistrationHandler(registrationService service.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{
		registrationService: registrationService,
	}
}

func (h *RegistrationHandler) CreateRegistration(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.registrationService.CreateRegistration(r.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrFormNotFound):
			respondWithError(w, http.StatusNotFound, "Form not found")
		case errors.Is(err, service.ErrDatabaseError):
			respondWithError(w, http.StatusInternalServerError, "Database error")
		case errors.Is(err, service.ErrDuplicateRegistration):
			respondWithError(w, http.StatusConflict, "Duplicate registration")
		default:
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
