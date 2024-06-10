package notes

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Content string
}

// Item constructor
func NewItem(content string) *Item {
	return &Item{
		Content: content,
	}
}

type List struct {
	gorm.Model
	Items []Item
}

// List constructor
func NewList() *List {
	return &List{}
}

// AddItem adds an item to the list
func (l *List) AddItem(item *Item) {
	l.Items = append(l.Items, *item)
}
