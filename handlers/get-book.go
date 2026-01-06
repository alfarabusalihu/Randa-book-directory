package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"randa/book-directory/storage"
)

func GetBookById(c *gin.Context) {
	id := c.Param("BookID")

	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "BookID is required"})
		return
	}

	for _, findBook := range storage.Books {
		if findBook.BookID == id {
			c.IndentedJSON(http.StatusOK, findBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book isn't found"})
}
