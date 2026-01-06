package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"randa/book-directory/interfaces"
	"randa/book-directory/storage"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatingBook(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/books", CreatingBook)

	// Initial length of Books slice
	initialLength := len(storage.Books)

	newBook := interfaces.BookType{
		BookID:   "1",
		Title:    "Test Book",
		AuthorID: "1",
	}
	jsonValue, _ := json.Marshal(newBook)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, initialLength+1, len(storage.Books))
	assert.Equal(t, newBook, storage.Books[len(storage.Books)-1])
}
