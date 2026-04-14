// models.go
package dto

import (
	"abondoe/spond-assignment/internal/types"
	"time"
)

type CreateRegistrationRequest struct {
	FormId       types.CompactUUID `json:"formId"`
	MemberTypeId types.CompactUUID `json:"memberTypeId"`
	Name         string            `json:"name"`
	Email        string            `json:"email"`
	PhoneNumber  string            `json:"phoneNumber"`
	BirthDate    time.Time         `json:"birthDate"`
}
