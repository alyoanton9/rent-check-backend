package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
	"testing"
)

func testMixedRepositories(t *testing.T, flatRepo FlatRepository, groupRepo GroupRepository,
	itemRepo ItemRepository, userId uint64) {

	testGetGroupFlats(t, flatRepo, groupRepo, userId)
}

func testGetGroupFlats(t *testing.T, flatRepo FlatRepository, groupRepo GroupRepository, userId uint64) {
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
	for i := range flats {
		err := flatRepo.CreateFlat(&(flats[i]))
		require.NoError(t, err)
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
	for i := range groups {
		err := groupRepo.CreateGroup(&(groups[i]))
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
			groupId:       groups[0].Id,
			flatId:        flats[0].Id,
			expectedFlats: []model.Flat{flats[0]},
		},
		{
			name:          "valid 2",
			groupId:       groups[0].Id,
			flatId:        flats[1].Id,
			expectedFlats: []model.Flat{flats[1], flats[0]},
		},
		{
			name:          "valid 3",
			groupId:       groups[1].Id,
			flatId:        flats[1].Id,
			expectedFlats: []model.Flat{flats[1]},
		},
		{
			name:          "group not found",
			groupId:       111,
			flatId:        flats[1].Id,
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
