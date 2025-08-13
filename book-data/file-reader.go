package fileReader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"randa/book-directory/interfaces"
)

var Books []interfaces.BookType

func LoadBooksFromFile() {
	// Open book-data
	workingDir, _ := os.Getwd()

	// Construct the correct path
	filePath := filepath.Join(workingDir, "book-data", "book-data.json")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	fileContents, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse JSON data and store to books var
	err = json.Unmarshal(fileContents, &Books)
	if len(Books)==0{
		fmt.Println("Error books is empty:", err)
	}
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
}