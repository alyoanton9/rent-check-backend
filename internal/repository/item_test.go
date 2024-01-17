package repository

import (
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
	"testing"
)

func testItemRepository(t *testing.T, repo ItemRepository, userId uint64) {
	var userIdNotFound uint64 = 111
	var flatId uint64 = 2
	var groupIds = []uint64{2, 3, 1}

	items := []model.Item{
		{
			Title:  "a",
			UserId: userId,
		},
		{
			Title:  "b",
			UserId: userId,
		},
		{
			Title:  "c",
			UserId: userId,
		},
		{
			Title:  "",
			UserId: userIdNotFound,
		},
	}

	testCreateItem(t, repo, &items)

	testGetItems(t, repo, &items, userId, userIdNotFound)

	testUpdateItems(t, repo, &items, userId, userIdNotFound)

	testHideItem(t, repo, &items, userId, userIdNotFound)

	testAddItemToGroup(t, repo, &items, flatId, groupIds, userId, userIdNotFound)

	testUpdateItemStatus(t, repo, &items, flatId, groupIds, userId)

	testGetFlatItems(t, repo, &items, flatId, groupIds, userId)

	testDeleteItemFromGroup(t, repo, &items, flatId, groupIds, userId, userIdNotFound)

}

// TODO add test cases after un-hide logic implemented
func testCreateItem(t *testing.T, repo ItemRepository, items *[]model.Item) {
	tests := []struct {
		name          string
		item          *model.Item
		expectedError error
		expectedId    uint64
	}{
		{
			name:          "valid 1",
			item:          &(*items)[0],
			expectedError: nil,
			expectedId:    1,
		},
		{
			name:          "valid 2",
			item:          &(*items)[1],
			expectedError: nil,
			expectedId:    2,
		},
		{
			name:          "valid 3",
			item:          &(*items)[2],
			expectedError: nil,
			expectedId:    3,
		},
		{
			name:          "title exists",
			item:          &(*items)[0],
			expectedError: &e.KeyAlreadyExist{Field: "title"},
		},
	}

	t.Run("CreateItem", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.CreateItem(test.item)

				if test.expectedError == nil {
					assert.Nil(t, err)
					assert.Equal(t, test.expectedId, test.item.Id)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

func testGetItems(t *testing.T, repo ItemRepository, items *[]model.Item, userId, userIdNotFound uint64) {
	tests := []struct {
		name          string
		userId        uint64
		expectedItems []model.Item
	}{
		{
			name:          "non-empty",
			userId:        userId,
			expectedItems: (*items)[:3],
		},
		{
			name:          "empty",
			userId:        userIdNotFound,
			expectedItems: nil,
		},
	}
	t.Run("GetItems", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actualGroups, err := repo.GetItems(test.userId)

				require.NoError(t, err)
				assert.Equal(t, test.expectedItems, actualGroups)
			})
		}
	})
}

