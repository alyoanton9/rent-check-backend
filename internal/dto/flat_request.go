package dto

import "rent-checklist/internal/models"

type FlatRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

func (flatRequest FlatRequest) ToModel() models.Flat {
	return models.Flat{
		Title:       flatRequest.Title,
		Description: flatRequest.Description,
		Address:     flatRequest.Address,
	}
}
