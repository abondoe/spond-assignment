package handler_test

import (
	"abondoe/spond-assignment/internal/handler"
	"abondoe/spond-assignment/internal/models/dto"
	"abondoe/spond-assignment/internal/service"
	"abondoe/spond-assignment/internal/types"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func TestGetFormHandler(t *testing.T) {
	tests := []struct {
		name           string
		idParam        string
		mockBehavior   func(ctx context.Context, id uuid.UUID) (*dto.GetFormResponse, error)
		expectedStatus int
		validate       func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name:           "Invalid UUID format gives 400",
			idParam:        "invalid-uuid",
			mockBehavior:   nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "Form not found gives 404",
			idParam: uuid.New().String(),
			mockBehavior: func(ctx context.Context, id uuid.UUID) (*dto.GetFormResponse, error) {
				return nil, service.ErrFormNotFound
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:    "Successful retrieval gives 200 and correct JSON",
			idParam: "550e8400-e29b-41d4-a716-446655440000",
			mockBehavior: func(ctx context.Context, id uuid.UUID) (*dto.GetFormResponse, error) {
				return &dto.GetFormResponse{FormId: types.CompactUUID(id)}, nil
			},
			expectedStatus: http.StatusOK,
			validate: func(t *testing.T, w *httptest.ResponseRecorder) {
				assertJSONContentType(t, w)

				// Check that body can be parsed and contains data
				var resp dto.GetFormResponse
				assertJSONBody(t, w.Body, &resp)

				if resp.FormId == types.CompactUUID(uuid.Nil) {
					t.Error("forventet FormId i responsen, fikk tom streng")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockService := &mockFormService{getFormFn: tt.mockBehavior}
			h := handler.NewFormHandler(mockService)

			// Prepare request
			req := httptest.NewRequest(http.MethodGet, "/forms/"+tt.idParam, nil)

			// Important for Go 1.22+
			req.SetPathValue("id", tt.idParam)

			w := httptest.NewRecorder()

			// Execute
			h.GetForm(w, req)

			// Assert status
			if w.Code != tt.expectedStatus {
				t.Errorf("forventet status %d, fikk %d", tt.expectedStatus, w.Code)
			}

			// Validate response body if needed
			if tt.validate != nil {
				tt.validate(t, w)
			}
		})
	}
}
