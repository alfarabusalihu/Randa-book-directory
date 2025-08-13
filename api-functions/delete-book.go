package apifunctions

import (
	"net/http"
	fileReader "randa/book-directory/book-data"
	"github.com/gin-gonic/gin"
)

func DeleteBookById(c *gin.Context) {
	id := c.Param("BookID")

	if id == "" {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "BookID is required"})
        return
    }

	for i, y := range fileReader.Books {
		if y.BookID == id {
			fileReader.Books = append(fileReader.Books[:i], fileReader.Books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Book is deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}