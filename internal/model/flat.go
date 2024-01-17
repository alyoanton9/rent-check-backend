package model

import (
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/entity"
)

type Flat struct {
	Id          uint64
	Title       string
	Description string
	Address     string
	OwnerId     uint64
}

func DtoToFlat(flatRequest dto.FlatRequest, ownerId uint64) Flat {
	return Flat{
		Title:       flatRequest.Title,
		Description: flatRequest.Description,
		Address:     flatRequest.Address,
		OwnerId:     ownerId,
	}
}

func FlatToDto(flat Flat) dto.FlatResponse {
	return dto.FlatResponse{
		Id:          flat.Id,
		Title:       flat.Title,
		Description: flat.Description,
		Address:     flat.Address,
	}
}

func EntityToFlat(flat entity.Flat) Flat {
	return Flat{
		Id:          flat.Id,
		Title:       flat.Title,
		Description: flat.Description,
		Address:     flat.Address,
		OwnerId:     flat.OwnerId,
	}
}

func FlatToEntity(flat Flat) entity.Flat {
	return entity.Flat{
		Id:          flat.Id,
		Title:       flat.Title,
		Description: flat.Description,
		Address:     flat.Address,
		OwnerId:     flat.OwnerId,
	}
}
