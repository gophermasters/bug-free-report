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

var database *gorm.DB;


func Connect(){
	
	conectDB := os.Getenv("DATABASE_URI");
	db, err := gorm.Open(postgres.Open(conectDB), &gorm.Config{});
		warning.FATAL_ERROR(err);
		fmt.Println("Database Connected");

	database = db;
		config, _ := database.DB();
			config.SetMaxIdleConns(10);
			config.SetMaxOpenConns(100);
			config.SetConnMaxLifetime(time.Hour);
		
		migrations.Run(database);
			fmt.Println("Migrations Finished");
	
}

func CloseConn(){
	config, err := database.DB();
		errors.CloseConnError(err);
	err = config.Close();
		errors.CloseConnError(err);
}

func GetDatabase() *gorm.DB {
	return database
}