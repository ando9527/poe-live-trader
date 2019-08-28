package cloud

import (
	"net/http"
	"os"
)

func handleAuth(h http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		user , pass, _ := r.BasicAuth()
		if user!=os.Getenv("APP_USER") || pass != os.Getenv("APP_PASS"){
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		h(w,r)
	}
}
