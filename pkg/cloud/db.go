package cloud

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ANCHOR="ANCHOR"
var SUCCESS="SUCCESS"

type SSID struct {
	ID int
	Content string
	Anchor string
}


func (s *Server) Connect() {
	db, e := gorm.Open("mysql", os.Getenv("APP_DSN"))
	if e != nil {
		panic(e)
	}
	db.DB().SetConnMaxLifetime(time.Minute*5);
	db.DB().SetMaxIdleConns(5);
	db.DB().SetMaxOpenConns(5);
	s.db = db
}

func (s *Server) InitTable() {
	s.db.CreateTable(&SSID{})
}

func (s *Server) DropTable() {
	s.db.DropTableIfExists(&SSID{})
}


