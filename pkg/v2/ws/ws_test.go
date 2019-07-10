package ws

import (
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetItemID(t *testing.T) {
	server := FakeWebsocketServer()
	defer server.Close()

	client := FakeNewWebsocketClient(server.URL)
	expect := []string{"6bf0738f765b4d364fc65105910493c13b3d89ded2797cbcca32b99ca0579825"}
	defer client.Conn.Close()

	select {
	case actual := <-client.GetItemID():
		assert.Equal(t, expect, actual)
	case <-time.After(time.Millisecond * 60):
		assert.Error(t, errors.New("timeout"))
	}

}
