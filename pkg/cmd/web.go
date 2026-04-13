package cmd

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/tecknohow/TaskDream/pkg/config"
	"github.com/tecknohow/TaskDream/pkg/db"
	"github.com/tecknohow/TaskDream/pkg/migration"
	"github.com/tecknohow/TaskDream/pkg/routes"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the TaskDream web server",
	Long:  `Start the TaskDream HTTP server with configured routes and middleware.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		// Initialize database
		engine, err := db.InitDB(cfg)
		if err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}
		defer engine.Close()

		// Run migrations
		if err := migration.RunMigrations(engine); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}

		// Store engine globally for route handlers
		db.Engine = engine

		// Create Echo instance
		e := echo.New()

		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
			AllowHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		}))

		// Setup routes
		routes.SetupRoutes(e)

		// Start server
		addr := fmt.Sprintf(":%d", cfg.Server.Port)
		log.Printf("Starting TaskDream server on %s", addr)
		if err := e.Start(addr); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	},
}
