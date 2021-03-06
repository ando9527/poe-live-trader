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

	"github.com/ando9527/poe-live-trader/pkg/item"
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
	ItemBuilderChan chan types.ItemBuilder
	Conn            *websocket.Conn
	Config          Config
	ServerURL       string
	ctx             context.Context

}

func NewClient(ctx context.Context, cfg Config) *Client {
	serverURL := getServerURL(cfg.League, cfg.Filter)
	return &Client{
		ItemBuilderChan: make(chan types.ItemBuilder),
		Conn:            nil,
		Config:          cfg,
		ServerURL:       serverURL,
		ctx:             ctx,
	}
}
func (client *Client) ReadMessage()(err error) {
		itemID := types.ItemID{}
		for {
			_, bytes, err := client.Conn.ReadMessage()
			if e, ok :=  err.(*websocket.CloseError); ok && e.Code == websocket.ClosePolicyViolation {
				return fmt.Errorf("error 1008, ggg server crashed")
			}

			if err != nil {
				return	fmt.Errorf("websocket read message error, %w", err)
			}

			logrus.Debug("Receive: ", string(bytes))

			err = json.Unmarshal(bytes, &itemID)
			if err != nil {
				logrus.Error("unmarshal json message from websocket server, ", err)
				continue
			}
			
			b:=item.Builder{
				IdList:   itemID.New,
				FilterID: client.Config.Filter,
			}
			client.ItemBuilderChan <- &b
		}
}

func (client *Client) Connect() {
	header:= client.Config.Header
	logrus.Infof("Connecting to %s", client.ServerURL)
	dialer:=websocket.DefaultDialer
	//dialer.HandshakeTimeout =90*time.Second
	var err error
	for{
		client.Conn, _, err = dialer.Dial(client.ServerURL, header)
		if err != nil {
			logrus.Warn("Dial error ", err)
			logrus.Info("Reconnect in 6 sec..")
			time.Sleep(time.Second*6)
			continue
		}
		break
	}
	logrus.Info("Connected websocket server!")
}

func (c *Client)Run(){
	go func(){
		//Do  the reconnect
		for{
			c.Connect()
			err := c.ReadMessage()
			if err != nil {
				logrus.Error(err)
			}
		}
	}()
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
