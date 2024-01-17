package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
	"testing"
)

func testGroupRepository(t *testing.T, repo GroupRepository, userId uint64) {
	var userIdNotFound uint64 = 111
	var flatId uint64 = 2

	groups := []model.Group{
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

	testCreateGroup(t, repo, &groups)

	testGetGroups(t, repo, &groups, userId, userIdNotFound)

	testUpdateGroup(t, repo, &groups, userId, userIdNotFound)

	testHideGroup(t, repo, &groups, userId, userIdNotFound)

	testAddGroupToFlat(t, repo, &groups, flatId, userId, userIdNotFound)

	testDeleteGroupFromFlat(t, repo, &groups, flatId, userId)

	// add groups back to use them in item repo tests
	testAddGroupToFlat(t, repo, &groups, flatId, userId, userIdNotFound)
}

func testCreateGroup(t *testing.T, repo GroupRepository, groups *[]model.Group) {
	tests := []struct {
		name          string
		group         *model.Group
		expectedError error
		expectedId    uint64
	}{
		{
			name:          "valid 1",
			group:         &(*groups)[0],
			expectedError: nil,
			expectedId:    1,
		},
		{
			name:          "valid 2",
			group:         &(*groups)[1],
			expectedError: nil,
			expectedId:    2,
		},
		{
			name:          "valid 3",
			group:         &(*groups)[2],
			expectedError: nil,
			expectedId:    3,
		},
		{
			name:          "title exists",
			group:         &(*groups)[0],
			expectedError: &e.KeyAlreadyExist{Field: "title"},
		},
		{
			name:          "userId not found",
			group:         &(*groups)[3],
			expectedError: &e.KeyNotFound{Field: "userId"},
		},
	}

	t.Run("CreateGroup", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.CreateGroup(test.group)

				if test.expectedError == nil {
					assert.Nil(t, err)
					assert.Equal(t, test.expectedId, test.group.Id)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

func testGetGroups(t *testing.T, repo GroupRepository, groups *[]model.Group, userId, userIdNotFound uint64) {
	tests := []struct {
		name           string
		userId         uint64
		groupIds       []uint64
		expectedGroups []model.Group
	}{
		{
			name:           "non-empty all groups",
			userId:         userId,
			expectedGroups: (*groups)[:3],
		},
		{
			name:           "non-empty certain groups",
			userId:         userId,
			groupIds:       []uint64{1, 2},
			expectedGroups: (*groups)[:2],
		},
		{
			name:           "empty",
			userId:         userIdNotFound,
			expectedGroups: nil,
		},
	}
	t.Run("GetGroups", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actualGroups, err := repo.GetGroups(test.userId, test.groupIds)

				require.NoError(t, err)
				assert.Equal(t, test.expectedGroups, actualGroups)
			})
		}
	})
}

func testUpdateGroup(t *testing.T, repo GroupRepository, groups *[]model.Group, userId, userIdNotFound uint64) {
	(*groups)[0].Title = "d"

	tests := []struct {
		name          string
		groupToUpdate *model.Group
		userId        uint64
		expectedError error
		expectedTitle string
	}{
		{
			name:          "valid",
			groupToUpdate: &(*groups)[0],
			userId:        userId,
			expectedError: nil,
			expectedTitle: "d",
		},
		{
			name: "not found",
			groupToUpdate: &model.Group{
				Id:     100,
				UserId: userId,
			},
			userId:        userId,
			expectedError: &e.KeyNotFound{Field: "groupId"},
		},
		{
			name: "title exists",
			groupToUpdate: &model.Group{
				Id:     1,
				Title:  "b",
				UserId: userId,
			},
			userId:        userId,
			expectedError: &e.KeyAlreadyExist{Field: "title"},
		},
		{
			name: "user has no access to group",
			groupToUpdate: &model.Group{
				Id:     1,
				UserId: userId,
			},
			userId:        userIdNotFound,
			expectedError: &e.NoAccess{Field: "userId,groupId"},
		},
	}
	t.Run("UpdateGroup", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.UpdateGroup(test.groupToUpdate, test.userId)

				if test.expectedError == nil {
					assert.Nil(t, err)
					assert.Equal(t, test.expectedTitle, test.groupToUpdate.Title)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

// TODO add more cases after un-hide logic implemented
func testHideGroup(t *testing.T, repo GroupRepository, groups *[]model.Group, userId, userIdNotFound uint64) {
	tests := []struct {
		name               string
		groupId            uint64
		userId             uint64
		expectedError      error
		expectedRestGroups []model.Group
	}{
		{
			name:               "valid",
			groupId:            (*groups)[0].Id,
			userId:             userId,
			expectedError:      nil,
			expectedRestGroups: (*groups)[1:3],
		},
		{
			name:          "not found",
			groupId:       (*groups)[0].Id,
			userId:        userId,
			expectedError: &e.KeyNotFound{Field: "groupId"},
		},
		{
			name:          "user has no access to group",
			groupId:       (*groups)[1].Id,
			userId:        userIdNotFound,
			expectedError: &e.NoAccess{Field: "userId,groupId"},
		},
	}

	t.Run("HideGroup", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.HideGroup(test.groupId, test.userId)

				if test.expectedError == nil {
					assert.Nil(t, err)

					var ids []uint64
					actualRestGroups, err := repo.GetGroups(test.userId, ids)

					assert.Nil(t, err)
					assert.Equal(t, test.expectedRestGroups, actualRestGroups)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

func testAddGroupToFlat(t *testing.T, repo GroupRepository, groups *[]model.Group,
	flatId uint64, userId, userIdNotFound uint64) {
	tests := []struct {
		name          string
		groupId       uint64
		userId        uint64
		expectedError error
	}{
		{
			name:          "valid 1",
			groupId:       (*groups)[1].Id,
			userId:        userId,
			expectedError: nil,
		},
		{
			name:          "valid 2",
			groupId:       (*groups)[2].Id,
			userId:        userId,
			expectedError: nil,
		},
		{
			name:          "user has no access to flat",
			groupId:       (*groups)[2].Id,
			userId:        userIdNotFound,
			expectedError: &e.NoAccess{Field: "userId,flatId"},
		},
		{
			name:          "not found",
			groupId:       (*groups)[0].Id,
			userId:        userId,
			expectedError: &e.KeyNotFound{Field: "groupId"},
		},
		{
			name:          "already added",
			groupId:       (*groups)[1].Id,
			userId:        userId,
			expectedError: &e.KeyAlreadyExist{Field: "groupId"},
		},
	}

	t.Run("AddGroupToFlat", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.AddGroupToFlat(test.groupId, flatId, test.userId)
				if test.expectedError == nil {
					assert.Nil(t, err)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

func testDeleteGroupFromFlat(t *testing.T, repo GroupRepository, groups *[]model.Group, flatId uint64, userId uint64) {
	tests := []struct {
		name          string
		groupId       uint64
		userId        uint64
		expectedError error
	}{
		{
			name:          "valid 1",
			groupId:       (*groups)[1].Id,
			userId:        userId,
			expectedError: nil,
		},
		{
			name:          "valid 2",
			groupId:       (*groups)[2].Id,
			userId:        userId,
			expectedError: nil,
		},
		{
			name:          "not found",
			groupId:       (*groups)[0].Id,
			userId:        userId,
			expectedError: &e.KeyNotFound{Field: "groupId"},
		},
	}

	t.Run("DeleteGroupFromFlat", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.DeleteGroupFromFlat(test.groupId, flatId, test.userId)
				if test.expectedError == nil {
					assert.Nil(t, err)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}
