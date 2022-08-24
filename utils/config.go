package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DEBUG string
var COMBAT_PAUSE_TIME int

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DEBUG = os.Getenv("DEBUG")
	COMBAT_PAUSE_TIME, err = strconv.Atoi(os.Getenv("COMBAT_PAUSE_TIME"))

	if err != nil {
		COMBAT_PAUSE_TIME = 1
	}
}
