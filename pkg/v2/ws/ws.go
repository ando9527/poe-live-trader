package ws

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/ando9527/poe-live-trader/conf"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	ItemID    chan []string
	Conn      *websocket.Conn
	ServerURL string
}

func (client *Client) GetItemID() (itemID chan []string) {
	return itemID
}
func (c *Client) ReConnect() {
	header := getHeader()
	for {
		logrus.Infof("Connecting to %s", c.ServerURL)
		conn, _, err := websocket.DefaultDialer.Dial(c.ServerURL, header)
		if err == nil {
			log.Info("Connected websocket server!")
			c.Conn = conn
			return
		} else {
			logrus.Fatal("dial:", err)
		}
		logrus.Info("Reconnect in 5 sec..")
		time.Sleep(5 * time.Second)
		logrus.Info("Reconnecting...")
	}
}

func NewWebsocketClient() (client *Client) {
	url := getServerURL()
	client = &Client{make(chan []string), nil, url}
	return client
}

func getServerURL() (serverUrl string) {
	urlPath := fmt.Sprintf("/api/trade/live/%s/%s", conf.Env.League, conf.Env.Filter)
	u := url.URL{Scheme: "wss", Host: "www.pathofexile.com", Path: urlPath}

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

func (c *Client) NotifyDC() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			select {
			case <-interrupt:
				log.Println("interrupt, sending close signal to ws server")

				err := c.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("write close:", err)
					return
				}
				select {
				case <-time.After(time.Second):
				}
				return
			}
		}
	}()

}
