package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main(){
	s := NewServer()
	s.Routes()
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "8080"
	}

	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), s.router))
}