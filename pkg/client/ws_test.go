package client

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/websocket"
)

type MockServer struct {
	upgrader websocket.Upgrader
	Want     string
}

func (m *MockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	// Do the echo
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		m.Want = string(message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func TestClient_ReConnect(t *testing.T) {
	// mock server initialize
	mockServer := &MockServer{websocket.Upgrader{}, ""}
	server := httptest.NewServer(mockServer)
	serverURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/"
	defer server.Close()

	// Testing client
	client := Client{ServerURL: serverURL}
	client.ReConnect()
	defer client.Conn.Close()
	want := "yolo"
	if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(want)); err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, want, mockServer.Want)
}

func TestReadMessage(t *testing.T) {
	// mock server initialize
	mockServer := &MockServer{websocket.Upgrader{}, ""}
	server := httptest.NewServer(mockServer)
	serverURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/"
	defer server.Close()

	// Testing client
	client := Client{ServerURL: serverURL}
	client.ReConnect()
	defer client.Conn.Close()
	want := "yolo"
	if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(want)); err != nil {
		panic(err)
	}

	receive := bytes.Buffer{}
	client.ReadMessage(&receive)
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, want, receive.String())
}
