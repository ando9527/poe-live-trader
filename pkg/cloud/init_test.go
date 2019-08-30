package cloud

var fakeData = "123"

var server *Server

func init(){
	server = NewServer()
	server.Connect()
	if server.db.HasTable(&SSID{}){
		server.DropTable()
	}
	server.InitTable()
}