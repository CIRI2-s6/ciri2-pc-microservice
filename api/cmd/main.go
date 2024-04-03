package main

import (
	"ciri2-pc-microservice/configs" //add this
	"ciri2-pc-microservice/internal/routes"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// Run database
	configs.ConnectDB()

	routes.ComponentRoutes(router)

	router.Run("0.0.0.0:6000")
}
