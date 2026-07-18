package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"destiny-backend/config"
	_ "destiny-backend/docs"
	"destiny-backend/internal/routes"
	"destiny-backend/pkg/database"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to PostgreSQL
	pool, err := database.NewPostgresPool(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer pool.Close()
	log.Println("Connected to PostgreSQL successfully")

	// Run migrations
	if err := runMigrations(cfg); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")

	// Connect to Redis (optional, non-fatal if unavailable)
	var redisClient *redis.Client
	redisClient, err = database.NewRedisClient(cfg.Redis)
	if err != nil {
		log.Printf("Warning: Redis connection failed: %v (continuing without Redis)", err)
		redisClient = nil
	} else {
		defer redisClient.Close()
		log.Println("Connected to Redis successfully")
	}

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status":  "ok",
			"service": "destiny-backend",
		})
	})

	// Swagger docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Setup routes
	routes.Setup(e, pool, redisClient, cfg.JWT.Secret)

	// Start server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", addr)

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("Shutting down server...")
		e.Close()
	}()

	if err := e.Start(addr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
	log.Println("Server stopped")
}

func runMigrations(cfg *config.Config) error {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.Host, cfg.Postgres.Port,
		cfg.Postgres.DBName, cfg.Postgres.SSLMode,
	)

	migrationsPath := "migrations"
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		altPaths := []string{
			"backend/migrations",
			"../migrations",
			"../../migrations",
		}
		found := false
		for _, p := range altPaths {
			if _, err := os.Stat(p); err == nil {
				migrationsPath = p
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("migrations directory not found in any expected location")
		}
	}

	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to resolve migrations path: %w", err)
	}

	m, err := migrate.New("file://"+absPath, dbURL)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration up failed: %w", err)
	}

	return nil
}
