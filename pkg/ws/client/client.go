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
	"github.com/sirupsen/logrus"
)


type Config struct{
	POESSID string
	League  string
	Filter  string
	Header http.Header
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
				logrus.Error("Websocket read message error, ", err)
				client.dcChan<-struct{}{}
				return
			}
			logrus.Debug("Receive: ", string(bytes))

			err = json.Unmarshal(bytes, &itemID)
			if err != nil {
				logrus.Error("Unmarshal json message from websocket server, ", err)
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
	header:= client.Config.Header
	logrus.Infof("Connecting to %s", client.ServerURL)
	dialer:=websocket.DefaultDialer
	//dialer.HandshakeTimeout =90*time.Second
	conn, _, err := dialer.Dial(client.ServerURL, header)
	if err != nil {
		return fmt.Errorf("dial error, %w", err)
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
					ticker:=time.NewTicker(time.Second*6)
					for _=range ticker.C{
						logrus.Info("reconnect in 6 sec..")
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
						logrus.Info("write close:", err)
						return
					}
					return
			}
		}
	}()

}

func (c *Client)Run(){
	err := c.Connect()
	if err != nil {
		logrus.Error(err)
	}
	c.MonitorStatus()
}


func getServerURL(league string, filter string) (serverUrl string) {
	urlPath := fmt.Sprintf("/api/trade/live/%s/%s", league, filter)
	u := url.URL{Scheme: "wss", Host: "www.pathofexile.com", Path: urlPath}
	return u.String()
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
