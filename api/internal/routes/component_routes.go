// Package routes holds the routes for the application
package routes

import (
	"ciri2-pc-microservice/internal/controllers"

	"github.com/gin-gonic/gin"
)

func ComponentRoutes(componentGroup *gin.Engine) {
	component := new(controllers.ComponentController)
	componentGroup.POST("/component", component.BatchCreateComponent())
	componentGroup.POST("/component/check", component.CheckAlreadyExistingComponents())
	componentGroup.GET("/component/:id", component.GetComponent())
	componentGroup.GET("/component", component.GetPaginatedComponent())
}
