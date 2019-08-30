package cloud

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostSSID(t *testing.T) {
	server := httptest.NewServer(server.handleSSID())
	PostSSID(server.URL, fakeData, "", "")
	ssid := GetPOESSID(server.URL)
	assert.Equal(t, fakeData, ssid)

}