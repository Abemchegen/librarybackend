package controller

import (
	"librarybackend/domain"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookUsecase domain.BookUseCase
}

func NewBookController(BookUsecase domain.BookUseCase) *BookController {
	return &BookController{
		BookUsecase: BookUsecase,
	}
}

func (p *BookController) CreateBook(c *gin.Context) {
	var Book domain.Book
	if err := c.ShouldBindJSON(&Book); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid request", "error": err.Error()})
		return
	}
	createdBook, err := p.BookUsecase.CreateBook(Book)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to create Book", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Book created successfully", "data": createdBook})
}

func (p *BookController) GetAllBook(c *gin.Context) {
	Books, err := p.BookUsecase.GetAllBook()
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to retrieve Books", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Books retrieved successfully", "data": Books})
}

func (p *BookController) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	Book, err := p.BookUsecase.GetBookByID(id)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to retrieve Book", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Book retrieved successfully", "data": Book})
}

func (p *BookController) UpdateBook(c *gin.Context) {
	var Book domain.Book
	if err := c.ShouldBindJSON(&Book); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid request", "error": err.Error()})
		return
	}
	updatedBook, err := p.BookUsecase.UpdateBook(Book)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to update Book", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Book updated successfully", "data": updatedBook})
}

func (p *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	deletedBook, err := p.BookUsecase.DeleteBook(id)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to delete Book", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Book deleted successfully", "data": deletedBook})
}
func (p *BookController) LendBook(c *gin.Context) {
	var request domain.LendBookRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid input", "error": err.Error()})
		return
	}

	record, err := p.BookUsecase.LendBook(
		request.BookID,
		request.StudentID,
		request.StudentName,
		request.LentDate,
		request.DueDate,
		request.LentType,
	)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to lend book", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Book lent successfully", "record": record})
}

func (p *BookController) ReturnBook(c *gin.Context) {

	var request domain.ReturnBookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid input", "error": err.Error()})
		return
	}

	err := p.BookUsecase.ReturnBook(request.BookID, request.StudentID, request.ReturnStatus, request.ReturnCondition)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to return book", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Book returned successfully"})
}
