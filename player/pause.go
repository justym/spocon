package player

import (
	"net/http"

	"golang.org/x/exp/errors"
)

const (
	pauseEndpoint   = "https://api.spotify.com/v1/me/player/pause"
	pauseSuccesCode = 204
)

//Pause pause playback
func Pause(client *http.Client) error {
	req, err := http.NewRequest(http.MethodPut, pauseEndpoint, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != pauseSuccesCode {
		return errors.New("Invalid status code: " + string(resp.StatusCode))
	}

	return err
}
