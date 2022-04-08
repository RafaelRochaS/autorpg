package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DEBUG string

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DEBUG = os.Getenv("DEBUG")
}
