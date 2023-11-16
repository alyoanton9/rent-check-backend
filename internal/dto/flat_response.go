package dto

import "rent-checklist/internal/model"

type FlatResponse struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

func FromModel(flat model.Flat) FlatResponse {
	return FlatResponse{
		Id:          flat.Id,
		Title:       flat.Title,
		Description: flat.Description,
		Address:     flat.Address,
	}
}
