// Package p contains an ht Cloud Function.
package cloud

import (

	"net/http"

	"github.com/ando9527/poe-live-trader/cloud/env"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/jinzhu/gorm"
)
var conf env.Config
func init(){
	log.InitCloudLogger(false)
	conf = env.NewConfig()
}








type Server struct {
	router *http.ServeMux
	db *gorm.DB
}

func NewServer()(s *Server){
	s = &Server{
		router: http.NewServeMux(),
		db:     nil,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.HandleFunc("/", handleAuth(s.handleSSID()))
}

