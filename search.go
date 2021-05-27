package tmdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type searchClient struct {
	*Client
}

// SearchClient returns a searchClient
func SearchClient(c *Client) *searchClient {
	return &searchClient{c}
}

// Movies searches for movies based on passed string, returns a Movies struct with MovieResults embedded
func (s *searchClient) Movies(name string) (*Movies, error) {
	if len(name) < 1 {
		return nil, errors.New("query string too short")
	}
	query := fmt.Sprintf("%s%s%s%s%s&query=%s", baseURL, searchURL, movieURL, api, s.apiKey, url.QueryEscape(name))

	res, err := s.http.Get(query)
	if err != nil {
		return nil, err
	}

	results := &Movies{}
	err = json.NewDecoder(res.Body).Decode(results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// TvClient searches for movies based on passed string, returns a Shows struct with ShowResults embedded
func (s *searchClient) TV(name string) (*Shows, error) {
	if len(name) < 1 {
		return nil, errors.New("query string too short")
	}
	query := fmt.Sprintf("%s%s%s%s%s&query=%s", baseURL, searchURL, tvURL, api, s.apiKey, url.QueryEscape(name))

	res, err := s.http.Get(query)
	if err != nil {
		return nil, err
	}

	results := &Shows{}
	err = json.NewDecoder(res.Body).Decode(results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

type Movies struct {
	Page         int            `json:"page"`
	Results      []MovieResults `json:"results"`
	TotalResults int            `json:"total_results"`
	TotalPages   int            `json:"total_pages"`
}

type MovieResults struct {
	PosterPath       string  `json:"poster_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	GenreIds         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Title            string  `json:"title"`
	BackdropPath     string  `json:"backdrop_path"`
	Popularity       float64 `json:"popularity"`
	VoteCount        int     `json:"vote_count"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
}

type Shows struct {
	Page         int           `json:"page"`
	Results      []ShowResults `json:"results"`
	TotalResults int           `json:"total_results"`
	TotalPages   int           `json:"total_pages"`
}
type ShowResults struct {
	PosterPath       string   `json:"poster_path"`
	Popularity       float64  `json:"popularity"`
	ID               int      `json:"id"`
	BackdropPath     string   `json:"backdrop_path"`
	VoteAverage      float64  `json:"vote_average"`
	Overview         string   `json:"overview"`
	FirstAirDate     string   `json:"first_air_date"`
	OriginCountry    []string `json:"origin_country"`
	GenreIds         []int    `json:"genre_ids"`
	OriginalLanguage string   `json:"original_language"`
	VoteCount        int      `json:"vote_count"`
	Name             string   `json:"name"`
	OriginalName     string   `json:"original_name"`
}
