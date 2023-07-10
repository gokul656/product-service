package configs

import (
	"log"
	"os"

	env "github.com/gokul656/product-service/app/types"

	"github.com/joho/godotenv"
)

func GetEnv() env.Env {
	err := godotenv.Load()
	if err != nil {
		log.Panicln("unable to load env")
	}

	return env.Env{
		Port:       os.Getenv("PORT"),
		DbURL:      os.Getenv("DB_URL"),
		DbName:     os.Getenv("DB_NAME"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
}
