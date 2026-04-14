package repository

import (
	"abondoe/spond-assignment/internal/models/db"
	"context"
	"database/sql"
)

type RegistrationRepository interface {
	CreateRegistration(ctx context.Context, registration db.Registration) error
}

type registrationRepository struct {
	db *sql.DB
}

func NewRegistrationRepository(db *sql.DB) RegistrationRepository {
	return &registrationRepository{db: db}
}

func (r *registrationRepository) CreateRegistration(ctx context.Context, registration db.Registration) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO registrations (form_id, member_type_id, name, email, phone_number, birth_date, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		registration.FormId,
		registration.MemberTypeId,
		registration.Name,
		registration.Email,
		registration.PhoneNumber,
		registration.BirthDate,
		registration.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
