package handler_test

import (
	"abondoe/spond-assignment/internal/handler"
	"abondoe/spond-assignment/internal/models/dto"
	"abondoe/spond-assignment/internal/service"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRegistrationHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    any
		mockBehavior   func(ctx context.Context, req dto.CreateRegistrationRequest) error
		expectedStatus int
	}{
		{
			name:           "Invalid JSON body gives 400",
			requestBody:    "invalid json",
			mockBehavior:   nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "Form not found gives 404",
			requestBody: validCreateRegistrationRequest(),
			mockBehavior: func(ctx context.Context, req dto.CreateRegistrationRequest) error {
				return service.ErrFormNotFound
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:        "Database error gives 500",
			requestBody: validCreateRegistrationRequest(),
			mockBehavior: func(ctx context.Context, req dto.CreateRegistrationRequest) error {
				return service.ErrDatabaseError
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:        "Duplicate registration gives 409",
			requestBody: validCreateRegistrationRequest(),
			mockBehavior: func(ctx context.Context, req dto.CreateRegistrationRequest) error {
				return service.ErrDuplicateRegistration
			},
			expectedStatus: http.StatusConflict,
		},
		{
			name:        "Successful registration gives 201 and data",
			requestBody: validCreateRegistrationRequest(),
			mockBehavior: func(ctx context.Context, req dto.CreateRegistrationRequest) error {
				return nil
			},
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockSvc := &mockRegistrationService{createRegistrationFn: tt.mockBehavior}
			h := handler.NewRegistrationHandler(mockSvc)

			// Prepare body
			var body []byte
			if s, ok := tt.requestBody.(string); ok {
				body = []byte(s)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			// Prepare request
			req := httptest.NewRequest(http.MethodPost, "/registrations", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			// Execute
			h.CreateRegistration(w, req)

			// Assert
			if w.Code != tt.expectedStatus {
				t.Errorf("forventet status %d, fikk %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
