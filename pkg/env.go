package pkg

import (
	"os"
)

var (
	clientID     string
	clientSecret string
)

//LoadEnv loads value from .env file
func LoadEnv() {
	clientID = os.Getenv("SPOTIFY_ID")
	clientSecret = os.Getenv("SPOTIFY_KEY")
}

//GetValues return clientID clientSecret
func GetValues() (string, string) {
	return clientID, clientSecret
}
