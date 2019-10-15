package traderV2

import (
	"testing"

	"github.com/ando9527/poe-live-trader/pkg/types"
)

func TestClient_Run(t *testing.T) {
	//c:=Client{
	//	env:        nil,
	//	database:   &FakeDatabase{},
	//	wsPool:     &FakeWsPool{},
	//	idCache:    &FakeIDCache{},
	//	notifier:   &FakeNotifier{},
	//	httpClient: &FakeHttpClient{},
	//}
	//c.Run()
}

type FakeHttpClient struct{}

func (f *FakeHttpClient) RequestItemDetail(idList []string, filterID string) (types.ItemDetail, error) {
	panic("implement me")
}

type FakeNotifier struct{}

func (f *FakeNotifier) Run() {
	panic("implement me")
}

func (f *FakeNotifier) SendToQueue(string) {
	panic("implement me")
}

type FakeIDCache struct{}

func (f *FakeIDCache) AllowSend(string) bool {
	panic("implement me")
}

func (f *FakeIDCache) Run() {
	panic("implement me")
}

type FakeWsPool struct{}

func (f *FakeWsPool) GetBuilderChannel() <-chan types.ItemBuilder {
	panic("implement me")
}

func (f *FakeWsPool) Run() {
	panic("implement me")
}

type FakeDatabase struct{}

func (f  *FakeDatabase) isIgnored(string) bool {
	panic("implement me")
}

func (f  *FakeDatabase) Connect() {
	panic("implement me")
}

func (f  *FakeDatabase) Migration() {
	panic("implement me")
}

