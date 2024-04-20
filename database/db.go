package database

import (
	"gin-task-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func InitDB() (*Database, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate schema (assuming tables don't exist)
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}
