package repository

import (
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
	"testing"
)

func testMixedRepositories(t *testing.T, flatRepo FlatRepository, groupRepo GroupRepository,
	itemRepo ItemRepository, userId uint64) {
	flats := []model.Flat{
		{
			Address: "d",
			OwnerId: userId,
		},
		{
			Address: "e",
			OwnerId: userId,
		},
	}

	groups := []model.Group{
		{
			Title:  "e",
			UserId: userId,
		},
		{
			Title:  "f",
			UserId: userId,
		},
	}

	testGetGroupFlats(t, flatRepo, groupRepo, &flats, &groups, userId)

	testCopyItemsFromFlatGroup(t, itemRepo, &flats, &groups, userId)
}

func testGetGroupFlats(t *testing.T, flatRepo FlatRepository, groupRepo GroupRepository,
	flats *[]model.Flat, groups *[]model.Group, userId uint64) {
	for i := range *flats {
		err := flatRepo.CreateFlat(&(*flats)[i])
		require.NoError(t, err)
	}

	for i := range *groups {
		err := groupRepo.CreateGroup(&(*groups)[i])
		require.NoError(t, err)
	}

	tests := []struct {
		name          string
		groupId       uint64
		flatId        uint64
		expectedFlats []model.Flat
		expectedError error
	}{
		{
			name:          "valid 1",
			groupId:       (*groups)[0].Id,
			flatId:        (*flats)[0].Id,
			expectedFlats: []model.Flat{(*flats)[0]},
		},
		{
			name:          "valid 2",
			groupId:       (*groups)[0].Id,
			flatId:        (*flats)[1].Id,
			expectedFlats: []model.Flat{(*flats)[1], (*flats)[0]},
		},
		{
			name:          "valid 3",
			groupId:       (*groups)[1].Id,
			flatId:        (*flats)[1].Id,
			expectedFlats: []model.Flat{(*flats)[1]},
		},
		{
			name:          "group not found",
			groupId:       111,
			flatId:        (*flats)[1].Id,
			expectedError: &e.KeyNotFound{Field: "groupId"},
		},
	}

	t.Run("GetGroupFlats", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := groupRepo.AddGroupToFlat(test.groupId, test.flatId, userId)
				if test.expectedError == nil {
					require.NoError(t, err)

					actualFlats, err := flatRepo.GetFlatsByGroupId(test.groupId, userId)
					require.NoError(t, err)
					assert.ElementsMatch(t, test.expectedFlats, actualFlats)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

func testCopyItemsFromFlatGroup(t *testing.T, repo ItemRepository,
	flats *[]model.Flat, groups *[]model.Group, userId uint64) {
	items := []model.Item{
		{
			Title:  "e",
			UserId: userId,
		},
		{
			Title:  "f",
			UserId: userId,
		},
	}

	tests := []struct {
		name          string
		groupId       uint64
		flatId        uint64
		flatIdCopy    uint64
		items         []model.Item
		expectedError error
	}{
		{
			name:       "valid",
			groupId:    (*groups)[0].Id,
			flatId:     (*flats)[1].Id,
			flatIdCopy: (*flats)[0].Id,
			items:      items,
		},
		{
			name:          "user has no access to flat flatId",
			groupId:       (*groups)[1].Id,
			flatId:        111,
			flatIdCopy:    1,
			expectedError: &e.NoAccess{Field: "userId,flatId"},
		},
		{
			name:          "user has no access to flat flatIdCopy",
			groupId:       (*groups)[1].Id,
			flatId:        1,
			flatIdCopy:    111,
			expectedError: &e.NoAccess{Field: "userId,flatId"},
		},
		{
			name:          "group not found",
			groupId:       111,
			flatId:        (*flats)[0].Id,
			flatIdCopy:    (*flats)[1].Id,
			expectedError: &e.KeyNotFound{Field: "groupId"},
		},
	}

	t.Run("CopyItemsFromFlatGroup", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				if test.items != nil {
					for i := range test.items {
						err := repo.CreateItem(&test.items[i])
						require.NoError(t, err)

						err = repo.AddItemToGroup(test.flatIdCopy, test.groupId, test.items[i].Id, userId)
						require.NoError(t, err)
					}
				}

				err := repo.CopyItemsFromFlatGroup(test.groupId, test.flatId, test.flatIdCopy, userId)

				if test.expectedError == nil {
					assert.Nil(t, err)

					groupItemsExpected, err := repo.GetFlatItems(test.flatIdCopy, userId)
					require.NoError(t, err)

					groupItemsActual, err := repo.GetFlatItems(test.flatId, userId)
					require.NoError(t, err)

					groupItemsExpected = lo.Filter(groupItemsExpected, func(groupItem model.GroupItems, _ int) bool {
						return groupItem.GroupId == test.groupId
					})

					groupItemsActual = lo.Filter(groupItemsActual, func(groupItem model.GroupItems, _ int) bool {
						return groupItem.GroupId == test.groupId
					})

					assert.ElementsMatch(t, groupItemsExpected, groupItemsActual)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}
