package server

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Message chan string
	ctx context.Context
	router *http.ServeMux
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	upgrader := websocket.Upgrader{CheckOrigin: Danger}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	logrus.Info("Headers", r.Header)
	// Read messages from socket

	for {
		select {
		case m:=<-s.Message:
			logrus.Debug("Recv ", m )
			send := []byte(m)
			err:= conn.WriteMessage(1, send)
			if err != nil {
				logrus.Warn("Closing this connection, err ", err)
				return
			}
		}
	}
}

func NewServer(ctx context.Context)( s *Server) {
	server := &Server{
		Message: make(chan string),
		ctx:     ctx,
		router:  http.NewServeMux(),
	}
	return server
}

func (s *Server)Run(){
	port:=":9527"
	s.router.HandleFunc("/",s.handler)
	httpserver:=&http.Server{Addr:port, Handler:s.router}

	go func(){
		//h:=http.HandlerFunc(s.handler)
		err := httpserver.ListenAndServe()
		if err == http.ErrServerClosed{
			return
		}
		if err != nil {
			panic(err)
		}
	}()

	go func(){
		for{
			select{
			case <-s.ctx.Done():
				logrus.Info("shutdown http server")
				err := httpserver.Shutdown(s.ctx)
				if err != nil {
					logrus.Warn(err)
					return
				}
				return
			}
		}
	}()

}