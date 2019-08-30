// Package p contains an ht Cloud Function.
package cloud

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)









type Server struct {
	router *http.ServeMux
	db *gorm.DB
	dsn string
	user string
	pass string
	logLevel string
}

func NewServer(dsn string, user string, pass string, logLevel string) (s *Server) {
	s = &Server{
		router: http.NewServeMux(),
		db:     nil,
		dsn:    dsn,
		user:   user,
		pass:   pass,
		logLevel: logLevel,
	}
	s.routes()

	return s
}

func (s *Server) routes() {
	s.router.HandleFunc("/", s.handleAuth(s.handleSSID()))
}

func (s *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s.Connect()
	logrus.Panic(http.ListenAndServe(fmt.Sprintf(":%s", port), s.router))
}

