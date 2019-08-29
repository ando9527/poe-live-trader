package cloud


func init()  {
	server = NewServer()
	server.Connect()
}

//
//func TestInitTable(t *testing.T) {
//	if server.db.HasTable(&SSID{}){
//		server.DropTable()
//	}
//	server.InitTable()
//}

