package dto

import "rent-checklist/internal/model"

type FlatRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

func ToModel(flatRequest FlatRequest) model.Flat {
	return model.Flat{
		Title:       flatRequest.Title,
		Description: flatRequest.Description,
		Address:     flatRequest.Address,
	}
}
