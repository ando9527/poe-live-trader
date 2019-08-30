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
}

func NewServer(dsn string, user string, pass string) (s *Server) {
	s = &Server{
		router: http.NewServeMux(),
		db:     nil,
		dsn:    dsn,
		user:   user,
		pass:   pass,
	}
	s.routes()

	return s
}

func (s *Server) routes() {
	s.router.HandleFunc("/", handleAuth(s.handleSSID()))
}

func (s *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s.Connect()
	logrus.Panic(http.ListenAndServe(fmt.Sprintf(":%s", port), s.router))
}

