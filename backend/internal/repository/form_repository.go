package repository

import (
	"abondoe/spond-assignment/internal/models/db"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type FormRepository interface {
	GetForm(ctx context.Context, id uuid.UUID) (*db.Form, error)
}

type formRepository struct{}

var forms = []db.Form{createForm()}

func NewFormRepository() FormRepository {

	return &formRepository{}
}

func (r *formRepository) GetForm(ctx context.Context, id uuid.UUID) (*db.Form, error) {

	for _, form := range forms {
		if form.FormId == id {
			return &form, nil
		}
	}

	return nil, sql.ErrNoRows
}

func createForm() db.Form {

	amId, _ := uuid.Parse("8FE4113D4E4020E0DCF887803A886981")
	smId, _ := uuid.Parse("4237C55C5CC3B4B082CBF2540612778E")
	formId, _ := uuid.Parse("B171388180BC457D9887AD92B6CCFC86")

	memberTypes := []db.MemberType{
		{
			Id:   amId,
			Name: "Active Member",
		},
		{
			Id:   smId,
			Name: "Social Member",
		},
	}

	registrationOpens := time.Date(2024, 12, 16, 0, 0, 0, 0, time.UTC)

	return db.Form{
		ClubId:            "britsport",
		MemberTypes:       memberTypes,
		FormId:            formId,
		Title:             "Coding camp summer 2025",
		RegistrationOpens: registrationOpens,
	}
}
