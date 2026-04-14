package service

import (
	"abondoe/spond-assignment/internal/models"
	"abondoe/spond-assignment/internal/models/dto"
	"abondoe/spond-assignment/internal/repository"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type FormService interface {
	GetForm(ctx context.Context, id uuid.UUID) (*dto.GetFormResponse, error)
}

type formService struct {
	formRepo repository.FormRepository
}

func NewFormService(formRepo repository.FormRepository) FormService {
	return &formService{
		formRepo: formRepo,
	}
}

func (s *formService) GetForm(ctx context.Context, id uuid.UUID) (*dto.GetFormResponse, error) {
	formRow, err := s.formRepo.GetForm(ctx, id)
	if err == sql.ErrNoRows {
		return nil, ErrFormNotFound
	} else if err != nil {
		return nil, ErrDatabaseError
	}

	form := models.MapFormToDTO(formRow)

	return form, nil
}
