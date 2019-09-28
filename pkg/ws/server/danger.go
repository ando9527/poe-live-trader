package server

import "net/http"

// Danger Valid nothing
func Danger(r *http.Request) bool {
	return true
}
