package graphql

var fakeData = "123"
var TestDsn = "root@tcp(localhost:3306)/test"
var TestUser ="user"
var TestPass ="pass"
var TestLogLevel= "warn"

var server *Server

func init(){
	server = NewServer(TestDsn, TestUser, TestPass, TestLogLevel)
	server.Connect()

	if server.resolver.db.HasTable(&Ssid{}){
		server.DropTable()
	}
	server.InitTable()

}