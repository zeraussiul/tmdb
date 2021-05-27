package tmdb

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTvClient(t *testing.T) {
	c := createNewClient(t)
	tv := TvClient(c)
	require.NotEmpty(t, tv)
}

// TODO: Properly write this test
func TestTvClient_GetDetails(t *testing.T) {
	c := createNewClient(t)
	tv := TvClient(c)

	_, err := tv.GetDetails(12)
	require.NoError(t, err)
	//require.NotEmpty(t, details)

}
