package main

import (
	"randa/book-directory/routes"
	"randa/book-directory/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	storage.LoadBooksFromFile()

	router := gin.Default()
	routes.SetupRouter(router)
	router.Run("0.0.0.0:8080")
}
