package trader

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/ando9527/poe-live-trader/pkg/key"
	"github.com/ando9527/poe-live-trader/pkg/request"
	"github.com/ando9527/poe-live-trader/pkg/ws"
	"github.com/sirupsen/logrus"
)

type Trader struct {
	Whisper         chan string
	WebsocketPool *ws.Pool
	RequestClient   *request.Client
	KeySim *key.Client
	IDCache map[string]bool
	sync.Mutex
	ctx context.Context
}

func NewTrader(ctx context.Context, wsConfig ws.Config) (t *Trader) {
	t = &Trader{
		Whisper:       make(chan string),
		WebsocketPool: ws.NewPool(ctx, wsConfig),
		RequestClient: request.NewRequestClient(),
		KeySim:        key.NewClient(ctx),
		IDCache:       map[string]bool{},
		Mutex:         sync.Mutex{},
		ctx:           ctx,
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
			case result := <-t.WebsocketPool.ItemStubChan:
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

func (t *Trader) isPortInUsed()( ans bool){
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", "9527"), time.Second)
	if err != nil {
		logrus.Debug("testing port in used", err)
	}
	if conn != nil {
		_ = conn.Close()
		return true
	}
	return false
}

func (t *Trader) CacheClearTask(){
	go func() {
		for{
			time.Sleep(time.Minute*10)
			t.Mutex.Lock()
			t.IDCache=make(map[string]bool)
			t.Mutex.Unlock()
		}
	}()
}

func (t *Trader) Launch() {
	if t.isPortInUsed(){
		logrus.Panic("9527 port in used")
	}
	err:=t.WebsocketPool.Run()
	if err != nil {
		logrus.Panic(err)
	}
	t.processItemID()
	t.KeySim.Run()
	t.CacheClearTask()

}

