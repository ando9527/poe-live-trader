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
		case http.MethodGet:
			ssid:=SSID{}
			e:=s.db.Where(SSID{Anchor: ANCHOR}).First(&ssid).Error
			if e != nil {
				logrus.Error(e)
				http.Error(w, "error", http.StatusInternalServerError)
				return
			}

			e = json.NewEncoder(w).Encode(&ssid)
			if e != nil {
				http.Error(w, "error", http.StatusInternalServerError)
				logrus.Error(e)
				return
			}
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
				http.Error(w, "error", http.StatusInternalServerError)
				logrus.Error(err)
				return
			}

		default:
			http.Error(w, "Sorry, only GET and POST methods are supported.", http.StatusInternalServerError)
		}

	}
}

