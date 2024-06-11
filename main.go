package main

import (
	"github.com/gin-gonic/gin"
	"github.com/handsomexdd1024/sp24-tsdt-go/notes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := notes.App("./tsdt.db")
	err := router.Run("127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
}
