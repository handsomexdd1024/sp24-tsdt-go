package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

	return r
}

func main() {
	router := setupRouter()
	err := router.Run("127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
}
