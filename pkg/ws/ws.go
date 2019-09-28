package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/ando9527/poe-live-trader/pkg/server"
	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)


type Config struct{
	CloudEnable bool
	envPOESSID string
	CloudURL string
	User string
	Pass string
	League string
	Filter string
}


type Client struct {
	ItemID    chan []string
	Conn      *websocket.Conn
	Config Config
	ServerURL string


}

func NewClient(cfg Config) *Client {
	serverURL := getServerURL(cfg.League, cfg.Filter)
	return &Client{
		ItemID:    make(chan []string),
		Conn:      nil,
		Config:    cfg,
		ServerURL: serverURL,
	}
}
func (client *Client) ReadMessage() {
	go func() {
		itemID := types.ItemID{}
		for {
			_, bytes, err := client.Conn.ReadMessage()

			if e, ok :=  err.(*websocket.CloseError); ok && e.Code == websocket.CloseNormalClosure {
				log.Info("WS close normal")
				return
			}
			if err != nil {
				log.Error(errors.Wrap(err, "websocket read message error"))
				logrus.Info("Reconnecting in 2 seconds..")
				time.Sleep(time.Second*5)
				client.ReConnect()
				return
			}
			log.Debug("Receive: ", string(bytes))

			err = json.Unmarshal(bytes, &itemID)
			if err != nil {
				logrus.Error(err)
				continue
			}
			client.ItemID <- itemID.New
		}

	}()
}

func (client *Client) ReConnect() {
	header := client.getHeader()
	for {
		logrus.Infof("Connecting to %s", client.ServerURL)
		dialer:=websocket.DefaultDialer
		dialer.HandshakeTimeout =90*time.Second
		conn, _, err := dialer.Dial(client.ServerURL, header)
		if err == nil {
			log.Info("Connected websocket server!")
			client.Conn = conn
			client.ReadMessage()
			return
		} else {
			logrus.Panic("dial:", err)
		}
		logrus.Info("Reconnect in 5 sec..")
		time.Sleep(5 * time.Second)
		logrus.Info("Reconnecting...")
	}

}


func getServerURL(league string, filter string) (serverUrl string) {
	urlPath := fmt.Sprintf("/api/trade/live/%s/%s", league, filter)
	u := url.URL{Scheme: "wss", Host: "www.pathofexile.com", Path: urlPath}
	return u.String()
}


func (c *Client)getHeader() (header http.Header) {
	header = make(http.Header)
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "en-US,en;q=0.9,zh-TW;q=0.8,zh;q=0.7,zh-CN;q=0.6,ja;q=0.5")
	header.Add("Cache-Control", "no-cache")
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

	if c.Config.CloudEnable == false {

		logrus.Debug("using local poessid, ", os.Getenv("CLIENT_POESESSID"))
		cookie := fmt.Sprintf("POESESSID=%s", os.Getenv("CLIENT_POESESSID"))
		header.Add("Cookie", cookie)
	}else{
		s, err := server.GetPOESSID(c.Config.CloudURL, c.Config.User, c.Config.Pass)
		if err != nil {
			panic(err)
		}
		cookie := fmt.Sprintf("POESESSID=%s", s)
		header.Add("Cookie", cookie)
	}

	return header
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
