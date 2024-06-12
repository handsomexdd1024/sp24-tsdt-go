package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tsdtModels "github.com/handsomexdd1024/sp24-tsdt-go/models"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}

func (ctrl *Controller) getTodoItems(todo_list_id uint) []tsdtModels.TodoItem {
	var items []tsdtModels.TodoItem
	ctrl.db.Where("todo_list_id = ?", todo_list_id).Find(&items)
	return items
}

func (ctrl *Controller) createTodoList() tsdtModels.TodoList {
	list := tsdtModels.TodoList{TodoItems: []tsdtModels.TodoItem{}}
	ctrl.db.Create(&list)
	return list
}

func (ctrl *Controller) createTodoItem(description string, todoListID uint) tsdtModels.TodoItem {
	item := tsdtModels.TodoItem{Description: description, TodoListID: todoListID}
	ctrl.db.Create(&item)
	return item
}

func (ctrl *Controller) HomePage(c *gin.Context) {
	apiAddress := tsdtModels.ApiAddress{NewItem: "/new"}
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
	var list tsdtModels.TodoList
	ctrl.db.First(&list, id)
	items := ctrl.getTodoItems(list.ID)
	apiAddress := tsdtModels.ApiAddress{NewItem: fmt.Sprintf("/%d/new", id)}
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
