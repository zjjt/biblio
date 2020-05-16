package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//GetEnvVariable load env variables from .env file usage in development only
func GetEnvVariable(varname string) string {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Couldnt load .env file")
	}
	return os.Getenv(varname)
}
