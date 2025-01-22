package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID              primitive.ObjectID `json:"mongo_id" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	Author          string             `json:"author" bson:"author"`
	Course          string             `json:"course" bson:"course"`
	PublicationDate string             `json:"publicationdate" bson:"publicationdate"`
	Quantity        int                `json:"quantity" bson:"quantity"`
	Bookid          string             `json:"bookid" bson:"bookid"`
	Image           string             `json:"image" bson:"image"`
}

type BookUseCase interface {
	CreateBook(Book Book) (Book, error)
	GetAllBook() ([]Book, error)
	GetBookByID(id string) (Book, error)
	UpdateBook(Book Book) (Book, error)
	DeleteBook(id string) (Book, error)
	LendBook(bookid string, studentid string, studentname string, lentdate string, duedate string, lenttype string) (Record, error)
	ReturnBook(bookid string, studentid string, returndate string, returnstatus string, returncondition string) error
	GetRecord() ([]Record, error)
	BooksBorrowed(id string) ([]Record, ErrorResponse)
}

type BookRepository interface {
	CreateBook(Book Book) (Book, error)
	GetAllBook() ([]Book, error)
	GetBookByID(id string) (Book, error)
	UpdateBook(Book Book) (Book, error)
	DeleteBook(id string) (Book, error)
	LendBook(bookid string, studentid string, studentname string) error
	ReturnBook(bookid string, studentid string) error
}
