package main

import (
	"github.com/handsomexdd1024/sp24-tsdt-go/notes"
)

func main() {
	router := notes.App()
	err := router.Run("127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
}
