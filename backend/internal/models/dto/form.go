// models.go
package dto

import (
	"abondoe/spond-assignment/internal/types"
	"time"
)

type MemberType struct {
	Id   types.CompactUUID `json:"id"`
	Name string            `json:"name"`
}

type GetFormResponse struct {
	FormId            types.CompactUUID `json:"formId"`
	ClubId            string            `json:"clubId"`
	MemberTypes       []MemberType      `json:"memberTypes"`
	Title             string            `json:"title"`
	RegistrationOpens time.Time         `json:"registrationOpens"`
}
