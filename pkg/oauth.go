package pkg

import (
	"errors"

	a2ns "github.com/nmrshll/oauth2-noserver"
	"golang.org/x/oauth2"
)

var authorizedClient *a2ns.AuthorizedClient

//NewClient returns authorized singleton http client
func NewClient(conf *oauth2.Config) (*a2ns.AuthorizedClient, error) {
	if conf == nil {
		return nil, errors.New("oauth2.Config is not defined")
	}

	if authorizedClient == nil {
		authorizedClient, err := a2ns.AuthenticateUser(conf)
		if err != nil {
			return nil, err
		}
		return authorizedClient, err
	}

	return authorizedClient, nil
}