func testUpdateItems(t *testing.T, repo ItemRepository, items *[]model.Item, userId, userIdNotFound uint64) {
	(*items)[0].Title = "d"

	tests := []struct {
		name          string
		itemToUpdate  *model.Item
		userId        uint64
		expectedError error
		expectedTitle string
	}{
		{
			name:          "valid",
			itemToUpdate:  &(*items)[0],
			userId:        userId,
			expectedError: nil,
			expectedTitle: "d",
		},
		{
			name: "not found",
			itemToUpdate: &model.Item{
				Id:     100,
				UserId: userId,
			},
			userId:        userId,
			expectedError: &e.KeyNotFound{Field: "itemId"},
		},
		{
			name: "title exists",
			itemToUpdate: &model.Item{
				Id:     1,
				Title:  "b",
				UserId: userId,
			},
			userId:        userId,
			expectedError: &e.KeyAlreadyExist{Field: "title"},
		},
		{
			name: "user has no access to item",
			itemToUpdate: &model.Item{
				Id:     1,
				UserId: userId,
			},
			userId:        userIdNotFound,
			expectedError: &e.NoAccess{Field: "userId,itemId"},
		},
	}
	t.Run("UpdateItem", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.UpdateItem(test.itemToUpdate, test.userId)

				if test.expectedError == nil {
					assert.Nil(t, err)
					assert.Equal(t, test.expectedTitle, test.itemToUpdate.Title)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

// TODO add more cases after un-hide logic implemented
func testHideItem(t *testing.T, repo ItemRepository, items *[]model.Item, userId, userIdNotFound uint64) {
	tests := []struct {
		name              string
		itemId            uint64
		userId            uint64
		expectedError     error
		expectedRestItems []model.Item
	}{
		{
			name:              "valid",
			itemId:            (*items)[0].Id,
			userId:            userId,
			expectedError:     nil,
			expectedRestItems: (*items)[1:3],
		},
		{
			name:          "not found",
			itemId:        (*items)[0].Id,
			userId:        userId,
			expectedError: &e.KeyNotFound{Field: "itemId"},
		},
		{
			name:          "user has no access to item",
			itemId:        (*items)[1].Id,
			userId:        userIdNotFound,
			expectedError: &e.NoAccess{Field: "userId,itemId"},
		},
	}

	t.Run("HideItem", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.HideItem(test.itemId, test.userId)

				if test.expectedError == nil {
					assert.Nil(t, err)

					actualRestItems, err := repo.GetItems(test.userId)

					assert.Nil(t, err)
					assert.Equal(t, test.expectedRestItems, actualRestItems)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

func testAddItemToGroup(t *testing.T, repo ItemRepository, items *[]model.Item,
	flatId uint64, groupIds []uint64, userId, userIdNotFound uint64) {
	tests := []struct {
		name          string
		itemId        uint64
		groupId       uint64
		flatId        uint64
		userId        uint64
		expectedError error
	}{
		{
			name:          "valid 1",
			itemId:        (*items)[1].Id,
			groupId:       groupIds[0],
			flatId:        flatId,
			userId:        userId,
			expectedError: nil,
		},
		{
			name:          "valid 2",
			itemId:        (*items)[2].Id,
			groupId:       groupIds[0],
			flatId:        flatId,
			userId:        userId,
			expectedError: nil,
		},
		{
			name:          "user has no access to flat",
			itemId:        (*items)[2].Id,
			userId:        userIdNotFound,
			expectedError: &e.NoAccess{Field: "userId,flatId"},
		},
		{
			name:          "not found",
			itemId:        (*items)[0].Id,
			userId:        userId,
			expectedError: &e.KeyNotFound{Field: "itemId"},
		},
		{
			name:          "item id already on flat group",
			itemId:        (*items)[2].Id,
			groupId:       groupIds[0],
			flatId:        flatId,
			userId:        userId,
			expectedError: &e.KeyAlreadyExist{Field: "flatId,groupId,itemId"},
		},
	}

	t.Run("AddItemToGroup", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.AddItemToGroup(test.flatId, test.groupId, test.itemId, test.userId)
				if test.expectedError == nil {
					assert.Nil(t, err)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

func testUpdateItemStatus(t *testing.T, repo ItemRepository, items *[]model.Item, flatId uint64, groupIds []uint64, userId uint64) {
	tests := []struct {
		name          string
		itemId        uint64
		groupId       uint64
		status        model.Status
		expectedError error
	}{
		{
			name:    "valid",
			itemId:  (*items)[1].Id,
			groupId: groupIds[0],
			status:  model.Ok,
		},
		{
			name:    "flat doesn't have group",
			itemId:  (*items)[1].Id,
			groupId: groupIds[2],
			//status: model.NotOk,
			expectedError: &e.NoAccess{Field: "flatId,groupId"},
		},
	}

	t.Run("UpdateItemStatus", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.UpdateItemStatus(flatId, test.groupId, test.itemId, test.status, userId)
				if test.expectedError == nil {
					assert.Nil(t, err)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

func testGetFlatItems(t *testing.T, repo ItemRepository, items *[]model.Item,
	flatId uint64, groupIds []uint64, userId uint64) {
	tests := []struct {
		name               string
		expectedGroupItems []model.GroupItems
	}{
		{
			name: "empty group and group with 2 items",
			expectedGroupItems: []model.GroupItems{
				{
					GroupId: groupIds[0],
					Items: []model.ItemWithStatus{
						{
							Item:   (*items)[1],
							Status: model.Ok,
						},
						{
							Item:   (*items)[2],
							Status: model.Unset,
						},
					},
				},
				{
					GroupId: groupIds[1],
					Items:   []model.ItemWithStatus{},
				},
			},
		},
	}

	t.Run("GetFlatItems", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actualGroupItems, err := repo.GetFlatItems(flatId, userId)

				require.NoError(t, err)

				// TODO consider more concise solution comparing "deep" objects
				test.expectedGroupItems = model.SortGroupItemsByGroupId(test.expectedGroupItems)
				actualGroupItems = model.SortGroupItemsByGroupId(actualGroupItems)

				expectedGroupIds := lo.Map(test.expectedGroupItems, func(groupItem model.GroupItems, _ int) uint64 {
					return groupItem.GroupId
				})

				actualGroupIds := lo.Map(actualGroupItems, func(groupItem model.GroupItems, _ int) uint64 {
					return groupItem.GroupId
				})

				// check group ids equality
				require.ElementsMatch(t, expectedGroupIds, actualGroupIds)

				// check each group items equality
				for i, _ := range test.expectedGroupItems {
					require.ElementsMatch(t, test.expectedGroupItems[i].Items, actualGroupItems[i].Items)
				}
			})
		}
	})
}

func testDeleteItemFromGroup(t *testing.T, repo ItemRepository, items *[]model.Item,
	flatId uint64, groupIds []uint64, userId, userIdNotFound uint64) {
	tests := []struct {
		name          string
		itemId        uint64
		groupId       uint64
		flatId        uint64
		userId        uint64
		expectedError error
	}{
		{
			name:          "valid 1",
			itemId:        (*items)[1].Id,
			groupId:       groupIds[0],
			flatId:        flatId,
			userId:        userId,
			expectedError: nil,
		},
		{
			name:          "user has no access to flat",
			itemId:        (*items)[2].Id,
			userId:        userIdNotFound,
			expectedError: &e.NoAccess{Field: "userId,flatId"},
		},
		{
			name:          "flat doesn't contain group",
			itemId:        (*items)[2].Id,
			groupId:       groupIds[2],
			flatId:        flatId,
			userId:        userId,
			expectedError: &e.NoAccess{Field: "flatId,groupId"},
		},
		{
			name:          "valid 2",
			itemId:        (*items)[2].Id,
			groupId:       groupIds[0],
			flatId:        flatId,
			userId:        userId,
			expectedError: nil,
		},
		{
			name:          "not found",
			itemId:        (*items)[0].Id,
			userId:        userId,
			expectedError: &e.KeyNotFound{Field: "itemId"},
		},
	}

	t.Run("DeleteItemFromGroup", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.DeleteItemFromGroup(test.flatId, test.groupId, test.itemId, test.userId)
				if test.expectedError == nil {
					assert.Nil(t, err)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}
