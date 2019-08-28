package cloud

import (
	"testing"
)


func init()  {
	server = NewServer()
	server.Connect()
}


func TestInitTable(t *testing.T) {
	server.InitTable()
	server.DropTable()
}

