package pkg

import (
	"os"

	"golang.org/x/exp/errors/fmt"
)

var (
	clientID     string
	clientSecret string
)

func init() {
	clientID = os.Getenv("SPOTIFY_ID")
	clientSecret = os.Getenv("SPOTIFY_KEY")
}

//GetValues return clientID clientSecret
func GetValues() (string, string, error) {
	if clientID == "" || clientSecret == "" {
		err := fmt.Errorf("SPOTIFY_ID or SPOTIFY_KEY is nil")
		return clientID, clientSecret, err
	}
	return clientID, clientSecret, nil
}
