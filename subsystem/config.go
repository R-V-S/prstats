package subsystem

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (err error) {
	environment := os.Getenv("environment")
	if environment == "" {
		environment = "dev"
	}
	envFilename := ".env." + environment
	err = godotenv.Load(envFilename)
	if err != nil {
		log.Fatal("Error loading ", envFilename, " file")
	}
	return nil
}
