package routers

import (
	"librarybackend/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	NewUserRouter(route, config, DB)
	NewBookRouter(route, config, DB)

}
