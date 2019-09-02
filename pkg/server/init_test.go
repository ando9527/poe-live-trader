package server

import (
	"github.com/ando9527/poe-live-trader/pkg/graphql/models"
)

var fakeData = "123"
var TestDsn = "root@tcp(localhost:3306)/test"
var TestUser ="user"
var TestPass ="pass"
var TestLogLevel= "warn"

var server *Server

func init(){
	server = NewServer(TestDsn, TestUser, TestPass, TestLogLevel)
	server.Connect()

	if server.resolver.db.HasTable(&models.Ssid{}){
		server.DropTable()
	}
	server.InitTable()

}