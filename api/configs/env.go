// Package configs holds all the logic that is necessary to configure the application
package configs

import (
	"os"

	"github.com/joho/godotenv"
)

// EnvMongoURI returns the MongoDB URI from the environment variables
func EnvMongoURI() string {
	println("MONGOURI: ", os.Getenv("MONGOURI"))
	envFile, err := godotenv.Read(".env")
	if envFile != nil && err == nil {
		println("MONGOURI: ", envFile["MONGOURI"])
		return envFile["MONGOURI"]
	}

	return os.Getenv("MONGOURI")
}

// EnvPort returns the port from the environment variables
func EnvPort() string {
	envFile, err := godotenv.Read(".env")
	if envFile != nil && err == nil {
		return envFile["PC_PORT"]
	}

	return os.Getenv("PC_PORT")
}
