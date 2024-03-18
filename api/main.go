package main

import (
	"ciri2-pc-microservice/configs" //add this
	"ciri2-pc-microservice/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	routes.ComponentRoutes(router)

	router.Run("0.0.0.0:6000")
}
