package ws

import (
	"testing"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestClient_GetItemID(t *testing.T) {
	server := FakeWebsocketServer()
	defer server.Close()

	client := FakeNewWebsocketClient(server.URL)
	expect := []string{"6bf0738f765b4d364fc65105910493c13b3d89ded2797cbcca32b99ca0579825"}
	client.ReConnect()
	defer client.Conn.Close()

	select {
	case actual := <-client.ItemID:
		assert.Equal(t, expect, actual)
		return
	case <-time.After(time.Millisecond * 60):
		assert.Error(t, errors.New("timeout"))
	}

}
