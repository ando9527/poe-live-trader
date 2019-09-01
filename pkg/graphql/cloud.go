// Package p contains an ht Cloud Function.
package graphql

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router *http.ServeMux
	dsn string
	user string
	pass string
	logLevel string
	resolver *Resolver
}

func NewServer(dsn string, user string, pass string, logLevel string) (s *Server) {
	s = &Server{
		router:   http.NewServeMux(),
		dsn:      dsn,
		user:     user,
		pass:     pass,
		logLevel: logLevel,
		resolver: &Resolver{
			db: nil,
		},
	}
	s.routes()

	return s
}

func (s *Server) routes() {

	p:= handler.Playground("GraphQL playground", "/query")
	s.router.HandleFunc("/", p)


	h:=handler.GraphQL(NewExecutableSchema(Config{Resolvers: s.resolver}))
	s.router.HandleFunc("/query", h)
}

func (s *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s.Connect()
	logrus.Debugf("connect to http://localhost:%s/ for GraphQL playground", port)
	logrus.Panic(http.ListenAndServe(fmt.Sprintf(":%s", port), s.router))
}

