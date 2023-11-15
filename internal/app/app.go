package app

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"rent-checklist/internal/config"
	"rent-checklist/internal/database/postgres"
	"rent-checklist/internal/handler"
	"rent-checklist/internal/repository"
	"rent-checklist/internal/route"
)

func App() {
	ctx := context.Background()

	appConfig := initConfig()

	var db *pg.DB
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