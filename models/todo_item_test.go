package models_test

import (
	"os"
	"testing"
	"time"

	. "github.com/handsomexdd1024/sp24-tsdt-go/db"
	. "github.com/handsomexdd1024/sp24-tsdt-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func createTempDatabase() (*gorm.DB, string) {
	// create a temporary database with current time
	// to avoid conflict with other tests
	dbName := time.Now().Format("test20060102150405.db")
	db, err := InitDB(dbName)
	if err != nil {
		panic("failed to connect database")
	}
	return db, dbName
}

func removeDatabase(dbName string) {
	// remove the temporary database
	err := os.Remove(dbName)
	if err != nil {
		panic("failed to remove database")
	}

}

func TestNewTodoList(t *testing.T) {
	db, dbName := createTempDatabase()
	list := NewTodoList()
	// save the new list
	db.Create(list)

	// read the list from the database
	var savedList TodoList
	db.First(&savedList)

	assert.NotEqual(t, 0, savedList.ID)
	removeDatabase(dbName)
}

func TestNewTodoItem(t *testing.T) {
	db, dbName := createTempDatabase()
	list := NewTodoList()
	// save the new list
	db.Create(list)

	item := NewTodoItem("Buy Milk", list.ID)
	// save the new item
	db.Create(item)

	// read the item from the database
	var savedItem TodoItem
	db.First(&savedItem)

	assert.NotEqual(t, 0, savedItem.ID)
	assert.Equal(t, "Buy Milk", savedItem.Description)

	removeDatabase(dbName)
}
