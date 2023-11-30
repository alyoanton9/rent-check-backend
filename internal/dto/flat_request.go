package dto

import "rent-checklist-backend/internal/model"

type FlatRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

func ToModel(flatRequest FlatRequest, ownerId string) model.Flat {
	return model.Flat{
		Title:       flatRequest.Title,
		Description: flatRequest.Description,
		Address:     flatRequest.Address,
		OwnerId:     ownerId,
	}
}
