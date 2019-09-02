package cloud

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostSSID(t *testing.T) {
	server := httptest.NewServer(FakeServer.handleSSID())
	PostSSID(server.URL, FakeSSID, "", "")
	ssid := GetPOESSID(server.URL, FakeUSER, FakePASS)
	assert.Equal(t, FakeSSID, ssid)

}