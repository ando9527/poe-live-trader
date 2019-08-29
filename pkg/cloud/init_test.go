package cloud

import (
	"github.com/ando9527/poe-live-trader/pkg/log"
)

var fakeData = "123"

var server *Server

func init(){
	log.InitLogger(true)
	server = NewServer()
	server.Connect()
	if server.db.HasTable(&SSID{}){
		server.DropTable()
	}
	server.InitTable()
}