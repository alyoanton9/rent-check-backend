package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"log"
	"os"
	"rent-checklist-backend/internal/config"
	"rent-checklist-backend/internal/database/postgres"
	"rent-checklist-backend/internal/handler"
	"rent-checklist-backend/internal/repository"
	"rent-checklist-backend/internal/route"
)

func App() {
	ctx := context.Background()

	appConfig := initConfig()

	var db *gorm.DB
	db, err := postgres.New(appConfig.Postgres, ctx)
	if err != nil {
		log.Fatal(err)
	}

	flat := repository.NewFlatRepository(db)
	item := repository.NewItemRepository(db)
	h := handler.NewHandler(flat, item)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())

	route.Setup(e, h)

	err = e.Start(fmt.Sprintf(":%s", appConfig.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() config.Config {
	var appConfig config.Config

	appConfig.Port = os.Getenv("HTTP_PORT")

	appConfig.Postgres = postgres.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	return appConfig
}
