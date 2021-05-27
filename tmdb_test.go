package tmdb

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	// init bad api key client
	apiKey := ""
	c, err := NewClient(apiKey)
	require.Error(t, err)
	require.Empty(t, c)
	// create valid client
	_ = createNewClient(t)

}

func TestClient_SetTimeout(t *testing.T) {
	c := createNewClient(t)
	to := time.Second * 30

	c.SetTimeout(to)
	require.Equal(t, to, c.http.Timeout)

}

func TestClient_SetApiKey(t *testing.T) {
	c := createNewClient(t)
	apiKey := ""
	err := c.SetApiKey(apiKey)
	require.Error(t, err)
	apiKey = "11223344"

	err = c.SetApiKey(apiKey)
	require.NoError(t, err)
	require.Equal(t, c.apiKey, apiKey)
}

func createNewClient(t *testing.T) *Client {
	apiKey := "1234567890"
	c, err := NewClient(apiKey)
	require.NoError(t, err)
	require.NotEmpty(t, c)
	require.NotEmpty(t, c.http)

	require.Equal(t, c.apiKey, apiKey)
	return c
}

// TODO: rewrite test
func TestWithClient(t *testing.T) {
	c := &http.Client{}
	c2, err := NewClient("123456", WithClient(c))
	require.NoError(t, err)
	require.NotEmpty(t, c2)
	require.Equal(t, c, c2.http)
}
