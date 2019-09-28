package trader

import (
	"sync"

	"github.com/ando9527/poe-live-trader/pkg/request"
	"github.com/ando9527/poe-live-trader/pkg/ws"
	"github.com/ando9527/poe-live-trader/pkg/ws/server"
)

type Trader struct {
	Whisper         chan string
	WebsocketClient *ws.Client
	RequestClient   *request.Client
	LocalServer *server.Server
	IDCache map[string]bool
	sync.Mutex
}

func NewTrader(wsConfig ws.Config) (t *Trader) {
	t = &Trader{}
	t.Whisper = make(chan string)
	t.WebsocketClient = ws.NewClient( wsConfig)
	t.RequestClient = request.NewRequestClient(wsConfig.Filter)
	t.LocalServer = server.NewServer()
	t.IDCache = map[string]bool{}

	//go func(){
	//	for{
	//		time.After(time.Minute*30)
	//		t.Mutex.Lock()
	//		t.IDCache = map[string]bool{}
	//		t.Mutex.Unlock()
	//	}
	//}()

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
