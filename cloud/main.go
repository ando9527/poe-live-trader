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

	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), s.router))
}