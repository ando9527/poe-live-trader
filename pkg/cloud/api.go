package cloud

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *Server)handleSSID() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		//case http.MethodGet:
		//	ssid:=SSID{}
		//	s.db.Where(SSID{Anchor: ANCHOR}).First(&ssid)
		//
		//	//if err != nil {
		//	//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	//	return
		//	//}
		//	_, err := fmt.Fprint(w, ssid)
		//	if err != nil {
		//		logrus.Error(err)
		//	}
		case http.MethodPost:
			ssid:=SSID{}
			e := json.NewDecoder(r.Body).Decode(&ssid)

			if e!=nil {
				logrus.Error(e)
				http.Error(w, "err", http.StatusInternalServerError)
				return
			}
			ssid.Anchor=ANCHOR
			e = s.db.FirstOrCreate(&ssid, SSID{Anchor: ANCHOR}).Error
			if e!=nil {
				logrus.Error(e)
				http.Error(w, "err", http.StatusInternalServerError)
				return
			}

			_, err := fmt.Fprint(w, html.EscapeString(SUCCESS))
			if err != nil {
				logrus.Error(err)
			}

		default:
			http.Error(w, "Sorry, only GET and POST methods are supported.", http.StatusInternalServerError)
		}

	}
}

