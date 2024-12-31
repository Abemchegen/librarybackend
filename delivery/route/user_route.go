package routers

import (
	"librarybackend/config"
	"librarybackend/delivery/controller"
	"librarybackend/infrastracture"
	"librarybackend/repository"
	"librarybackend/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	timeout := time.Duration(config.ContextTimeout) * time.Second

	repo := repository.NewUserRepository(DB, config.UserCollection)

	tokenGen := infrastracture.NewTokenGenerator()
	passwordSvc := infrastracture.NewPasswordService()

	usecase := usecase.NewUserUseCase(repo, timeout, tokenGen, passwordSvc)

	userController := controller.NewUserController(usecase)

	user := route.Group("/user")
	{

		user.POST("/register", userController.CreateAccount)
		user.POST("/login", userController.Login)

	}

}
