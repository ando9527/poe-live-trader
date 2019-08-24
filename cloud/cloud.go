// Package p contains an ht Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/ando9527/poe-live-trader/cloud/env"
	"github.com/ando9527/poe-live-trader/pkg/log"
)
var conf env.Config
func init(){
	log.InitCloudLogger(false)
	conf = env.NewConfig()
}

//func basicAuth(username, password string) string {
//	auth := username + ":" + password
//	return base64.StdEncoding.EncodeToString([]byte(auth))
//}
//
//func redirectPolicyFunc(req *http.Request, via []*http.Request) error{
//	req.Header.Add("Authorization","Basic " + basicAuth("admin","cycloneOP"))
//	return nil
//}
var mux = newMux()


func C(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w,r)
}


func newMux() *http.ServeMux{
	mux:=http.NewServeMux()
	mux.HandleFunc("/", handleAuth(handleSSID()))

	return mux
}


func handleSSID() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var d struct {
			Message string `json:"message"`
		}
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			fmt.Fprint(w, "123")
			return
		}
		if d.Message == "" {
			fmt.Fprint(w, "123")
			return
		}
		fmt.Fprint(w, html.EscapeString(d.Message))
	}
}

func handleAuth(h http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		user , pass, _ := r.BasicAuth()
		if user!=conf.User || pass !=conf.Pass{
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		h(w,r)
	}
}



