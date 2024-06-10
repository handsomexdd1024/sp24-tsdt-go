package notes

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// query all todo items from given todo list id
func getTodoItems(db *gorm.DB, todoListID uint) []TodoItem {
	var items []TodoItem
	db.Where("todo_list_id = ?", todoListID).Find(&items)
	return items
}

func App(templatesDir string, dbPath string) *gin.Engine {
	r := gin.Default()

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TodoList{}, &TodoItem{})
	db.Create(&TodoList{})
	var defaultList TodoList
	db.First(&defaultList)

	r.LoadHTMLGlob(filepath.Join(templatesDir, "*.html"))

	r.GET("/", func(c *gin.Context) {
		items := getTodoItems(db, defaultList.ID)
		c.HTML(200, "homepage.html", gin.H{
			"Items": items,
		})
	})

	r.POST("/", func(c *gin.Context) {
		description := c.PostForm("description")
		item := NewTodoItem(description, defaultList.ID)
		db.Create(item)
		c.Redirect(302, "/")
	})
	return r
}
