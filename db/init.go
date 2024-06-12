package db

import (
	"github.com/handsomexdd1024/sp24-tsdt-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB creates a new database connection and returns it.
func InitDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.TodoList{}, &models.TodoItem{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
