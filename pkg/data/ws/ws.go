package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/atotto/clipboard"

	"github.com/ando9527/poe-live-trader/pkg/audio"
	http2 "github.com/ando9527/poe-live-trader/pkg/data/http"

	"github.com/sirupsen/logrus"

	"github.com/ando9527/poe-live-trader/pkg/conf"
	"github.com/gorilla/websocket"
)

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

type LiveData struct {
	New []string `json:"new"`
}
type ItemHandler func(itemID []string) (itemDetail http2.ItemDetail)

func reconnect() (conn *websocket.Conn) {
	urlPath := fmt.Sprintf("/api/trade/live/%s/%s", conf.Env.League, conf.Env.Filter)

	u := url.URL{Scheme: "wss", Host: "www.pathofexile.com", Path: urlPath}
	logrus.Infof("connecting to %s", u.String())

	header := getHeader()
	for {
		c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
		if err == nil {
			return c
		} else {
			logrus.Fatal("dial:", err)
		}
		logrus.Info("Reconnect in 5 sec..")
		time.Sleep(5 * time.Second)
		logrus.Info("Reconnecting...")
	}

}
func Connect(itemHandler ItemHandler) {

	ctx, cancel := context.WithCancel(context.Background())

	// Connect to Server
	conn := reconnect()
	logrus.Info("websocket server connected, start receiving message.. ")
	defer conn.Close()

	// Read Message
	go func(ctx2 context.Context) {
		defer cancel()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				logrus.Error("websocket read message error: ", err)
				return
			}
			logrus.Debugf("recv: %s", message)
			liveData := LiveData{}
			err = json.Unmarshal([]byte(message), &liveData)
			if err != nil {
				logrus.Fatalf("decode json message from ws server failed, message: %s", message)
			}
			go func() {
				itemDetail := itemHandler(liveData.New)
				for _, result := range itemDetail.Result {
					fmt.Println(result.Listing.Whisper)
					err := clipboard.WriteAll(result.Listing.Whisper)
					if err != nil {
						logrus.Warn("failed copy whisper to clipboard.")
					}
					audio.Play()
				}
			}()

		}
	}(ctx)

	// Cleanly close the connection by sending a close message and then
	// waiting (with timeout) for the server to close the connection.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			log.Println("interrupt, sending close signal to ws server")

			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			cancel()
			select {
			case <-time.After(time.Second):
			}
			return
		}
	}
}
