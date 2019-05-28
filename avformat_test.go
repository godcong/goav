package goav_test

import (
	"testing"

	"github.com/godcong/goav"
	"github.com/stretchr/testify/assert"
)

// TestAvUrlSplitWithSufficientSizes ...
func TestAvUrlSplitWithSufficientSizes(t *testing.T) {
	url := "rtsp://user:password@example.com:554/path"
	var port int
	proto, authorization, hostname, path := goav.AVURLSplit(100, 100, 100, &port, 100, url)

	assert.Equal(t, "rtsp", proto)
	assert.Equal(t, "user:password", authorization)
	assert.Equal(t, "example.com", hostname)
	assert.Equal(t, 554, port)
	assert.Equal(t, "/path", path)
}

// TestAvUrlSplitWithInsufficientSizes ...
func TestAvUrlSplitWithInsufficientSizes(t *testing.T) {
	url := "https://user:password@example.com:443/here/is/the/path"
	proto, authorization, hostname, path := goav.AVURLSplit(3, 5, 5, nil, 5, url)

	assert.Equal(t, "ht", proto)
	assert.Equal(t, "user", authorization)
	assert.Equal(t, "exam", hostname)
	assert.Equal(t, "/her", path)
}
