package player

import (
	"errors"
	"net/http"
	"strconv"
)

const (
	startEndpoint    = "https://api.spotify.com/v1/me/player/play"
	startSuccessCode = 204
)

//Start starts/resume playback
func Start(client *http.Client) error {
	req, err := http.NewRequest(http.MethodPut, startEndpoint, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != startSuccessCode {
		return errors.New("Invalid Status Code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}
