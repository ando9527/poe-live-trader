package trader

import (
	"fmt"
	"net/http"

	"github.com/ando9527/poe-live-trader/conf"
	"github.com/ando9527/poe-live-trader/pkg/v2/request"
	"github.com/ando9527/poe-live-trader/pkg/v2/ws"
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

type Trader struct {
	Whisper         chan string
	WebsocketClient *ws.Client
	RequestClient   *request.Client
}

func NewTrader() (t *Trader) {
	t = &Trader{}
	t.Whisper = make(chan string)
	t.WebsocketClient = ws.NewWebsocketClient()
	t.RequestClient = request.NewRequestClient()
	return t
}
func (t *Trader) Launch() {
	t.WebsocketClient.ReConnect()
	t.WebsocketClient.NotifyDC()
	defer t.WebsocketClient.Conn.Close()

	// get item id from websocket server
	itemID := t.WebsocketClient.GetItemID()
	go func() {
		for {
			select {
			case result := <-itemID:
				// get detail of item from http server
				go func() {
					itemDetail := t.RequestClient.RequestItemDetail(result)
					for _, result := range itemDetail.Result {
						t.Whisper <- result.Listing.Whisper
					}
				}()
			}
		}
	}()

}
func (c *Trader) GetWhisper() (whisper chan string) {
	return c.Whisper
}
