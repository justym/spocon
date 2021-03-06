package pkg

import (
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

const (
	redirectURL = "http://127.0.0.1:14565/oauth/callback"
	scopeUMPS   = "user-modify-playback-state"
	scopeURCP   = "user-read-currently-playing"
)

//NewConfig return singleton state imformation for OAuth2.0 configurations.
func NewConfig() *oauth2.Config {
	id, key, err := GetValues()
	if err != nil {
		log.Fatal(err)
	}

	config := &oauth2.Config{
		ClientID:     id,
		ClientSecret: key,
		Endpoint:     spotify.Endpoint,
		Scopes:       []string{scopeUMPS, scopeURCP},
	}

	return config
}
