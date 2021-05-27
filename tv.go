package tmdb

import (
	"encoding/json"
	"fmt"
)

type tvClient struct {
	*Client
}

// TvClient returns a tvClient
func TvClient(c *Client) *tvClient {
	return &tvClient{c}
}

// GetDetails gets details on a TvClient show based on the passed id.
// https://developers.themoviedb.org/3/tv/get-tv-details
func (t *tvClient) GetDetails(id int) (*TVDetails, error) {
	query := fmt.Sprintf("%s%s/%d%s%s", baseURL, tvURL, id, api, t.apiKey)
	res, err := t.http.Get(query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	details := &TVDetails{}
	err = json.NewDecoder(res.Body).Decode(details)
	if err != nil {
		return nil, err
	}

	return details, nil
}

// GetVideos returns a list of videos, trailers, intros, etc, based on id
func (t *tvClient) GetVideos(id int) (*Videos, error) {
	query := fmt.Sprintf("%s%s/%d/videos%s%s", baseURL, tvURL, id, api, t.apiKey)
	res, err := t.http.Get(query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	details := &Videos{}
	err = json.NewDecoder(res.Body).Decode(details)
	if err != nil {
		return nil, err
	}

	return details, nil
}
