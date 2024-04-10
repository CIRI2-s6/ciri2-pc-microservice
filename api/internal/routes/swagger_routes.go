package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoutes(swaggerGroup *gin.Engine) {

	swaggerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
