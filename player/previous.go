package player

import (
	"errors"
	"net/http"
	"strconv"
)

const (
	prevEndpoint    = "https://api.spotify.com/v1/me/player/previous"
	prevSuccessCode = 204
)

func Previous(client *http.Client) error {
	req, err := http.NewRequest(http.MethodPost, prevEndpoint, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != prevSuccessCode {
		return errors.New("Invalid status code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}
