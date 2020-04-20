package player

import (
	"net/http"
)

const statusEndpoint = "https://api.spotify.com/v1/me/player/currently-playing"

func Status(client *http.Client) error {

}
