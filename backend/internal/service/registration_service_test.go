package service_test

import (
	"abondoe/spond-assignment/internal/service"
	"context"
	"database/sql"
	"testing"

	"github.com/lib/pq"
)

func TestCreateRegistration(t *testing.T) {
	tests := []struct {
		name     string
		formRepo mockFormRepository
		regRepo  mockRegistrationRepository
		wantErr  error
	}{
		{
			name:     "happy path",
			formRepo: mockFormRepository{form: validForm()},
			regRepo:  mockRegistrationRepository{},
			wantErr:  nil,
		},
		{
			name:     "form not found",
			formRepo: mockFormRepository{err: sql.ErrNoRows},
			regRepo:  mockRegistrationRepository{},
			wantErr:  service.ErrFormNotFound,
		},
		{
			name:     "duplicate registration",
			formRepo: mockFormRepository{form: validForm()},
			regRepo:  mockRegistrationRepository{err: &pq.Error{Code: "23505"}},
			wantErr:  service.ErrDuplicateRegistration,
		},
		{
			name:     "database error",
			formRepo: mockFormRepository{form: validForm()},
			regRepo:  mockRegistrationRepository{err: sql.ErrConnDone},
			wantErr:  service.ErrDatabaseError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			service := service.NewRegistrationService(&tt.regRepo, &tt.formRepo)

			// Execute
			err := service.CreateRegistration(context.Background(), validCreateRegistrationRequest())

			// Assert status
			if err != tt.wantErr {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
