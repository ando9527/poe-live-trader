package cloud

import (
	"net/http"
)

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
