// models.go
package db

import (
	"time"

	"github.com/google/uuid"
)

type Registration struct {
	FormId       uuid.UUID
	MemberTypeId uuid.UUID
	Name         string
	Email        string
	PhoneNumber  string
	BirthDate    time.Time
	CreatedAt    time.Time
}
