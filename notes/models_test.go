package notes_test

import (
	"os"
	"testing"
	"time"

	"github.com/handsomexdd1024/sp24-tsdt-go/notes"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createTempDatabase() (*gorm.DB, string) {
	// create a temporary database with current time
	// to avoid conflict with other tests
	dbName := time.Now().Format("test20060102150405.db")
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&notes.TodoList{}, &notes.TodoItem{})
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
	list := notes.NewTodoList()
	// save the new todo list
	db.Create(list)

	// read the todo list from the database
	var savedList notes.TodoList
	db.First(&savedList)

	assert.NotEqual(t, 0, savedList.ID)
	removeDatabase(dbName)
}

func TestNewTodoItem(t *testing.T) {
	db, dbName := createTempDatabase()
	list := notes.NewTodoList()
	// save the new todo list
	db.Create(list)

	item := notes.NewTodoItem("Buy Milk", list.ID)
	// save the new todo item
	db.Create(item)

	// read the todo item from the database
	var savedItem notes.TodoItem
	db.First(&savedItem)

	assert.NotEqual(t, 0, savedItem.ID)
	assert.Equal(t, "Buy Milk", savedItem.Description)

	removeDatabase(dbName)
}
