package apifunctions

import (
	"net/http"
	fileReader "randa/book-directory/book-data"
	"randa/book-directory/interfaces"

	"github.com/gin-gonic/gin"
)

func CreatingBook(c *gin.Context) {
	var newBook interfaces.BookType

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	for _, checkData := range fileReader.Books {
		if newBook.BookID == checkData.BookID && newBook.Title == checkData.Title {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "Book already Exists"})
			return
		}
	}
	fileReader.Books = append(fileReader.Books, newBook)
	c.IndentedJSON(http.StatusCreated, gin.H{
        "message": "Book added successfully",
        "book":    newBook,
    })
}