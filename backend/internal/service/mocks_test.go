package service_test

import (
	"abondoe/spond-assignment/internal/models/db"
	"context"

	"github.com/google/uuid"
)

type mockFormRepository struct {
	form *db.Form
	err  error
}

func (m *mockFormRepository) GetForm(ctx context.Context, id uuid.UUID) (*db.Form, error) {
	return m.form, m.err
}

type mockRegistrationRepository struct {
	err error
}

func (m *mockRegistrationRepository) CreateRegistration(ctx context.Context, registration db.Registration) error {
	return m.err
}
