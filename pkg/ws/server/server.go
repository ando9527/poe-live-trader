package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Server struct {
	MessageChan  chan string
	ClientNumber int
	address string
	connList []*websocket.Conn
	sync.Mutex
}

func NewServer(address string) *Server {
	return &Server{
		MessageChan:  make(chan string),
		ClientNumber: 0,
		address:      address,
		connList:     []*websocket.Conn{},
		Mutex:        sync.Mutex{},
	}
}


func (s *Server) SendToServer(m string) {
	s.MessageChan <-m
}



// Danger Valid nothing
func Danger(r *http.Request) bool {
	return true
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	upgrader := websocket.Upgrader{CheckOrigin: Danger}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	s.ClientNumber +=1
	logrus.Debug("ClientNumber number: ", s.ClientNumber)
	logrus.Info("Headers", r.Header)
	s.Mutex.Lock()
	s.connList = append(s.connList, conn)
	s.Mutex.Unlock()
	// Read messages from socket
	//go func(){
	//	for {
	//		_, msg, err := conn.ReadMessage()
	//		if err != nil {
	//			_ = conn.Close()
	//			logrus.Error("Recv error ", err)
	//			return
	//		}
	//		logrus.Info("Recv ", string(msg))
	//
	//	}
	//}()


}

func (s *Server) Run() {
	go func(){
		if err := http.ListenAndServe(s.address, s); err != nil {
			logrus.Fatal(err)
		}
		logrus.Info("server")
	}()

	go func(){
		for m:=range s.MessageChan{
			for index,conn:=range s.connList{
				send := []byte(m)
				err := conn.WriteMessage(websocket.TextMessage, send)
				if err != nil {
					logrus.Error("Sent error ",err)
					s.Mutex.Lock()
					s.connList= remove(s.connList, index)
					s.Mutex.Unlock()
				}
			}

		}
	}()
}
func remove(s []*websocket.Conn, i int) []*websocket.Conn {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}