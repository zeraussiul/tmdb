package tmdb

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchClient(t *testing.T) {
	createSearchClient(t)
}

func TestSearchClient_Movies(t *testing.T) {
	s := createSearchClient(t)

	res, err := s.Movies("")
	require.Error(t, err)
	require.Empty(t, res)

	_, err = s.Movies("fight club")
	require.NoError(t, err)
	//require.NotEmpty(t, movs)

}

func TestSearchClient_TV(t *testing.T) {
	s := createSearchClient(t)

	res, err := s.TV("")
	require.Error(t, err)
	require.Empty(t, res)

	_, err = s.TV("game of thrones")
	require.NoError(t, err)
	//require.NotEmpty(t, shws)
}

func createSearchClient(t *testing.T) *searchClient {
	c := createNewClient(t)
	s := SearchClient(c)
	require.NotEmpty(t, s)

	return s
}
