package routes

import (
	"randa/book-directory/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/books", handlers.GetAllBooks)
	router.GET("/books/:bookID", handlers.GetBookById)
	router.POST("/books", handlers.CreatingBook)
	router.PUT("/books/:bookID", handlers.UpdateBookById)
	router.DELETE("/books/:bookID", handlers.DeleteBookById)
	router.GET("/books/search", handlers.SearchBook)
}
