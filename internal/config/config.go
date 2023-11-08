package config

import "rent-checklist/internal/database/postgres"

type Config struct {
	Port     string
	Postgres postgres.Config
}
