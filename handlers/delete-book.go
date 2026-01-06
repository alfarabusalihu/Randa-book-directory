package handlers

import (
	"net/http"
	"randa/book-directory/storage"
	"github.com/gin-gonic/gin"
)

func DeleteBookById(c *gin.Context) {
	id := c.Param("BookID")

	if id == "" {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "BookID is required"})
        return
    }

	for i, y := range storage.Books {
		if y.BookID == id {
			storage.Books = append(storage.Books[:i], storage.Books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Book is deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}