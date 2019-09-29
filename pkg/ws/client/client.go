package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)


type Config struct{
	POESSID string
	League  string
	Filter  string
}


type Client struct {
	ItemStub    chan types.ItemStub
	Conn      *websocket.Conn
	Config Config
	ServerURL string
	dcChan chan struct{}
	ctx context.Context

}

func NewClient(ctx context.Context, cfg Config) *Client {
	serverURL := getServerURL(cfg.League, cfg.Filter)
	return &Client{
		ItemStub:        make(chan types.ItemStub),
		Conn:          nil,
		Config:        cfg,
		ServerURL:     serverURL,
		dcChan:        make(chan struct{}),
		ctx:           ctx,
	}
}
func (client *Client) ReadMessage() {
	go func() {
		itemID := types.ItemID{}
		for {
			_, bytes, err := client.Conn.ReadMessage()

			if e, ok :=  err.(*websocket.CloseError); ok && e.Code == websocket.ClosePolicyViolation {
				client.dcChan<-struct{}{}
				logrus.Warn("error 1008, ggg server crashed")
				return
			}
		//websocket: close 1006 (abnormal closure): unexpected EOF
			if err != nil {
				logrus.Error(errors.Wrap(err, "websocket read message error"))
				client.dcChan<-struct{}{}
				return
			}
			logrus.Debug("Receive: ", string(bytes))

			err = json.Unmarshal(bytes, &itemID)
			if err != nil {
				logrus.Error("unmarshal json message from websocket server, ", err)
				continue
			}
			stub:=types.ItemStub{
				ID:     itemID.New,
				Filter: client.Config.Filter,
			}
			client.ItemStub <- stub
		}

	}()
}

func (client *Client) Connect()(err error) {
	header:= client.getHeader()
	logrus.Infof("Connecting to %s", client.ServerURL)
	dialer:=websocket.DefaultDialer
	//dialer.HandshakeTimeout =90*time.Second
	conn, _, err := dialer.Dial(client.ServerURL, header)
	if err != nil {
		return errors.Wrapf(err, "Dial error")
	}
	logrus.Info("Connected websocket server!")
	client.Conn = conn
	client.ReadMessage()
	return nil

}


func (c *Client)MonitorStatus(){
	go func(){
		for{
			select {
				case <-c.dcChan:
					for{
						logrus.Info("reconnect in 5 sec..")

						time.Sleep(time.Second*5)
						err := c.Connect()
						if err != nil {
							logrus.Error(err)
						}else{
							break
						}

					}
				case <-c.ctx.Done():
					logrus.Println("interrupt, sending close signal to ws server")
					err := c.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					if err != nil {
						logrus.Println("write close:", err)
						return
					}
					return
			}
		}
	}()

}

func (c *Client)Run()(err error){
	err = c.Connect()
	if err != nil {
		return err
	}
	c.MonitorStatus()
	return nil
}


func getServerURL(league string, filter string) (serverUrl string) {
	urlPath := fmt.Sprintf("/api/trade/live/%s/%s", league, filter)
	u := url.URL{Scheme: "ws", Host: "www.pathofexile.com", Path: urlPath}
	return u.String()
}


func (c *Client)getHeader() (header http.Header) {
	header = getSimChromeCookie()
	logrus.Debug("using local poessid, ", os.Getenv("CLIENT_POESESSID"))
	cookie := fmt.Sprintf("POESESSID=%s", os.Getenv("CLIENT_POESESSID"))
	header.Add("Cookie", cookie)

	return header
}

func getSimChromeCookie() (header http.Header) {
	header = make(http.Header)
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "en-US,en;q=0.9,zh-TW;q=0.8,zh;q=0.7,zh-CN;q=0.6,ja;q=0.5")
	header.Add("Cache-Control", "no-cache")
	//header.Add("Connection", "Upgrade")
	header.Add("Host", "www.pathofexile.com")
	header.Add("Origin", "https://www.pathofexile.com")
	header.Add("Pragma", "no-cache")
	//header.Add("Sec-WebSocket-Extensions", "permessage-deflate; client_max_window_bits")
	//header.Add("Sec-WebSocket-Key", "Oa+B/nEJMeezec/bNsjTwg==")
	//header.Add("Sec-WebSocket-Version", "13")
	//header.Add("Upgrade", "websocket")
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	return header
}



func (client *Client) NotifyDC() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			select {
			case <-interrupt:
				logrus.Println("interrupt, sending close signal to ws server")

				err := client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					logrus.Println("write close:", err)
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
func (client *Client) Disconnect() {
	err := client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		logrus.Error("write close:", err)
	}
}
