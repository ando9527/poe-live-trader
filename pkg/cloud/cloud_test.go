package cloud

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)
var fakeSSID="123"

func TestNewClient(t *testing.T) {
	cb := context.Background()
	c ,e:= NewClient(cb)
	if e != nil {
		logrus.Panic(e)
	}

	defer c.Close()
	err := c.UpdateInsert(fakeSSID)
	if err != nil {
		panic(err)
	}
	actual, err := c.QuerySSID()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, actual, fakeSSID)

}

