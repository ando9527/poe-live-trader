package cloud

import (
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
	db, e := gorm.Open("mysql", s.dsn)
	if e != nil {
		panic(e)
	}
	db.DB().SetConnMaxLifetime(time.Minute*5)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(5)
	if s.logLevel=="debug"{
		db.LogMode(true)
	}
	s.db = db
}

func (s *Server) InitTable() {
	s.db.CreateTable(&SSID{})
}

func (s *Server) DropTable() {
	s.db.DropTableIfExists(&SSID{})
}

func (s *Server) Migration() {
	s.db.AutoMigrate(&SSID{})
}

