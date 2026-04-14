package service_test

import (
	"abondoe/spond-assignment/internal/models/dto"
	"abondoe/spond-assignment/internal/service"
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestGetForm(t *testing.T) {
	tests := []struct {
		name     string
		formId   uuid.UUID
		formRepo mockFormRepository
		wantErr  error
		validate func(t *testing.T, resp *dto.GetFormResponse)
	}{
		{
			name:     "happy path",
			formId:   validFormId(),
			formRepo: mockFormRepository{form: validForm()},
			wantErr:  nil,
			validate: func(t *testing.T, resp *dto.GetFormResponse) {
				formId := validCompactFormId()
				if resp == nil {
					t.Fatal("expected response, got nil")
				}
				if resp.FormId != formId {
					t.Errorf("expected form ID %v, got %v", formId, resp.FormId)
				}
			},
		},
		{
			name:     "form not found",
			formId:   validFormId(),
			formRepo: mockFormRepository{err: sql.ErrNoRows},
			wantErr:  service.ErrFormNotFound,
		},
		{
			name:     "database error",
			formId:   validFormId(),
			formRepo: mockFormRepository{err: sql.ErrConnDone},
			wantErr:  service.ErrDatabaseError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := service.NewFormService(&tt.formRepo)

			resp, err := service.GetForm(context.Background(), tt.formId)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}

			if tt.validate != nil {
				tt.validate(t, resp)
			}
		})
	}
}
