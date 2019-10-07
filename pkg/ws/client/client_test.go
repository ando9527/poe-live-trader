package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ando9527/poe-live-trader/pkg/types"
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
			time.Sleep(time.Second)
		}
	}
	handler := http.HandlerFunc(f)
	server = httptest.NewServer(handler)
	return server
}

func FakeNewWebsocketClient(ctx context.Context, serverURL string) (client *Client) {
	newURL := "ws" + strings.TrimPrefix(serverURL, "http") + "/"
	client = &Client{
		ItemStub: make(chan types.ItemStub),
		ServerURL: newURL,
		ctx: ctx,
		}
	return client
}

func TestClient_GetItemID(t *testing.T) {
	server := FakeWebsocketServer()
	defer server.Close()
	ctx, cancel := context.WithCancel(context.Background())
	client := FakeNewWebsocketClient(ctx,server.URL)
	expect := types.ItemStub{
		ID:     []string{"6bf0738f765b4d364fc65105910493c13b3d89ded2797cbcca32b99ca0579825"},
		Filter: "",
	}

	go func(){
		select {
		case actual := <-client.ItemStub:
			logrus.Info("recv message, ", actual)
			assert.Equal(t, expect, actual)
			client.Disconnect()
			client.Conn.Close()
			cancel()

			//time.Sleep(10 * time.Millisecond)
			return
		case <-time.After(time.Millisecond * 60):
			t.Error(errors.New("timeout"))
		}
	}()
	err := client.Run()
	if err != nil {
		t.Fatal(err)
	}

}
