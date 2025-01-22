package usecase

import (
	"errors"
	"fmt"
	"librarybackend/domain"
)

type BookUseCase struct {
	BookRepository   domain.BookRepository
	RecordRepository domain.RecordRepository
}

func NewBookUseCase(BookRepository domain.BookRepository, RecordRepository domain.RecordRepository) domain.BookUseCase {
	return &BookUseCase{
		BookRepository:   BookRepository,
		RecordRepository: RecordRepository,
	}
}

func (p *BookUseCase) CreateBook(Book domain.Book) (domain.Book, error) {

	fmt.Println("Name:", Book.Name)
	fmt.Println("Author:", Book.Author)
	fmt.Println("Book ID:", Book.Bookid)
	fmt.Println("Course:", Book.Course)
	fmt.Println("Publication Date:", Book.PublicationDate)
	fmt.Println("Quantity:", Book.Quantity)
	if Book.Name == "" || Book.Author == "" || Book.Bookid == "" || Book.Course == "" || Book.PublicationDate == "" {
		return domain.Book{}, errors.New("fill out all book details please")
	}
	if Book.Quantity == 0 {
		return domain.Book{}, errors.New("atleast one book should be available")
	}

	Book, err := p.BookRepository.CreateBook(Book)
	return Book, err
}

func (p *BookUseCase) GetAllBook() ([]domain.Book, error) {
	return p.BookRepository.GetAllBook()
}

func (p *BookUseCase) GetBookByID(id string) (domain.Book, error) {
	return p.BookRepository.GetBookByID(id)
}

func (p *BookUseCase) UpdateBook(Book domain.Book) (domain.Book, error) {
	if Book.Name == "" || Book.Author == "" || Book.Quantity == 0 || Book.Bookid == "" || Book.Course == "" || Book.PublicationDate == "" {
		return domain.Book{}, errors.New("fill out all book details please")
	}
	_, err := p.BookRepository.UpdateBook(Book)
	return Book, err
}

func (p *BookUseCase) DeleteBook(id string) (domain.Book, error) {
	Book, err := p.BookRepository.GetBookByID(id)
	if err != nil {
		return domain.Book{}, err
	}
	_, err = p.BookRepository.DeleteBook(id)
	return Book, err
}

func (p *BookUseCase) LendBook(bookid string, studentid string, studentname string, lentdate string, duedate string, lenttype string) (domain.Record, error) {
	book, err := p.BookRepository.GetBookByID(bookid)
	if err != nil {
		return domain.Record{}, err
	}
	if book.Quantity <= 0 {
		return domain.Record{}, errors.New("book is not available")
	}
	err = p.BookRepository.LendBook(bookid, studentid, studentname)
	if err != nil {
		return domain.Record{}, err
	}

	student := domain.Student{
		SchoolID: studentid,
		Name:     studentname,
	}

	record, err := p.RecordRepository.CreateRecord(student, book, lentdate, duedate, lenttype, "Not Returned", "")

	if err != nil {
		return domain.Record{}, err
	}

	return record, nil
}

func (p *BookUseCase) ReturnBook(bookid string, studentid string, returndate string, returnstatus string, returncondition string) error {

	record, err := p.RecordRepository.GetRecordByID(studentid, bookid)

	if err != nil {
		return err
	}

	if record.ReturnStatus == "Returned" {
		return errors.New("book is already returned")
	}

	err = p.BookRepository.ReturnBook(bookid, studentid)
	if err != nil {
		return err
	}

	_, err = p.RecordRepository.UpdateRecord(studentid, bookid, returndate, returnstatus, returncondition)

	return err
}
func (p *BookUseCase) GetRecord() ([]domain.Record, error) {
	return p.RecordRepository.GetAllRecord()
}

func (p *BookUseCase) BooksBorrowed(id string) ([]domain.Record, domain.ErrorResponse) {

	result, err := p.RecordRepository.BooksBorrowed(id)

	if err != nil {
		return []domain.Record{}, domain.ErrorResponse{Message: "books this student borrowed can't be retrieved", Status: 404}
	}

	return result, domain.ErrorResponse{}
}
