package router

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	tsdtDB "github.com/handsomexdd1024/sp24-tsdt-go/db"
	"github.com/handsomexdd1024/sp24-tsdt-go/service"
)

func App(dbPath string) *gin.Engine {
	r := gin.Default()

	db, err := tsdtDB.InitDB(dbPath)
	if err != nil {
		panic("failed to connect database")
	}

	ctrl := service.NewController(db)

	r.LoadHTMLGlob(filepath.Join("./templates", "*.tmpl"))

	r.GET("/", ctrl.HomePage)

	r.POST("/new", ctrl.NewList)

	r.GET("/:id/", ctrl.GetList)

	r.POST("/:id/new", ctrl.NewItem)

	r.Static("/static", "./static")

	return r
}
