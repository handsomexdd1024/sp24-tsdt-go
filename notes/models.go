package notes

import (
	"gorm.io/gorm"
)

type TodoItem struct {
	gorm.Model
	Description string `gorm:"type:text"`
	TodoListID  uint   `gorm:"not null"`
}

type TodoList struct {
	gorm.Model
	TodoItems []TodoItem `gorm:"foreignKey:TodoListID"`
}
