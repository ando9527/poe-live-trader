package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func FakeWebsocketServer() (server *httptest.Server) {
	upgrader := websocket.Upgrader{}
	f := func(w http.ResponseWriter, r *http.Request) {
		conn, e := upgrader.Upgrade(w, r, nil)
		if e != nil {
			panic(e)
		}
		for {
			message := `{"new":["6bf0738f765b4d364fc65105910493c13b3d89ded2797cbcca32b99ca0579825"]}`
			mt := websocket.TextMessage
			err:= conn.WriteMessage(mt, []byte(message))
			if err != nil {
				logrus.Error(err)
			}
			time.Sleep(10 * time.Millisecond)
		}

	}
	handler := http.HandlerFunc(f)
	server = httptest.NewServer(handler)
	return server
}

func FakeNewWebsocketClient(serverURL string) (client *Client) {
	newURL := "ws" + strings.TrimPrefix(serverURL, "http") + "/"
	client = &Client{
		ItemID: make(chan []string),
		ServerURL: newURL,
		ctx: context.Background(),
		}
	return client
}

func TestClient_GetItemID(t *testing.T) {
	server := FakeWebsocketServer()
	defer server.Close()

	client := FakeNewWebsocketClient(server.URL)
	expect := []string{"6bf0738f765b4d364fc65105910493c13b3d89ded2797cbcca32b99ca0579825"}
	client.Run()

	select {
	case actual := <-client.ItemID:
		logrus.Info("recv message, ", actual)
		assert.Equal(t, expect, actual)
		client.Disconnect()
		client.Conn.Close()

		time.Sleep(10 * time.Millisecond)
		return
	case <-time.After(time.Millisecond * 60):
		t.Error(errors.New("timeout"))
	}

}
