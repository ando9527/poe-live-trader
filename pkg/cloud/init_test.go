package cloud

var FakeSSID = "123"
var FakeDSN = "root@tcp(localhost:3306)/test"
var FakeUSER ="user"
var FakePASS ="pass"
var FakeLogLevel = "warn"

var FakeServer *Server

func init(){
	FakeServer = NewServer(FakeDSN, FakeUSER, FakePASS, FakeLogLevel)
	FakeServer.Connect()
	if FakeServer.db.HasTable(&SSID{}){
		FakeServer.DropTable()
	}
	FakeServer.InitTable()
}