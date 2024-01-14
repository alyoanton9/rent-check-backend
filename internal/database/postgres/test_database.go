package postgres

import (
	"context"
	testContainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

func NewTestDB(migrationPath string) (testContainers.Container, *gorm.DB, error) {
	// TODO consider moving to some test '.env' file
	cfg := Config{
		User:     "postgres",
		Password: "password",
		Database: "db",
	}

	request := testContainers.ContainerRequest{
		Image:        "postgres:15.2-alpine",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_USER":     cfg.User,
			"POSTGRES_PASSWORD": cfg.Password,
			"POSTGRES_DB":       cfg.Database,
		},
	}

	container, err := testContainers.GenericContainer(
		context.Background(),
		testContainers.GenericContainerRequest{
			ContainerRequest: request,
			Started:          true,
		})
	if err != nil {
		return nil, nil, err
	}

	host, err := container.Host(context.Background())
	if err != nil {
		return nil, nil, err
	}

	port, err := container.MappedPort(context.Background(), "5432")
	if err != nil {
		return nil, nil, err
	}

	cfg.Host = host
	cfg.Port = port.Port()

	db, err := New(cfg, migrationPath, context.Background())
	if err != nil {
		return nil, nil, err
	}

	return container, db, nil
}
