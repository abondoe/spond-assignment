package handler_test

import (
	"abondoe/spond-assignment/internal/models/dto"
	"context"

	"github.com/google/uuid"
)

type mockFormService struct {
	getFormFn func(ctx context.Context, id uuid.UUID) (*dto.GetFormResponse, error)
}

func (m *mockFormService) GetForm(ctx context.Context, id uuid.UUID) (*dto.GetFormResponse, error) {
	return m.getFormFn(ctx, id)
}

type mockRegistrationService struct {
	createRegistrationFn func(ctx context.Context, req dto.CreateRegistrationRequest) error
}

func (m *mockRegistrationService) CreateRegistration(ctx context.Context, req dto.CreateRegistrationRequest) error {
	return m.createRegistrationFn(ctx, req)
}
