package pkg

import (
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var config *oauth2.Config

const redirectURL = "http://127.0.0.1:14565/oauth/callback"
const scopeUMPS = "user-modify-playback-state"

//NewConfig return singleton state imformation for OAuth2.0 configurations.
func NewConfig() *oauth2.Config {
	if config != nil {
		return config
	}

	LoadEnv()
	id, key := GetValues()
	if id == "" || key == "" {
		log.Fatalln("id or key is none")
	}

	config = &oauth2.Config{
		ClientID:     id,
		ClientSecret: key,
		Endpoint:     spotify.Endpoint,
		Scopes:       []string{scopeUMPS},
	}

	return config
}
