package models

import (
	"time"

	"gorm.io/gorm"
)

type Bugs struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	BugName     string         `json:"bugname"`
	Description string         `json:"description"`
	ErrorCode   float32        `json:"error_code"`
	Username    string         `json:"username"`
	ImageURL    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}