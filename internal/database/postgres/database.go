package postgres

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

const migrationPath string = "file://internal/database/postgres/migrations"

func New(config Config, ctx context.Context) (*pg.DB, error) {
	connString := makeConnectionString(config)

	log.Printf("connection string: %s", connString)

	m, err := migrate.New(migrationPath, connString)
	if err != nil {
		return nil, fmt.Errorf("error migrating database: %s", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	dbClient := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
	})

	return dbClient, nil
}

func makeConnectionString(cfg Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}
