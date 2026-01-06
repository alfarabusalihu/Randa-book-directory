package handlers

import (
	"net/http"
	"randa/book-directory/interfaces"
	"randa/book-directory/storage"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func SearchBook(c *gin.Context) {
	searchedKeyword := strings.TrimSpace(strings.ToLower(c.Query("keyword")))

	if searchedKeyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Keyword is required"})
		return
	}

	var matchingBooks []interfaces.BookType
	results := make(chan interfaces.BookType)
	var wg sync.WaitGroup

	for _, b := range storage.Books {
		wg.Add(1)
		go func(b interfaces.BookType) {
			defer wg.Done()
			if strings.Contains(strings.ToLower(b.Title), searchedKeyword) || strings.Contains(strings.ToLower(b.Description), searchedKeyword) {
				results <- b
			}
		}(b)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for book := range results {
		matchingBooks = append(matchingBooks, book)
	}

	if len(matchingBooks) > 0 {
		c.JSON(http.StatusOK, matchingBooks)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}
}
