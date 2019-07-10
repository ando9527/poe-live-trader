package ws

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/gorilla/websocket"
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
			e := conn.WriteMessage(mt, []byte(message))
			if e != nil {
				panic(e)
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

	conn, _, e := websocket.DefaultDialer.Dial(newURL, nil)
	if e != nil {
		panic(e)
	}
	client = &Client{make(chan []string), conn}
	return client
}
