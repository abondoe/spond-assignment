package service_test

import (
	"abondoe/spond-assignment/internal/models/db"
	"abondoe/spond-assignment/internal/models/dto"
	"abondoe/spond-assignment/internal/types"
	"time"

	"github.com/google/uuid"
)

func validFormId() uuid.UUID {
	return uuid.MustParse("B171388180BC457D9887AD92B6CCFC86")
}

func validActiveMemberTypeId() uuid.UUID {
	return uuid.MustParse("8FE4113D4E4020E0DCF887803A886981")
}

func validSocialMemberTypeId() uuid.UUID {
	return uuid.MustParse("4237C55C5CC3B4B082CBF2540612778E")
}

func validCompactFormId() types.CompactUUID {
	return types.CompactUUID(validFormId())
}

func validCompactActiveMemberTypeId() types.CompactUUID {
	return types.CompactUUID(validActiveMemberTypeId())
}

func validCreateRegistrationRequest() dto.CreateRegistrationRequest {
	return dto.CreateRegistrationRequest{
		FormId:       validCompactFormId(),
		MemberTypeId: validCompactActiveMemberTypeId(),
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		PhoneNumber:  "12345678",
		BirthDate:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func validForm() *db.Form {
	return &db.Form{
		FormId: validFormId(),
		Title:  "Awesome Camp",
		ClubId: "Awesome Club",
		MemberTypes: []db.MemberType{
			{
				Id:   validActiveMemberTypeId(),
				Name: "Active Member",
			},
			{
				Id:   validSocialMemberTypeId(),
				Name: "Social Member",
			},
		},
		RegistrationOpens: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}
