package models

import (
	"abondoe/spond-assignment/internal/models/db"
	"abondoe/spond-assignment/internal/models/dto"
	"abondoe/spond-assignment/internal/types"
)

func MapFormToDTO(f *db.Form) *dto.GetFormResponse {
	memberTypes := make([]dto.MemberType, len(f.MemberTypes))
	for i, mt := range f.MemberTypes {
		memberTypes[i] = dto.MemberType{
			Id:   types.CompactUUID(mt.Id),
			Name: mt.Name,
		}
	}

	return &dto.GetFormResponse{
		ClubId:            f.ClubId,
		MemberTypes:       memberTypes,
		FormId:            types.CompactUUID(f.FormId),
		Title:             f.Title,
		RegistrationOpens: f.RegistrationOpens,
	}
}
