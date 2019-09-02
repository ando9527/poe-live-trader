package server

import (
	"github.com/ando9527/poe-live-trader/pkg/graphql/models"
)

var FakeData = "123"
var FakeDSN = "root@tcp(localhost:3306)/test"
var FakeUser ="user"
var FakePass ="pass"
var FakeLogLevel = "warn"

var FakeServer *Server

func init(){
	FakeServer = NewServer(FakeDSN, FakeUser, FakePass, FakeLogLevel)
	FakeServer.Connect()

	if FakeServer.resolver.db.HasTable(&models.Ssid{}){
		FakeServer.DropTable()
	}
	FakeServer.InitTable()

}