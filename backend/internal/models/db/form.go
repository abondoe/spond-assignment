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
	// Litt rart at ClubId ikke er en UUID her - jeg ville kallet den clubName eller lignende
	ClubId            string
	MemberTypes       []MemberType
	FormId            uuid.UUID
	Title             string
	RegistrationOpens time.Time
	createdAt         time.Time
}
