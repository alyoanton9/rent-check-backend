package postgres

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func New(config Config, migrationPath string, ctx context.Context) (*gorm.DB, error) {
	connString := makeConnectionString(config)

	m, err := migrate.New(migrationPath, connString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database to migrate: %s", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	version, dirty, err := m.Version()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("applied migration: %d, dirty: %t\n", version, dirty)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{TranslateError: true})

	return db, nil
}

func makeConnectionString(cfg Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}
