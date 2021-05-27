package tmdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type movieClient struct {
	*Client
}

// MovieClient returns a movieClient
func MovieClient(c *Client) *movieClient {
	return &movieClient{c}
}

// GetDetails gets details on a TvClient show based on the passed id.
// https://developers.themoviedb.org/3/movies/get-movie-details
func (m *movieClient) GetDetails(ctx context.Context, id int) (*MovieDetails, error) {
	//apiKey := url.QueryEscape(m.apiKey)
	query := fmt.Sprintf("%s%s/%d%s%s", baseURL, movieURL, id, api, m.apiKey)

	req, err := m.newRequest(ctx, http.MethodGet, query)
	if err != nil {
		return nil, err
	}
	res, err := m.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	details := &MovieDetails{}
	if err = json.NewDecoder(res.Body).Decode(details); err != nil {
		return nil, err
	}

	return details, nil
}

// GetAlternativeTitles returns alternative titles to a given movie
func (m *movieClient) GetAlternativeTitles(ctx context.Context, id int) (*AlternativeTitles, error) {
	query := fmt.Sprintf("%s%s/%d%s%s%s", baseURL, movieURL, id, alternativeTitlesURL, api, m.apiKey)
	req, err := m.newRequest(ctx, http.MethodGet, query)
	if err != nil {
		return nil, err
	}

	res, err := m.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	titles := &AlternativeTitles{}
	if err := json.NewDecoder(res.Body).Decode(titles); err != nil {
		return nil, err
	}

	return titles, nil
}

// GetExternalIDs return movie ids found on different prodivders, such as IMDB, Facebook, Instagram, Twitter
func (m *movieClient) GetExternalIDs(ctx context.Context, id int) (*MoviesExternalIDs, error) {
	query := fmt.Sprintf("%s%s/%d%s%s%s", baseURL, movieURL, id, externalIDsURL, api, m.apiKey)
	req, err := m.newRequest(ctx, http.MethodGet, query)
	if err != nil {
		return nil, err
	}

	res, err := m.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	ids := &MoviesExternalIDs{}
	if err := json.NewDecoder(res.Body).Decode(ids); err != nil {
		return nil, err
	}

	return ids, nil
}

// GetVideos returns list of videos that have been added to a movie, such as a trailer
func (m *movieClient) GetVideos(ctx context.Context, id int) (*Videos, error) {
	query := fmt.Sprintf("%s%s/%d/videos%s%s", baseURL, movieURL, id, api, m.apiKey)

	req, err := m.newRequest(ctx, http.MethodGet, query)
	if err != nil {
		return nil, err
	}

	res, err := m.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	videos := &Videos{}
	if err := json.NewDecoder(res.Body).Decode(videos); err != nil {
		return nil, err
	}

	return videos, nil
}

// GetNowPlaying returns list of movies currently in theaters
func (m *movieClient) GetNowPlaying(ctx context.Context) (*NowPlaying, error) {
	query := fmt.Sprintf("%s%s%s%s%s", baseURL, movieURL, nowPlayingURL, api, m.apiKey)
	req, err := m.newRequest(ctx, http.MethodGet, query)
	if err != nil {
		return nil, err
	}

	res, err := m.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	playing := &NowPlaying{}
	if err := json.NewDecoder(res.Body).Decode(playing); err != nil {
		return nil, err
	}

	return playing, nil
}

// GetPopular returns current popular movies in TMDB, updated daily
func (m *movieClient) GetPopular(ctx context.Context) (*Popular, error) {
	query := fmt.Sprintf("%s%s%s%s%s", baseURL, movieURL, popularURL, api, m.apiKey)
	req, err := m.newRequest(ctx, http.MethodGet, query)
	if err != nil {
		return nil, err
	}

	res, err := m.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	popular := &Popular{}
	if err := json.NewDecoder(res.Body).Decode(popular); err != nil {
		return nil, err
	}

	return popular, nil
}

// GetTopRated returns top rated movies in TMDB
func (m *movieClient) GetTopRated(ctx context.Context) (*TopRated, error) {
	query := fmt.Sprintf("%s%s%s%s%s", baseURL, movieURL, topRatedURL, api, m.apiKey)
	req, err := m.newRequest(ctx, http.MethodGet, query)
	if err != nil {
		return nil, err
	}

	res, err := m.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	top := &TopRated{}
	if err := json.NewDecoder(res.Body).Decode(top); err != nil {
		return nil, err
	}

	return top, nil
}

// GetUpcoming returns upcoming movies in theaters
func (m *movieClient) GetUpcoming(ctx context.Context) (*Upcoming, error) {
	query := fmt.Sprintf("%s%s%s%s%s", baseURL, movieURL, upcomingURL, api, m.apiKey)
	req, err := m.newRequest(ctx, http.MethodGet, query)
	if err != nil {
		return nil, err
	}

	res, err := m.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	upcoming := &Upcoming{}
	if err := json.NewDecoder(res.Body).Decode(upcoming); err != nil {
		return nil, err
	}

	return upcoming, nil
}
