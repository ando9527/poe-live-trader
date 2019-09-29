package server

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Message chan string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func NewServer()(s *Server) {
	server := &Server{
		Message: make(chan string),
	}
	return server
}

func (s *Server)Run(){
	go func(){
		port:=":9527"
		logrus.Debug("Local server listening port ", port)
		if err := http.ListenAndServe(port , s); err != nil {
			logrus.Panic(err)
		}
	}()

}