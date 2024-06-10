package notes

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

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

func App(dbPath string) *gin.Engine {
	r := gin.Default()

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TodoList{}, &TodoItem{})

	r.LoadHTMLGlob(filepath.Join("./templates", "*.html"))

	r.GET("/", func(c *gin.Context) {
		apiAddress := ApiAddress{NewItem: "/new"}
		c.HTML(http.StatusOK, "homepage.html", gin.H{
			"Items": nil,
			"Title": "Start a new to-do list",
			"Api":   apiAddress,
		})
	})

	r.POST("/new", func(c *gin.Context) {
		description := c.PostForm("description")
		newList := TodoList{TodoItems: []TodoItem{}}
		db.Create(&newList)
		newItem := TodoItem{Description: description, TodoListID: newList.ID}
		db.Create(&newItem)
		redirectUrl := fmt.Sprintf("/%d/", newList.ID)
		c.Redirect(http.StatusFound, redirectUrl)
	})

	r.GET("/:id/", func(c *gin.Context) {
		id := c.Param("id")
		var list TodoList
		db.First(&list, id)
		items := getTodoItems(db, list.ID)
		apiAddress := ApiAddress{NewItem: fmt.Sprintf("/%s/new", id)}
		c.HTML(200, "homepage.html", gin.H{
			"Items": items,
			"Title": "Your to-do list",
			"Api":   apiAddress,
		})
	})

	r.POST("/:id/new", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		}
		description := c.PostForm("description")
		newItem := TodoItem{Description: description, TodoListID: uint(id)}
		db.Create(&newItem)
		redirectUrl := fmt.Sprintf("/%d/", id)
		c.Redirect(302, redirectUrl)
	})

	r.Static("/static", "./static")

	return r
}
