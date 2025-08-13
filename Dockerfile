# initializing the base image in this case I use golang base iage
FROM golang:1.22-alpine

# initializing a seperate directory
WORKDIR /app

# Copying go files into the directory
COPY go.mod go.sum ./

# installing go modules insides the app directory
RUN go mod download

# copy other files inside the DIR
COPY . .

# build the container
RUN CGO_ENABLED=0 GOOS=linux go build -o /randoli-book-system

EXPOSE 8080

#Run file
CMD ["/randa-book-system"]
