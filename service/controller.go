package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	. "github.com/handsomexdd1024/sp24-tsdt-go/models"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}

func (ctrl *Controller) getTodoItems(todo_list_id uint) []TodoItem {
	var items []TodoItem
	ctrl.db.Where("todo_list_id = ?", todo_list_id).Find(&items)
	return items
}

func (ctrl *Controller) createTodoList() TodoList {
	list := TodoList{TodoItems: []TodoItem{}}
	ctrl.db.Create(&list)
	return list
}

func (ctrl *Controller) createTodoItem(description string, todoListID uint) TodoItem {
	item := TodoItem{Description: description, TodoListID: todoListID}
	ctrl.db.Create(&item)
	return item
}

func (ctrl *Controller) HomePage(c *gin.Context) {
	apiAddress := ApiAddress{NewItem: "/new"}
	c.HTML(http.StatusOK, "homepage.tmpl", gin.H{
		"Items": nil,
		"Title": "Start a new to-do list",
		"Api":   apiAddress,
	})
}

func (ctrl *Controller) NewList(c *gin.Context) {
	description := c.PostForm("description")
	newList := ctrl.createTodoList()
	ctrl.createTodoItem(description, newList.ID)
	redirectUrl := "/" + strconv.Itoa(int(newList.ID)) + "/"
	c.Redirect(http.StatusFound, redirectUrl)
}

func (ctrl *Controller) GetList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}
	var list TodoList
	ctrl.db.First(&list, id)
	items := ctrl.getTodoItems(list.ID)
	apiAddress := ApiAddress{NewItem: fmt.Sprintf("/%d/new", id)}
	c.HTML(200, "homepage.tmpl", gin.H{
		"Items": items,
		"Title": "Your to-do list",
		"Api":   apiAddress,
	})
}

func (ctrl *Controller) NewItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}
	description := c.PostForm("description")
	ctrl.createTodoItem(description, uint(id))
	redirectUrl := fmt.Sprintf("/%d/", id)
	c.Redirect(http.StatusFound, redirectUrl)
}
