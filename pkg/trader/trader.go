package trader

import (
	"github.com/ando9527/poe-live-trader/pkg/request"
	"github.com/ando9527/poe-live-trader/pkg/ws"
)

type Trader struct {
	Whisper         chan string
	WebsocketClient *ws.Client
	RequestClient   *request.Client
}

func NewTrader(wsConfig ws.Config) (t *Trader) {
	t = &Trader{}
	t.Whisper = make(chan string)
	t.WebsocketClient = ws.NewClient(wsConfig)
	t.RequestClient = request.NewRequestClient(wsConfig.Filter)
	return t
}
func (t *Trader) Launch() {
	t.WebsocketClient.Connect()
	t.WebsocketClient.NotifyDC()

	// get item id from websocket server
	go func() {
		for {
			select {
			case result := <-t.WebsocketClient.ItemID:
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
