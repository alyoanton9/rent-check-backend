package config

import "rent-checklist-backend/internal/database/postgres"

type Config struct {
	Port     string
	Postgres postgres.Config
}
