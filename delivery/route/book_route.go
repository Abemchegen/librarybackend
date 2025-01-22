package routers

import (
	"librarybackend/config"
	"librarybackend/delivery/controller"
	"librarybackend/infrastracture"
	"librarybackend/repository"
	"librarybackend/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBookRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	BookRepository := repository.NewBookRepository(DB, config.BookCollection)
	RecordRepository := repository.NewRecordRepository(DB, config.RecordCollection)

	BookUsecase := usecase.NewBookUseCase(BookRepository, RecordRepository)
	BookController := controller.NewBookController(BookUsecase)

	BookRouter := route.Group("/book")
	BookRouter.Use(infrastracture.AuthMiddleware())
	{
		BookRouter.POST("/", BookController.CreateBook)
		BookRouter.GET("/all", BookController.GetAllBook)
		BookRouter.GET("/", BookController.GetBookByID)
		BookRouter.PUT("/", BookController.UpdateBook)
		BookRouter.DELETE("/", BookController.DeleteBook)
		BookRouter.POST("/lend", BookController.LendBook)
		BookRouter.POST("/return", BookController.ReturnBook)
		BookRouter.GET("/record", BookController.GetRecord)

	}

}
