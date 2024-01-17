package model

import (
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/entity"
)

type Item struct {
	Id          uint64
	UserId      uint64
	Title       string
	Description string
	Hide        bool
}

func DtoToItem(itemRequest dto.ItemRequest, userId uint64) Item {
	return Item{
		UserId:      userId,
		Title:       itemRequest.Title,
		Description: itemRequest.Description,
	}
}

func ItemToDto(item Item) dto.ItemResponse {
	return dto.ItemResponse{
		Id:          item.Id,
		Title:       item.Title,
		Description: item.Description,
	}
}

func EntityToItem(item entity.Item) Item {
	return Item{
		Id:          item.Id,
		UserId:      item.UserId,
		Title:       item.Title,
		Description: item.Description,
		Hide:        item.Hide,
	}
}

func ItemToEntity(item Item) entity.Item {
	return entity.Item{
		Id:          item.Id,
		UserId:      item.UserId,
		Title:       item.Title,
		Description: item.Description,
		Hide:        item.Hide,
	}
}
