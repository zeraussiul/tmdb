package tmdb

import (
	"context"
	"errors"
	"net/http"
	"time"
)

const Version = ""

const (
	baseURL              = "https://api.themoviedb.org/3"
	api                  = "?api_key="
	searchURL            = "/search"
	movieURL             = "/movie"
	nowPlayingURL        = "/now_playing"
	upcomingURL          = "/upcoming"
	topRatedURL          = "/top_rated"
	popularURL           = "/popular"
	externalIDsURL       = "/external_ids"
	alternativeTitlesURL = "alternative_titles"
	tvURL                = "/tv"
	youtubeURL           = "https://www.youtube.com/watch?v="

	imageURL = "https://image.tmdb.org/t/p"
)

// Client is a client for TMDB with an embedded http client and api key
type Client struct {
	apiKey string
	http   *http.Client
}

// Option used to override default client options
type Option func(*Client)

// NewClient creates a new TMDB Client with a default timeout duration
func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("not a valid API key")
	}
	c := &Client{
		apiKey: apiKey,
		http:   http.DefaultClient,
	}
	// default timeout of 10 seconds
	c.http.Timeout = time.Minute

	// range through options and apply them.
	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

func WithClient(client *http.Client) Option {
	return func(c *Client) {
		c.http = client
	}
}

// SetTimeout changes the timeout duration of the TMDB Client
func (c *Client) SetTimeout(duration time.Duration) {
	c.http.Timeout = duration
}

// SetApiKey changes the api key on the client
func (c *Client) SetApiKey(apiKey string) error {
	if len(apiKey) == 0 {
		return errors.New("not a valid API key")
	}
	c.apiKey = apiKey
	return nil
}

func (c *Client) newRequest(ctx context.Context, method string, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
