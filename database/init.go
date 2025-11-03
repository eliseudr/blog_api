package db

import (
	"fmt"
	"os"

	"github.com/eliseudr/blog_api/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// LoadConfig loads database configuration from environment variables
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

// Validate the configuration
func validateConfig(config *Config) error {
	envVars := map[string]string{
		"DB_HOST":     config.Host,
		"DB_PORT":     config.Port,
		"DB_USER":     config.User,
		"DB_PASSWORD": config.Password,
		"DB_NAME":     config.Name,
	}

	for name, value := range envVars {
		if value == "" {
			return fmt.Errorf("environment variable %s is required", name)
		}
	}
	return nil
}

// Initialize creates the database if it doesn't exist and runs migrations
func Initialize(config *Config) (*gorm.DB, error) {
	// 1. Connect to the MySQL server (In case the database doesn't exist)
	serverDB, err := OpenServer(config.User, config.Password, config.Host, config.Port)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL server: %w", err)
	}
	fmt.Println("Connected to MySQL server")

	// 2. Get the database instance
	sqlDB, err := serverDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// 3. Create the database if it doesn't exist
	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS " + config.Name)
	if err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("failed to create database: %w", err)
	}
	fmt.Println("Database created/verified:", config.Name)

	sqlDB.Close()

	// 4. Connect to the database
	database, err := Open(config.User, config.Password, config.Host, config.Port, config.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("Connected to database:", config.Name)

	// 5. Auto migrate the models (Create the tables in the database)
	err = database.AutoMigrate(&models.BlogPost{}, &models.Comment{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	fmt.Println("Database migration completed")

	return database, nil
}
