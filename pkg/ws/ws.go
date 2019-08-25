package ws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/ando9527/poe-live-trader/cmd/client/conf"
	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	ItemID    chan []string
	Conn      *websocket.Conn
	ServerURL string
}

func (client *Client) ReadMessage() {
	go func() {
		itemID := types.ItemID{}
		for {
			_, bytes, err := client.Conn.ReadMessage()
			if err != nil {
				log.Error(errors.Wrap(err, "websocket read message error"))
				return
			}
			log.Debug("Receive: ", string(bytes))

			err = json.Unmarshal(bytes, &itemID)
			if err != nil {
				panic(err)
			}
			client.ItemID <- itemID.New
		}

	}()
}

func (client *Client) ReConnect() {
	header := getHeader()
	for {
		logrus.Infof("Connecting to %s", client.ServerURL)
		conn, _, err := websocket.DefaultDialer.Dial(client.ServerURL, header)
		if err == nil {
			log.Info("Connected websocket server!")
			client.Conn = conn
			client.ReadMessage()
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
	serverURL := getServerURL()
	client = &Client{make(chan []string), nil, serverURL}
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

	cookie := fmt.Sprintf("POESESSID=%s", getPOESSID())
	header.Add("Cookie", cookie)

	return header
}

func getPOESSID() (ssid string) {
	if conf.Env.CloudEnable==false{
		logrus.Debug("using local poessid")
		return conf.Env.Poesessid
	}
	c:=http.Client{Timeout:time.Second*10}
	req, e := http.NewRequest("GET", conf.Env.CloudUrl, nil)
	if e != nil {
		logrus.Fatal(e)
	}
	req.SetBasicAuth(conf.Env.User,conf.Env.Pass)
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	all, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		logrus.Fatal(e)
	}
	logrus.Debug("using cloud poessid")
	return string(all)
}

func (client *Client) NotifyDC() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			select {
			case <-interrupt:
				log.Println("interrupt, sending close signal to ws server")

				err := client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
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
