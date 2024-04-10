// Package main holds the main function that starts the application
package main

import (
	"ciri2-pc-microservice/configs"
	docs "ciri2-pc-microservice/docs"
	"ciri2-pc-microservice/internal/routes"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// @title PC Microservice documentation
// @version 0.0.5
func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = "localhost:" + configs.EnvPort()

	// Run database
	configs.ConnectDB()

	routes.ComponentRoutes(router)
	routes.SwaggerRoutes(router)

	router.Run("0.0.0.0:" + configs.EnvPort())
}
