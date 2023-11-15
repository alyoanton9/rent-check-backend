package dto

import "rent-checklist/internal/models"

type FlatResponse struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

func (flatResponse *FlatResponse) FromModel(flat models.Flat) {
	flatResponse.Id = flat.Id
	flatResponse.Title = flat.Title
	flatResponse.Description = flat.Description
	flatResponse.Address = flat.Address
}
