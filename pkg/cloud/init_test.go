package cloud

var fakeData = "123"
var TestDsn = "root@tcp(localhost:3306)/test"
var TestUser ="user"
var TestPass ="pass"

var server *Server

func init(){
	server = NewServer(TestDsn, TestUser, TestPass)
	server.Connect()
	if server.db.HasTable(&SSID{}){
		server.DropTable()
	}
	server.InitTable()
}