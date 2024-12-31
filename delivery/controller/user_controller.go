package controller

import (
	"fmt"
	"librarybackend/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUseCase
}

func NewUserController(userUsecase domain.UserUseCase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

func (uc *UserController) CreateAccount(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := uc.UserUsecase.CreateAccount(user)
	if err.Message != "" {
		c.JSON(400, gin.H{
			"status":  err.Status,
			"message": err.Message,
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Account created successfully",
		"data":    result})
}

func (uc *UserController) Login(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := uc.UserUsecase.Login(user)
	if err.Message != "" {
		c.JSON(400, gin.H{
			"status":  err.Status,
			"message": err.Message,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Login successful",

		"data": result})
}

func (uc *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := uc.UserUsecase.GetByID(id)

	if err.Message != "" {
		c.JSON(400, gin.H{
			"status":  err.Status,
			"message": err.Message,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "User found",

		"data": result})
}

func (uc *UserController) UpdateProfile(c *gin.Context) {

	var user domain.User

	userid := c.GetString("user_id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(userid)
	result, err := uc.UserUsecase.UpdateProfile(userid, user)

	if err.Message != "" {
		c.JSON(400, gin.H{
			"status":  err.Status,
			"message": err.Message,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Profile updated successfully",

		"data": result})
}

func (uc *UserController) GetAllUser(c *gin.Context) {
	result, err := uc.UserUsecase.GetAllUser()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(200, gin.H{"result": result})

}

func (uc *UserController) GetMe(c *gin.Context) {
	userID := c.GetString("user_id")
	result, err := uc.UserUsecase.GetUserByID(userID)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{

		"status":  200,
		"message": "User found",
		"user":    result})

}
