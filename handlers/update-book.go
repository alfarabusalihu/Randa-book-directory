package handlers

import (
	"net/http"
	"randa/book-directory/interfaces"
	"randa/book-directory/storage"

	"github.com/gin-gonic/gin"
)

func UpdateBookById(c *gin.Context) {
	id := c.Param("BookID")

	for i, x := range storage.Books {
		if x.BookID == id {
			var updatedBook interfaces.BookType
			if err := c.BindJSON(&updatedBook); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
				return
			}

			if updatedBook.Title != "" {
				x.Title = updatedBook.Title
			}
			if updatedBook.AuthorID != "" {
				x.AuthorID = updatedBook.AuthorID
			}
			if updatedBook.PublisherID != "" {
				x.PublisherID = updatedBook.PublisherID
			}
			if updatedBook.PublicationDate != "" {
				x.PublicationDate = updatedBook.PublicationDate
			}
			if updatedBook.ISBN != "" {
				x.ISBN = updatedBook.ISBN
			}
			if updatedBook.Pages != 0 {
				x.Pages = updatedBook.Pages
			}
			if updatedBook.Genre != "" {
				x.Genre = updatedBook.Genre
			}
			if updatedBook.Description != "" {
				x.Description = updatedBook.Description
			}
			if updatedBook.Price != 0 {
				x.Price = updatedBook.Price
			}
			if updatedBook.Quantity != 0 {
				x.Quantity = updatedBook.Quantity
			}

			storage.Books[i] = x
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Book updated successfully",
				"book":    x,
			})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}