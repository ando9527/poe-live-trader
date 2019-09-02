package server

import (
	"time"

	"github.com/ando9527/poe-live-trader/pkg/graphql/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ANCHOR="ANCHOR"
var SUCCESS="SUCCESS"



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
	s.resolver.db = db
}

func (s *Server) InitTable() {
	s.resolver.db.CreateTable(&models.Ssid{})
}

func (s *Server) DropTable() {
	s.resolver.db.DropTableIfExists(&models.Ssid{})
}

func (s *Server) Migration() {
	s.resolver.db.AutoMigrate(&models.Ssid{})
}


