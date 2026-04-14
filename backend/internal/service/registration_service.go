package service

import (
	"abondoe/spond-assignment/internal/models/db"
	"abondoe/spond-assignment/internal/models/dto"
	"abondoe/spond-assignment/internal/repository"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type RegistrationService interface {
	CreateRegistration(ctx context.Context, req dto.CreateRegistrationRequest) error
}

type registrationService struct {
	registrationRepo repository.RegistrationRepository
	formRepo         repository.FormRepository
}

func NewRegistrationService(registrationRepo repository.RegistrationRepository, formRepo repository.FormRepository) RegistrationService {
	return &registrationService{
		registrationRepo: registrationRepo,
		formRepo:         formRepo,
	}
}

func (s *registrationService) CreateRegistration(ctx context.Context, req dto.CreateRegistrationRequest) error {
	formId := uuid.UUID(req.FormId)
	memberTypeId := uuid.UUID(req.MemberTypeId)

	// Sjekk at skjemaet eksisterer
	_, err := s.formRepo.GetForm(ctx, formId)
	if err != nil {
		return ErrFormNotFound
	}

	registration := db.Registration{
		FormId:       formId,
		MemberTypeId: memberTypeId,
		Name:         req.Name,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		BirthDate:    req.BirthDate,
		CreatedAt:    time.Now().UTC().Truncate(time.Second), // Truncate to second for consistent formatting with form data
	}

	if err := s.registrationRepo.CreateRegistration(ctx, registration); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrDuplicateRegistration
		}
		return ErrDatabaseError
	}

	return nil
}
