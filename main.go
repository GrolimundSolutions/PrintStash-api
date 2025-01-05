package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GrolimundSolutions/PrintStash-api/database"
	"github.com/GrolimundSolutions/PrintStash-api/handlers"
	"github.com/GrolimundSolutions/PrintStash-api/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run database migrations
	db.AutoMigrate(&models.Color{}, &models.Manufacturer{}, &models.Material{}, &models.FilamentSpool{}, &models.PrintSetting{})

	// Setup Gin router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	v1 := router.Group("/api/v1")
	{
		// Manufacturers routes
		v1.GET("/manufacturers", handlers.GetAllManufacturers)
		v1.GET("/manufacturers/:id", handlers.GetManufacturer)
		v1.POST("/manufacturers", handlers.CreateManufacturer)
		v1.PUT("/manufacturers/:id", handlers.UpdateManufacturer)
		v1.DELETE("/manufacturers/:id", handlers.DeleteManufacturer)

		// Materials routes
		v1.GET("/materials", handlers.GetAllMaterials)
		v1.GET("/materials/:id", handlers.GetMaterial)
		v1.POST("/materials", handlers.CreateMaterial)
		v1.PUT("/materials/:id", handlers.UpdateMaterial)
		v1.DELETE("/materials/:id", handlers.DeleteMaterial)

		// Colors routes
		v1.GET("/colors", handlers.GetAllColors)
		v1.GET("/colors/:id", handlers.GetColor)
		v1.POST("/colors", handlers.CreateColor)
		v1.PUT("/colors/:id", handlers.UpdateColor)
		v1.DELETE("/colors/:id", handlers.DeleteColor)

		// Filament spools routes
		v1.GET("/filament-spools", handlers.GetAllFilamentSpools)
		v1.GET("/filament-spools/:id", handlers.GetFilamentSpool)
		v1.POST("/filament-spools", handlers.CreateFilamentSpool)
		v1.PUT("/filament-spools/:id", handlers.UpdateFilamentSpool)
		v1.DELETE("/filament-spools/:id", handlers.DeleteFilamentSpool)

		// Print settings routes
		v1.GET("/print-settings", handlers.GetAllPrintSettings)
		v1.GET("/print-settings/:id", handlers.GetPrintSetting)
		v1.POST("/print-settings", handlers.CreatePrintSetting)
		v1.PUT("/print-settings/:id", handlers.UpdatePrintSetting)
		v1.DELETE("/print-settings/:id", handlers.DeletePrintSetting)
	}

	// Configure server
	srv := &http.Server{
		Addr:    ":" + os.Getenv("API_PORT"),
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Give outstanding requests 5 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
