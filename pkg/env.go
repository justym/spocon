package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	clientID     string
	clientSecret string
)

//LoadEnv loads value from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Cannot read .env file:", err)
	}
	clientID = os.Getenv("ID")
	clientSecret = os.Getenv("KEY")
}

//GetValues return clientID clientSecret
func GetValues() (string, string) {
	return clientID, clientSecret
}
