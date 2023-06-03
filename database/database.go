package database

import (
	"fmt"
	"os"
	"time"

	"github.com/gophermasters/bug-free-report/database/utils/migrations"
	"github.com/gophermasters/bug-free-report/database/utils/errors"
	"github.com/theGOURL/warning"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

// Connect establishes the connection to the database.
func Connect() {
	connectionURI := os.Getenv("DATABASE_URI")
	db, err := gorm.Open(postgres.Open(connectionURI), &gorm.Config{})
	errors.FatalError(err)

	fmt.Println("Database Connected")

	database = db

	config, err := database.DB()
	errors.FatalError(err)

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.Run(database)

	fmt.Println("Migrations Finished")
}

// CloseConn closes the connection to the database.
func CloseConn() {
	config, err := database.DB()
	errors.CloseConnError(err)

	err = config.Close()
	errors.CloseConnError(err)
}

// GetDatabase returns the database instance.
func GetDatabase() *gorm.DB {
	return database
}
