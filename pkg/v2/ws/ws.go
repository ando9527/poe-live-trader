package ws

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ando9527/poe-live-trader/conf"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Client struct {
	ItemID chan []string
	Conn   *websocket.Conn
}

func (client *Client) GetItemID() (itemID chan []string) {
	return itemID
}

func NewWebsocketClient() (client *Client) {
	serverURL := getServerURL()
	header := getHeader()
	conn, _, e := websocket.DefaultDialer.Dial(serverURL, header)
	if e != nil {
		panic(e)
	}
	client = &Client{make(chan []string), conn}
	return client
}

func getServerURL() (serverUrl string) {
	urlPath := fmt.Sprintf("/api/trade/live/%s/%s", conf.Env.League, conf.Env.Filter)
	u := url.URL{Scheme: "wss", Host: "www.pathofexile.com", Path: urlPath}
	logrus.Infof("connecting to %s", u.String())
	return u.String()
}

func getHeader() (header http.Header) {
	header = make(http.Header)
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "en-US,en;q=0.9,zh-TW;q=0.8,zh;q=0.7,zh-CN;q=0.6,ja;q=0.5")
	header.Add("Cache-Control", "no-cache")
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

	cookie := fmt.Sprintf("POESESSID=%s", conf.Env.Poesessid)
	header.Add("Cookie", cookie)

	return header
}
