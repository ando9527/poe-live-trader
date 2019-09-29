package trader

import (
	"context"
	"sync"

	"github.com/ando9527/poe-live-trader/pkg/request"
	"github.com/ando9527/poe-live-trader/pkg/ws"
	"github.com/ando9527/poe-live-trader/pkg/ws/server"
	"github.com/sirupsen/logrus"
)

type Trader struct {
	Whisper         chan string
	WebsocketClient *ws.Client
	RequestClient   *request.Client
	LocalServer *server.Server
	IDCache map[string]bool
	sync.Mutex
	ctx context.Context
}

func NewTrader(ctx context.Context, wsConfig ws.Config) (t *Trader) {
	t = &Trader{
		Whisper:         make(chan string),
		WebsocketClient: ws.NewClient( ctx,wsConfig),
		RequestClient:   request.NewRequestClient(wsConfig.Filter),
		LocalServer:     server.NewServer(),
		IDCache:         map[string]bool{},
		ctx:             ctx,
	}
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

func (t *Trader) processItemID() {
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

func (t *Trader) Launch() {
	err:=t.WebsocketClient.Run()
	if err != nil {
		logrus.Panic(err)
	}
	t.processItemID()
	//t.WebsocketClient.NotifyDC()
	t.LocalServer.Run()

	// get item id from websocket server


}
