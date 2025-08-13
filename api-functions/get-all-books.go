package apifunctions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	fileReader "randa/book-directory/book-data"
	"strconv"
)

func GetAllBooks(c *gin.Context) {
	setOffset := c.Query("offset")
	setLimit := c.Query("limit")

	if setOffset != "" || setLimit != "" {

		offset, err := strconv.Atoi(setOffset)
		if err != nil || offset < 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid offset value"})
			return
		}

		limit, err := strconv.Atoi(setLimit)
		if err != nil || limit <= 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Limit must have a proper value"})
			return
		}

		if offset >= len(fileReader.Books) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Offset exceeds available books"})
			return
		}

		if offset+limit > len(fileReader.Books) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": " limit should't exceed the length books store has"})
		}

		slicedObjects := fileReader.Books[offset : offset+limit]
		c.IndentedJSON(http.StatusOK, slicedObjects)
	} else {
		c.IndentedJSON(http.StatusOK, fileReader.Books)
	}
}
