package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func Connect() (*gorm.DB, error) {
	// Get database configuration with defaults
	host := getEnvOrDefault("DB_HOST", "localhost")
	port := getEnvOrDefault("DB_PORT", "5432")
	user := getEnvOrDefault("DB_USER", "postgres")
	password := getEnvOrDefault("DB_PASSWORD", "postgres")
	dbname := getEnvOrDefault("DB_NAME", "postgres")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Configure GORM with pgx driver
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Retry connection up to 5 times
	var err error
	for i := 0; i < 5; i++ {
		db, err := gorm.Open(postgres.New(postgres.Config{
			DriverName: "pgx",
			DSN:        dsn,
		}), gormConfig)
		if err == nil {
			// Configure connection pool
			sqlDB, err := db.DB()
			if err != nil {
				return nil, fmt.Errorf("failed to get database instance: %w", err)
			}
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(time.Hour)

			log.Printf("Database connection established to %s:%s/%s using pgx", host, port, dbname)
			DB = db
			return db, nil
		}
		log.Printf("Failed to connect to database (attempt %d/5): %v", i+1, err)
		time.Sleep(time.Second * 5)
	}

	return nil, fmt.Errorf("failed to connect to database after 5 attempts: %w", err)
}

func GetDB(ctx context.Context) *gorm.DB {
	return DB.WithContext(ctx)
}
