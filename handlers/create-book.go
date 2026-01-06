package handlers

import (
	"net/http"
	"randa/book-directory/interfaces"
	"randa/book-directory/storage"

	"github.com/gin-gonic/gin"
)

func CreatingBook(c *gin.Context) {
	var newBook interfaces.BookType

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	for _, checkData := range storage.Books {
		if newBook.BookID == checkData.BookID && newBook.Title == checkData.Title {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "Book already Exists"})
			return
		}
	}
	storage.Books = append(storage.Books, newBook)
	c.IndentedJSON(http.StatusCreated, gin.H{
        "message": "Book added successfully",
        "book":    newBook,
    })
}