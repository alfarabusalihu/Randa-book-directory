package main

import (
	apifunctions "randa/book-directory/api-functions"
	fileReader "randa/book-directory/book-data"

	"github.com/gin-gonic/gin"
)

func main() {
	fileReader.LoadBooksFromFile()

	router := gin.Default()
	router.GET("/books", apifunctions.GetAllBooks)
	router.GET("/books/:bookID", apifunctions.GetBookById)
	router.POST("/books", apifunctions.CreatingBook)
	router.PUT("/books/:bookID", apifunctions.UpdateBookById)
	router.DELETE("/books/:bookID", apifunctions.DeleteBookById)
	router.GET("/books/search", apifunctions.SearchBook)
	router.Run("0.0.0.0:8080")
}
