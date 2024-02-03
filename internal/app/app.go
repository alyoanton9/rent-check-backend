package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"rent-checklist-backend/internal/config"
	"rent-checklist-backend/internal/database/postgres"
	"rent-checklist-backend/internal/encrypt"
	"rent-checklist-backend/internal/handler"
	"rent-checklist-backend/internal/middleware"
	"rent-checklist-backend/internal/repository"
	"rent-checklist-backend/internal/route"
	"rent-checklist-backend/internal/service"
)

func App() {
	ctx := context.Background()

	appConfig := initConfig()

	const migrationPath string = "file://internal/database/postgres/migrations"

	var db *gorm.DB
	db, err := postgres.New(appConfig.Postgres, migrationPath, ctx)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(db)
	flatRepository := repository.NewFlatRepository(db)
	groupRepository := repository.NewGroupRepository(db)
	itemRepository := repository.NewItemRepository(db)
	hasher := encrypt.NewHasher()
	authService := service.NewAuthService("secret string!")

	h := handler.NewHandler(userRepository, flatRepository, groupRepository, itemRepository, hasher, authService)

	e := echo.New()

	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.BodyDump(middleware.LogRequestAndResponseBody))

	e.Pre(echoMiddleware.RemoveTrailingSlash())
	e.Use(echoMiddleware.KeyAuthWithConfig(middleware.MakeAuthConfig(authService, userRepository)))

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType},
	}))

	e.Use(echoMiddleware.Recover())

	route.Setup(e, h)

	domain := os.Getenv("DOMAIN")
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("/var/www/.cache"),
		HostPolicy: autocert.HostWhitelist(domain),
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", appConfig.Port),
		Handler: e,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	// http server is needed to pass Let's Encrypt HTTP-01 challenge
	httpPort := os.Getenv("HTTP_PORT")
	go http.ListenAndServe(fmt.Sprintf(":%s", httpPort), certManager.HTTPHandler(nil))

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

func initConfig() config.Config {
	var appConfig config.Config

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to read .env file: %s", err.Error())
	}

	appConfig.Port = os.Getenv("HTTPS_PORT")

	appConfig.Postgres = postgres.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	return appConfig
}
