package trader

import (
	"fmt"
	"net/http"

	"github.com/ando9527/poe-live-trader/conf"
	"github.com/ando9527/poe-live-trader/pkg/v2/types"
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
	WebsocketClient *WebsocketClient
	RequestClient   *RequestClient
}

func NewTrader() (t *Trader) {
	t = &Trader{}
	t.Whisper = make(chan string)
	t.WebsocketClient = NewWebsocketClient()
	t.RequestClient = NewRequestClient()

	// get item id from websocket server
	itemID := t.WebsocketClient.GetItemID()
	// get detail of item from http server
	itemDetail := t.RequestClient.RequestItemDetail(itemID)
	go func() {
		for _, result := range itemDetail.Result {
			t.Whisper <- result.Listing.Whisper
		}
	}()
	return t
}

func (c *Trader) GetWhisper() (whisper chan string) {
	return c.Whisper
}

type RequestClient struct {
}

func (client *RequestClient) RequestItemDetail(strings []string) (itemDetail types.ItemDetail) {
	return itemDetail
}

func NewRequestClient() (client *RequestClient) {
	return &RequestClient{}
}

type WebsocketClient struct {
}

func (client *WebsocketClient) GetItemID() []string {
	return nil
}

func NewWebsocketClient() (client *WebsocketClient) {
	return &WebsocketClient{}
}
