package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
	"testing"
)

func testFlatRepository(t *testing.T, repo FlatRepository, ownerId string) {
	ownerIdNotFound := "111"

	flats := []model.Flat{
		{
			Address: "a",
			OwnerId: ownerId,
		},
		{
			Address: "b",
			OwnerId: ownerId,
		},
		{
			Address: "",
			OwnerId: ownerIdNotFound,
		},
	}

	testCreateFlat(t, repo, &flats)

	testGetFlats(t, repo, &flats, ownerId, ownerIdNotFound)

	testUpdateFlat(t, repo, ownerId, ownerIdNotFound)

	testDeleteFlat(t, repo, &flats, ownerId, ownerIdNotFound)
}

func testCreateFlat(t *testing.T, repo FlatRepository, flats *[]model.Flat) {
	tests := []struct {
		name          string
		flat          *model.Flat
		expectedError error
		expectedId    uint64
	}{
		{
			name:          "valid 1",
			flat:          &(*flats)[0],
			expectedError: nil,
			expectedId:    1,
		},
		{
			name:          "valid 2",
			flat:          &(*flats)[1],
			expectedError: nil,
			expectedId:    2,
		},
		{
			name:          "address exists",
			flat:          &(*flats)[0],
			expectedError: &e.KeyAlreadyExist{Field: "address"},
		},
		{
			name:          "ownerId not found",
			flat:          &(*flats)[2],
			expectedError: &e.KeyNotFound{Field: "ownerId"},
		},
	}

	t.Run("CreateFlat", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.CreateFlat(test.flat)

				if test.expectedError == nil {
					assert.Nil(t, err)
					assert.Equal(t, test.expectedId, test.flat.Id)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

func testGetFlats(t *testing.T, repo FlatRepository, flats *[]model.Flat, ownerId, ownerIdNotFound string) {
	tests := []struct {
		name          string
		userId        string
		expectedFlats []model.Flat
	}{
		{
			name:          "non-empty",
			userId:        ownerId,
			expectedFlats: (*flats)[:2],
		},
		{
			name:          "empty",
			userId:        ownerIdNotFound,
			expectedFlats: nil,
		},
	}
	t.Run("GetFlats", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actualFlats, err := repo.GetFlats(test.userId)

				require.NoError(t, err)
				assert.Equal(t, test.expectedFlats, actualFlats)

			})
		}
	})
}

func testUpdateFlat(t *testing.T, repo FlatRepository, ownerId, ownerIdNotFound string) {
	tests := []struct {
		name            string
		flatToUpdate    *model.Flat
		userId          string
		expectedError   error
		expectedAddress string
	}{
		{
			name: "valid",
			flatToUpdate: &model.Flat{
				Id:      1,
				Address: "c",
				OwnerId: ownerId,
			},
			userId:          ownerId,
			expectedError:   nil,
			expectedAddress: "c",
		},
		{
			name: "not found",
			flatToUpdate: &model.Flat{
				Id:      100,
				OwnerId: ownerId,
			},
			userId:        ownerId,
			expectedError: &e.KeyNotFound{Field: "id"},
		},
		{
			name: "address exists",
			flatToUpdate: &model.Flat{
				Id:      1,
				Address: "b",
				OwnerId: ownerId,
			},
			userId:        ownerId,
			expectedError: &e.KeyAlreadyExist{Field: "address"},
		},
		{
			name: "user has no access to flat",
			flatToUpdate: &model.Flat{
				Id:      1,
				OwnerId: ownerId,
			},
			userId:        ownerIdNotFound,
			expectedError: &e.NoAccess{Field: "userId,flatId"},
		},
	}
	t.Run("UpdateFlat", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.UpdateFlat(test.flatToUpdate, test.userId)

				if test.expectedError == nil {
					assert.Nil(t, err)
					assert.Equal(t, test.expectedAddress, test.flatToUpdate.Address)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}

func testDeleteFlat(t *testing.T, repo FlatRepository, flats *[]model.Flat, ownerId, ownerIdNotFound string) {
	tests := []struct {
		name          string
		flatId        uint64
		userId        string
		expectedError error
	}{
		{
			name:          "valid",
			flatId:        (*flats)[0].Id,
			userId:        ownerId,
			expectedError: nil,
		},
		{
			name:          "not found",
			flatId:        (*flats)[0].Id,
			userId:        ownerId,
			expectedError: &e.KeyNotFound{Field: "id"},
		},
		{
			name:          "user has no access to flat",
			flatId:        (*flats)[1].Id,
			userId:        ownerIdNotFound,
			expectedError: &e.NoAccess{Field: "userId,flatId"},
		},
	}

	t.Run("DeleteFlat", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := repo.DeleteFlat(test.flatId, test.userId)

				if test.expectedError == nil {
					assert.Nil(t, err)
				} else {
					assert.EqualError(t, err, test.expectedError.Error())
				}
			})
		}
	})
}
