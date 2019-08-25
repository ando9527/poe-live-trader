// Package p contains an ht Cloud Function.
package p

import (
	"context"
	"fmt"
	"html"
	"net/http"

	"github.com/ando9527/poe-live-trader/cloud/env"
	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/sirupsen/logrus"
)
var conf env.Config
func init(){
	log.InitCloudLogger(false)
	conf = env.NewConfig()
}


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
		ctx := context.Background()
		client, err := cloud.NewClient(ctx)
		if err != nil {
			logrus.Fatalf("Failed to create client: %v", err)
		}
		defer client.Close()

		switch r.Method {
		case http.MethodGet:
			ssid, err := client.QuerySSID()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, err = fmt.Fprint(w, html.EscapeString(ssid))
			if err != nil {
				logrus.Error(err)
			}
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusInternalServerError)
				return
			}
			poessid := r.FormValue("poessid")
			err = client.UpdateInsert(poessid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			_, err := fmt.Fprint(w, html.EscapeString("success"))
			if err != nil {
				logrus.Error(err)
			}

		default:
			http.Error(w, "Sorry, only GET and POST methods are supported.", http.StatusInternalServerError)
		}

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



