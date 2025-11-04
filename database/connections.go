package db

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Open a connection to the MySQL database (localhost)
func Open(user, pass, host, port, name string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	// Get server mode from environment variable (defaults to PRODUCTION if not set)
	serverMode := strings.ToUpper(os.Getenv("SERVER_MODE"))
	if serverMode == "" {
		serverMode = "PRODUCTION"
	}

	// Configure GORM config with conditional logging
	var gormConfig *gorm.Config

	if serverMode == "DEV" {
		// Configure logger to show SQL queries in DEV mode
		gormLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             0,           // Log all queries
				LogLevel:                  logger.Info, // Log level: Silent, Error, Warn, Info
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Enable color
			},
		)
		gormConfig = &gorm.Config{
			Logger: gormLogger,
		}
	} else {
		// Production mode: disable SQL logging
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	return gorm.Open(mysql.Open(dsn), gormConfig)
}

// Open a connection to the MySQL server (localhost)
func OpenServer(user, pass, host, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
