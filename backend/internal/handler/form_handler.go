package handler

import (
	"abondoe/spond-assignment/internal/service"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type FormHandler struct {
	formService service.FormService
}

func NewFormHandler(formService service.FormService) *FormHandler {
	return &FormHandler{
		formService: formService,
	}
}

func (h *FormHandler) GetForm(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	formId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid form ID", http.StatusBadRequest)
		return
	}

	form, err := h.formService.GetForm(r.Context(), formId)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrFormNotFound):
			respondWithError(w, http.StatusNotFound, "Form not found")
		case errors.Is(err, service.ErrDatabaseError):
			respondWithError(w, http.StatusInternalServerError, "Database error")
		default:
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
	}

	respondWithJSON(w, http.StatusOK, form)
}
