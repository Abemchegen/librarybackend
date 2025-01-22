package controller

import (
	"fmt"
	"librarybackend/domain"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	StudentUsecase domain.StudentUseCase
}

func NewStudentController(StudentUsecase domain.StudentUseCase) *StudentController {
	return &StudentController{StudentUsecase: StudentUsecase}
}

func (uc *StudentController) EnterLibrary(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	image := c.Query("image")

	s := domain.Student{Name: name, SchoolID: id, Image: image}
	fmt.Printf(s.Name, s.SchoolID)

	_, err := uc.StudentUsecase.EnterLibrary(s)
	if err.Message != "" {
		c.JSON(400, gin.H{
			"status":  err.Status,
			"message": err.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "student entered successfully",
	})
}

func (uc *StudentController) LeaveLibrary(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	image := c.Query("image")

	s := domain.Student{Name: name, SchoolID: id, Image: image}

	_, err := uc.StudentUsecase.LeaveLibrary(s)
	if err.Message != "" {
		c.JSON(400, gin.H{
			"status":  err.Status,
			"message": err.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "student left successfully",
	})
}

func (uc *StudentController) GetStudentActivity(c *gin.Context) {

	result, err := uc.StudentUsecase.GetStudentActivity()

	if err.Message != "" {
		c.JSON(400, gin.H{
			"status":  err.Status,
			"message": err.Message,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "activity found",

		"data": result})
}

func (uc *StudentController) GetUniqueStudentCountPerDay(c *gin.Context) {

	result, err := uc.StudentUsecase.GetUniqueStudentCountPerDay()

	if err.Message != "" {
		c.JSON(400, gin.H{
			"status":  err.Status,
			"message": err.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "uniques student count retrieved",
		"data":    result,
	})

}

func (uc *StudentController) GetCurrentVisitors(c *gin.Context) {

	result, err := uc.StudentUsecase.GetCurrentVisitors()

	if err.Message != "" {
		c.JSON(400, gin.H{
			"status":  err.Status,
			"message": err.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "current visiters retrieved",
		"data":    result,
	})

}
