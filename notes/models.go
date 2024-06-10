package notes

import (
	"gorm.io/gorm"
)

type TodoItem struct {
	gorm.Model
	Description string `gorm:"type:text"`
	TodoListID  uint   `gorm:"not null"`
}

func NewTodoItem(description string, todoListID uint) *TodoItem {
	return &TodoItem{
		Description: description,
		TodoListID:  todoListID,
	}
}

type TodoList struct {
	gorm.Model
	TodoItems []TodoItem `gorm:"foreignKey:TodoListID"`
}

func NewTodoList() *TodoList {
	return &TodoList{}
}
