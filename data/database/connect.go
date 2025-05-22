package database

import (
	"fmt"
	"strconv"

	"github.com/gabereiser/blog/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	// This function is used to connect to a production database
	// It uses PostgreSQL as the database driver
	// and retrieves the connection details from the config env

	var err error
	p := config.Get("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}
	// formatting the connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Get("DB_HOST"), port, config.Get("DB_USER"), config.Get("DB_PASSWORD"), config.Get("DB_NAME"))

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	Migrate()
	fmt.Println("Database Migrated")
}

func ConnectTest() {
	// This function is used to connect to a test database
	var err error
	// Use an in-memory SQLite database for testing
	dsn := "file::memory:?cache=shared"
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	Migrate()
	fmt.Println("Database Migrated")
}
