package routers

import (
	"librarybackend/config"
	"librarybackend/delivery/controller"
	"librarybackend/repository"
	"librarybackend/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewStudentRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	repo := repository.NewStudentRepository(DB, config.StudentCollection)
	usecase := usecase.NewStudentUseCase(repo)
	StudentController := controller.NewStudentController(usecase)

	Student := route.Group("/Student")
	{

		Student.GET("/enterlibrary", StudentController.EnterLibrary)
		Student.GET("/leavelibrary", StudentController.LeaveLibrary)
		Student.GET("/getactivity", StudentController.GetStudentActivity)
		Student.GET("/getcount", StudentController.GetUniqueStudentCountPerDay)

	}

}
