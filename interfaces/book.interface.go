package interfaces

type BookType struct {
	BookID          string `json:"bookID"`
	AuthorID        string `json:"authorID"`
	PublisherID     string `json:"publisherID"`
	Title           string `json:"title"`
	PublicationDate string `json:"publicationDate"`
	ISBN            string `json:"ISBN"`
	Pages           int    `json:"pages"`
	Genre           string `json:"genre"`
	Description     string `json:"description"`
	Price           float64 `json:"price"`
	Quantity        int     `json:"quantity"`
}