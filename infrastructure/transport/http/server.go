// neodata-go/http/server.go
package http

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/neodata-io/neodata-go/config"
)

// SetupHTTPServer initializes a new HTTP server with the provided configuration and middleware.
func SetupHTTPServer(cfg *config.AppConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.App.ReadTimeout * time.Second,
		WriteTimeout: cfg.App.WriteTimeout * time.Second,
		AppName:      cfg.App.Name,
	})

	// Middleware setup
	app.Use(LoggerMiddleware()) // Log all incoming requests
	// app.Use(RateLimiterMiddleware(100, time.Minute)) // Rate limiting for protected endpoints

	return app
}

// StartServer starts the HTTP server on the specified port.
func StartServer(app *fiber.App, cfg *config.AppConfig) (*fiber.App, error) {
	if err := app.Listen(fmt.Sprintf(":%d", cfg.App.Port), fiber.ListenConfig{
		DisableStartupMessage: true,
	}); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return nil, err
	}

	return app, nil
}
