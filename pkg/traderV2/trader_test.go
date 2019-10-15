package traderV2

import (
	"testing"
	"time"

	"github.com/ando9527/poe-live-trader/pkg/item"
	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/matryer/is"
)


func TestClient_Run(t *testing.T) {
	is:=is.New(t)
	wp:=&FakeWsPool{
		ItemBuilderChan: make(chan types.ItemBuilder),
	}
	no:=&FakeNotifier{
		Receive: make(chan string),
	}
	c:=Client{
		env:        nil,
		database:   &FakeDatabase{},
		wsPool:     wp,
		idCache:    &FakeIDCache{},
		notifier:   no,
		httpClient: &FakeHttpClient{},
	}
	go func(){
		c.Run()
	}()
	b:=&item.Builder{
		ItemList: []types.Item{},
	}
	b.ItemList = append(b.ItemList, &item.Item{
		Notification: "@yolo hihi",
	})
	wp.ItemBuilderChan<-b

	for{
		select{
			case m:=<-no.Receive:
				is.Equal(m , "@yolo hihi")
				return
			case <-time.After(time.Millisecond*500):
				t.Errorf("timeout")
				return
		}
	}
}

type FakeHttpClient struct{}

func (f *FakeHttpClient) RequestItemDetail(idList []string, filterID string) (i types.ItemDetail, e error) {
	return i,nil
}

type FakeNotifier struct{
	Receive chan string
}

func (f *FakeNotifier) Run() {
	//panic("implement me")
}

func (f *FakeNotifier) SendToQueue(s string) {
	f.Receive<-s
	//panic("implement me")
}

type FakeIDCache struct{}

func (f *FakeIDCache) AllowSend(string) bool {
	return true
}

func (f *FakeIDCache) Run() {
	//panic("implement me")
}

type FakeWsPool struct{
	ItemBuilderChan chan types.ItemBuilder
}

func (f *FakeWsPool) GetBuilderChannel() <-chan types.ItemBuilder {
	return f.ItemBuilderChan
}

func (f *FakeWsPool) Run() {
	//panic("implement me")
}

type FakeDatabase struct{}

func (f  *FakeDatabase) IsIgnored(string) bool {
	return false
	//panic("implement me")
}

func (f  *FakeDatabase) Connect(name string) {
	//panic("implement me")
}

func (f  *FakeDatabase) Migration() {
	//panic("implement me")
}

