package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	println("MONGOURI: ", os.Getenv("MONGOURI"))
	envFile, err := godotenv.Read(".env")
	if envFile != nil && err == nil {
		println("MONGOURI: ", envFile["MONGOURI"])
		return envFile["MONGOURI"]
	}

	return os.Getenv("MONGOURI")
}
