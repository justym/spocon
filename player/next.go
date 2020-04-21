package player

import (
	"net/http"
	"strconv"

	"golang.org/x/exp/errors"
)

const (
	nextEndpoint    = "https://api.spotify.com/v1/me/player/next"
	nextSuccessCode = 204
)

func Next(client *http.Client) error {
	req, err := http.NewRequest(http.MethodPost, nextEndpoint, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != nextSuccessCode {
		return errors.New("Invalid status code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}
