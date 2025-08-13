# Randoli-book-system

This is a backend service written in Go to maintain a book system.

## Table of Contents

- [Description](#description)
- [Steps](#Steps)
- [Configuration](#configuration)
- [Prerequisites](#prerequisites)
- [API Requests](#api_Requests)

## Description

This project contains the backend part of a book system. The program allows users to perform CRUD operations using API endpoints, with features like searching for a book.

You can retrieve all books, with or without query parameters, to filter the results. The system supports pagination using offset and limit to control the number of books displayed.

You can also retrieve or delete a book by specifying its ID in the URL. Creating a book is done by providing a correctly structured JSON object, and you can update specific fields of a book.

Gin is used to define API endpoints, parse requests, send responses, handle errors, and start the server from main.go. Docker is used to containerize the project, creating a base image based on the Go image, with the rest of the application built on top of it.

## Steps

- Download and install Go language.

- Build the go.mod and go.sum files.

- Create a separate .json file to store the data.

- Implement a file-reader method to store books in a shared variable.

- Write the necessary API functions.

- Run the file-reader function and route configuration in main.go.

- Run main.go to ensure the endpoints work correctly.

 


### Prerequisites

- Go version: `Go 1.22.0`
- Docker version `26.0.0`

### Configuration

Since this project is containerized with Docker, it can be started with a few simple steps:

- Run docker-compose build
 to build the image from the Dockerfile, and use the port specified in docker-compose.yml.

- Run docker-compose up  
to start the container. The container will be initialized with the port and name.

- Run docker-compose stop 
 to stop the container.

The image and container can be managed using PowerShell or Git with these commands. Ensure that the host port is correctly set to 8080:8080:

- export PORT=8080
- docker build --tag randoli-book-system . 
- docker run -p 8080:8080 randoli-book-system

Once the container is running, API requests can be sent/received using tools like Postman.

The testing function was created to clarify if the function is creating a book, as for that go defaultly have a go test method. In this project we run a verbose command to test the create-book function and have a detailed output.
-  go test ./api-functions -verbose

### API_Requests

-Base URL
- http://localhost:8080

- router.GET("/books")
Example request: http://localhost:8080/books
With query params: http://localhost:8080/books?offset=2&limit=10 

- router.GET("/books/:BookID")
Example request: http://localhost:8080/books/0j1k2l3m-0123-4567-2345-678901234567

- router.POST("/books")
Example request: http://localhost:8080/books

- router.PUT("/books/:BookID")
Example request: http://localhost:8080/books/0j1k2l3m-0123-4567-2345-678901234567

- router.DELETE("/books/:BookID")
Example request: http://localhost:8080/books/0j1k2l3m-0123-4567-2345-678901234567

- router.GET("books/search")
Example request: http://localhost:8080/books?keyword=and


