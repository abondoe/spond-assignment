// models.go
package db

import (
	"time"

	"github.com/google/uuid"
)

type MemberType struct {
	Id   uuid.UUID
	Name string
}

type Form struct {
	// Odd that ClubId isn´t a UUID - I would have called it clubName or similar
	ClubId            string
	MemberTypes       []MemberType
	FormId            uuid.UUID
	Title             string
	RegistrationOpens time.Time
}
