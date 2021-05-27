package tmdb

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

// TODO: Write proper tests

func createMovieClient(t *testing.T) *movieClient {
	c := createNewClient(t)
	m := MovieClient(c)
	require.NotEmpty(t, m)
	return m
}

func TestMovieClient(t *testing.T) {
	createMovieClient(t)
}

func TestMovieClient_GetDetails(t *testing.T) {
	m := createMovieClient(t)
	_, err := m.GetDetails(context.Background(), 123)
	require.NoError(t, err)
}

func TestMovieClient_GetAlternativeTitles(t *testing.T) {
	m := createMovieClient(t)
	_, err := m.GetAlternativeTitles(context.Background(), 123)
	require.NoError(t, err)
}

func TestMovieClient_GetExternalIDs(t *testing.T) {
	m := createMovieClient(t)
	_, err := m.GetExternalIDs(context.Background(), 123)
	require.NoError(t, err)
}

func TestMovieClient_GetVideos(t *testing.T) {
	m := createMovieClient(t)

	_, err := m.GetVideos(context.Background(), 567189)
	require.NoError(t, err)
}

func TestMovieClient_GetNowPlaying(t *testing.T) {
	m := createMovieClient(t)
	_, err := m.GetNowPlaying(context.Background())
	require.NoError(t, err)
}

func TestMovieClient_GetPopular(t *testing.T) {
	m := createMovieClient(t)
	_, err := m.GetPopular(context.Background())
	require.NoError(t, err)
}

func TestMovieClient_GetTopRated(t *testing.T) {
	m := createMovieClient(t)
	_, err := m.GetTopRated(context.Background())
	require.NoError(t, err)
}

func TestMovieClient_GetUpcoming(t *testing.T) {
	m := createMovieClient(t)
	_, err := m.GetUpcoming(context.Background())
	require.NoError(t, err)
}
