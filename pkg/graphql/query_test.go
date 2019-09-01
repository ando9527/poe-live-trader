package graphql

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	ssid:=&Ssid{}
	e:=server.resolver.db.Where(Ssid{Anchor: ANCHOR}).First(&ssid).Error
	if e != nil {
		t.Error(e)
	}
	fmt.Println(ssid)
}

