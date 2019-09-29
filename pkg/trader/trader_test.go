package trader

//import (
//	"testing"
//	"time"
//
//	"github.com/ando9527/poe-live-trader/pkg/request"
//	"github.com/ando9527/poe-live-trader/pkg/ws"
//	"github.com/pkg/errors"
//	"github.com/stretchr/testify/assert"
//)

//func TestTrader_GetWhisper(t *testing.T) {
//	// ws server
//	fakeWSServer := ws.FakeWebsocketServer()
//	defer fakeWSServer.Close()
//
//	// ws client
//	fakeWSClient := ws.FakeNewWebsocketClient(fakeWSServer.URL)
//
//	// http server
//	fakeRequestServer := request.NewFakeRequestServer()
//	defer fakeRequestServer.Close()
//
//	// http client
//	fakeRequestClient := request.NewFakeRequestClient(fakeRequestServer.URL)
//
//	client := &Trader{}
//	client.RequestClient = fakeRequestClient
//	client.WebsocketClient = fakeWSClient
//	client.Whisper = make(chan string)
//
//	// start testing
//	client.Launch()
//	expect := "@Taranis__R_n_B___Legion Hi, I would like to buy your Exalted Orb listed for 166 chaos in Legion (stash tab \"GG\"; position: left 12, top 1)"
//	duration := time.Duration(20 * time.Millisecond)
//	for {
//		select {
//		case <-time.After(duration):
//			t.Error(errors.New("timeout"))
//		case actual := <-client.Whisper:
//			assert.Equal(t, expect, actual)
//			return
//		}
//	}
//
//}
