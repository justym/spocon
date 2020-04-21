package player

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const statusEndpoint = "https://api.spotify.com/v1/me/player/currently-playing"

//Status gets currently playback
func Status(client *http.Client) error {
	req, err := http.NewRequest(http.MethodGet, statusEndpoint, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var out CurrentStatus
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return err
	}

	fmt.Println(out.Item.Name)

	return nil
}
