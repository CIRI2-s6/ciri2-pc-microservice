package routes

import (
	"ciri2-pc-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func ComponentRoutes(router *gin.Engine) {
	router.POST("/component", controllers.BatchCreateComponent())
	router.GET("/component/:id", controllers.GetComponent())
}
