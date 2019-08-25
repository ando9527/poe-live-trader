package ssid

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSSID(t *testing.T) {
	c:=NewClient()
	ssid, err := c.GetSSID()
	if err != nil {
		logrus.Panic(err)
	}
	assert.Equal(t, ssid, "123")
}

