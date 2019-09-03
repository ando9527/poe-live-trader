// Package p contains an ht Cloud Function.
package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/ando9527/glimiter"
	"github.com/ando9527/poe-live-trader/pkg/graphql/graph"
	"github.com/didip/tollbooth"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/singleflight"
)

type Server struct {
	router *http.ServeMux
	dsn string
	user string
	pass string
	logLevel string
	production bool
	resolver *Resolver
	requestGroup *singleflight.Group
}

func NewServer(dsn string, user string, pass string, logLevel string, production bool) (s *Server) {
	s = &Server{
		router:   http.NewServeMux(),
		dsn:      dsn,
		user:     user,
		pass:     pass,
		logLevel: logLevel,
		production: production,
		resolver: &Resolver{
			db: nil ,
		},
		requestGroup: &singleflight.Group{},
	}
	s.routes()

	return s
}


func (s *Server) routes() {
	h:=handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: s.resolver}))
	h=s.handleAuth(h)
	//add limiter
	max:=1.0
	if s.production == false{
		max=20.0
	}
	lmt := tollbooth.NewLimiter(max, nil)
	//behind proxy
	lmt.SetIPLookups([]string{ "X-Forwarded-For","RemoteAddr", "X-Real-IP"})
	f:=glimiter.LimitFuncHandler(lmt, h)
	s.router.HandleFunc("/graphql", f.ServeHTTP)
}

func (s *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	s.Connect()
	//logrus.Debugf("connect to http://localhost:%s/ for GraphQL playground", port)
	logrus.Panic(http.ListenAndServe(fmt.Sprintf(":%s", port), s.router))
}

func temp(){
	//x:=func(w http.ResponseWriter, r *http.Request){
	//
	//}
	//
	//h:=http.HandlerFunc(x)
	//h.ServeHTTP()
}