package apifunctions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	fileReader "randa/book-directory/book-data"
	"randa/book-directory/interfaces"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatingBook(t *testing.T){
	gin.SetMode(gin.TestMode)

	fileReader.Books = []interfaces.BookType{}

	newBook:=interfaces.BookType{
		BookID: "ef0123456789",
		AuthorID: "bcde2345-fg67-8901-hijk-lmnopqrstuvw",
		PublisherID: "8765bcde-3210-ghij-klmn-678901opqrst",
		Title: "10840989",
		PublicationDate: "1949-06-08",
		ISBN: "9780451524935",
		Pages: 328,
		Genre: "Dysto",
		Description: "A novel about a totalitarian regime and surveillance.",
		Price: 14.99,
		Quantity: 6,
	}

	bookData, _ :=json.Marshal(newBook)

	req, _ :=http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/books", CreatingBook)
	router.ServeHTTP(w, req)

    // to ensure the staus on both testing and handler files are correct
	assert.Equal(t, http.StatusCreated, w.Code)
    
	// make sure the book is inserted into the array
	for _, book := range fileReader.Books {
		if book.BookID == newBook.BookID {
			assert.Equal(t, newBook.Title, book.Title, "The comapred titles of the books should be same")
			break
		}
	}	
}