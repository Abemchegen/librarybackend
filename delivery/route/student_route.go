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

func NewStudentRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	repo := repository.NewStudentRepository(DB, config.StudentCollection)
	usecase := usecase.NewStudentUseCase(repo)
	StudentController := controller.NewStudentController(usecase)

	route.GET("/enterlibrary", StudentController.EnterLibrary)
	route.GET("/leavelibrary", StudentController.LeaveLibrary)

	Student := route.Group("/Student")
	Student.Use(infrastracture.AuthMiddleware())

	{

		Student.GET("/getactivity", StudentController.GetStudentActivity)
		Student.GET("/getcount", StudentController.GetUniqueStudentCountPerDay)
		Student.GET("/currentvisitors", StudentController.GetCurrentVisitors)

	}

}
