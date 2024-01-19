package repository

import (
	"context"
	"github.com/stretchr/testify/require"
	"rent-checklist-backend/internal/database/postgres"
	"rent-checklist-backend/internal/model"
	"testing"
)

func TestRepository(t *testing.T) {
	const migrationPath string = "file://../database/postgres/migrations"
	container, db, err := postgres.NewTestDB(migrationPath)
	require.NoError(t, err)

	defer container.Terminate(context.Background())

	userRepo := NewUserRepository(db)

	user := model.User{Login: "login", PasswordHash: "hash"}
	err = userRepo.CreateUser(&user)

	require.NoError(t, err)

	flatRepo := NewFlatRepository(db)

	t.Run("FlatRepository", func(t *testing.T) {
		testFlatRepository(t, flatRepo, user.Id)
	})

	groupRepo := NewGroupRepository(db)

	t.Run("GroupRepository", func(t *testing.T) {
		testGroupRepository(t, groupRepo, user.Id)
	})

	itemRepo := NewItemRepository(db)

	t.Run("ItemRepository", func(t *testing.T) {
		testItemRepository(t, itemRepo, user.Id)
	})

	t.Run("Mixed", func(t *testing.T) {
		testMixedRepositories(t, flatRepo, groupRepo, itemRepo, user.Id)
	})
}
