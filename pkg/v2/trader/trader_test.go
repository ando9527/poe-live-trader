package trader

import (
	"testing"
	"time"

	"github.com/ando9527/poe-live-trader/pkg/v2/ws"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestTrader_GetWhisper(t *testing.T) {
	client := Trader{}
	fakeWSServer := ws.FakeWebsocketServer()
	fakeWSClient := ws.FakeNewWebsocketClient(fakeWSServer.URL)
	client.WebsocketClient = fakeWSClient

	w := client.GetWhisper()
	expect := "@Taranis__R_n_B___Legion Hi, I would like to buy your Exalted Orb listed for 166 chaos in Legion (stash tab \"GG\"; position: left 12, top 1)"
	duration := time.Duration(20 * time.Millisecond)
	select {
	case <-time.After(duration):
		assert.Error(t, errors.New("timeout"))
	case actual := <-w:
		assert.Equal(t, expect, actual)
	}

}
