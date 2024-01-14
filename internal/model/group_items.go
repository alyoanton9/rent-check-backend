package model

import (
	"github.com/samber/lo"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/entity"
	"sort"
)

type GroupItems struct {
	GroupId uint64
	Items   []ItemWithStatus
}

type ItemWithStatus struct {
	Item   Item
	Status Status
}

func EntitiesToGroupItems(groupItemsRecords []entity.GroupItem) []GroupItems {
	itemsByGroupId := lo.GroupBy(groupItemsRecords, func(groupItemsRecord entity.GroupItem) uint64 {
		return groupItemsRecord.GroupId
	})

	groupItemsList := make([]GroupItems, 0)
	for groupId, recordsList := range itemsByGroupId {
		items := make([]ItemWithStatus, 0)
		if !(len(recordsList) == 1 && recordsList[0].Status == "") {
			items = lo.Map(recordsList, func(record entity.GroupItem, _ int) ItemWithStatus {
				status, _ := ParseStatus(record.Status)
				return ItemWithStatus{
					Item: Item{
						Id:          record.Id,
						UserId:      record.UserId,
						Title:       record.Title,
						Description: record.Description,
						Hide:        record.Hide,
					},
					Status: status,
				}
			})
		}

		groupItemsList = append(groupItemsList, GroupItems{GroupId: groupId, Items: items})
	}

	return groupItemsList
}

func GroupItemsToDto(groupItemsList []GroupItems) []dto.GroupItemsResponse {
	groupItemsResponseList := lo.Map(groupItemsList, func(groupItems GroupItems, _ int) dto.GroupItemsResponse {
		items := lo.Map(groupItems.Items, func(itemWithStatus ItemWithStatus, _ int) dto.ItemWithStatus {
			return dto.ItemWithStatus{
				Status: itemWithStatus.Status.String(),
				Item: dto.ItemResponse{
					Id:          itemWithStatus.Item.Id,
					Title:       itemWithStatus.Item.Title,
					Description: itemWithStatus.Item.Description,
				},
			}
		})

		return dto.GroupItemsResponse{
			GroupId: groupItems.GroupId,
			Items:   items,
		}
	})

	return groupItemsResponseList
}

func SortGroupItemsByGroupId(groupItemsList []GroupItems) []GroupItems {
	sort.Slice(groupItemsList, func(i, j int) bool {
		return groupItemsList[i].GroupId < groupItemsList[j].GroupId
	})

	return groupItemsList
}
