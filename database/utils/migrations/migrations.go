package migrations

import (
	"github.com/gophermasters/bug-free-report/database/models"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(models.Bugs{})
}