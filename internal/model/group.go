package model

import (
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/entity"
)

type Group struct {
	Id     uint64
	UserId string
	Title  string
	Hide   bool
}

func DtoToGroup(groupRequest dto.GroupRequest, userId string) Group {
	return Group{
		Title:  groupRequest.Title,
		UserId: userId,
	}
}

func GroupToDto(group Group) dto.GroupResponse {
	return dto.GroupResponse{
		Id:    group.Id,
		Title: group.Title,
	}
}

func EntityToGroup(group entity.Group) Group {
	return Group{
		Id:     group.Id,
		UserId: group.UserId,
		Title:  group.Title,
		Hide:   group.Hide,
	}
}

func GroupToEntity(group Group) entity.Group {
	return entity.Group{
		Id:     group.Id,
		UserId: group.UserId,
		Title:  group.Title,
		Hide:   group.Hide,
	}
}
