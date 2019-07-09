package trader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrader_GetWhisper(t *testing.T) {
	client := Trader{}
	w := client.GetWhisper()
	actual := <-w
	assert.Equal(t, "yolo", actual)
}
