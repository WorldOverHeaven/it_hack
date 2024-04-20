package main

import (
	"context"
	"errors"
	echoSwagger "github.com/swaggo/echo-swagger"
	"mephi_hack/backend/internal/config"
	"mephi_hack/backend/internal/database"
	"mephi_hack/backend/internal/handler"
	"mephi_hack/backend/internal/service"
	"mephi_hack/pkg/auth"

	"flag"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"
	_ "mephi_hack/backend/docs"
)

const secret = "niggers"

func parseRootPath() string {
	var rootPath string
	flag.StringVar(&rootPath, "rootPath", "", "root folder")
	flag.Parse()
	return rootPath
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:8080
// @securityDefinitions.basic  BasicAuth
func main() {
	rootPath := parseRootPath()
	cfg, err := setupCfg(rootPath)
	fmt.Println("CFG", cfg)
	if err != nil {
		log.Fatalf("failed to parse config: %e\n", err)
	}

	db, err := setupDb(cfg)
	if err != nil {
		log.Fatalf("db error %e\n", err)
	}

	a := auth.New(secret)

	s := service.New(db, a)

	h := handler.New(s)

	e := setupEcho()

	handler.SetupRoutes(e, h)

	go func() {
		if err := e.Start(":" + cfg.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func setupDb(cfg config.Config) (database.Database, error) {
	// create a database connection
	db, err := database.NewDatabase(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("failed to create database conn %e\n", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, errors.New("failed to connect to database")
	}

	return db, nil
}

func setupEcho() *echo.Echo {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(Logger)

	return e
}

func setupCfg(rootPath string) (config.Config, error) {
	ctx := context.Background()

	var cfg config.Config

	err := confita.NewLoader(
		file.NewBackend(fmt.Sprintf("%s/deploy/back/default.yaml", rootPath)),
		env.NewBackend(),
	).Load(ctx, &cfg)

	if err != nil {
		return config.Config{}, err
	}

	envDocker := os.Getenv("ENV")
	if envDocker != "docker" {
		cfg.Postgres.Host = "localhost"
	}

	return cfg, nil
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		if err := next(c); err != nil {
			c.Error(err)
		}

		log.Printf(
			"%s %s",
			c.Path(),
			time.Since(start),
		)

		return nil
	}
}

func CreateJwtMiddlewareWithService(a auth.Service) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Error(err)
			}

			return nil
		}
	}
}
