package main

import (
	"flag"

	"github.com/ando9527/poe-live-trader/pkg/conf"
	"github.com/ando9527/poe-live-trader/pkg/data/http"
	"github.com/ando9527/poe-live-trader/pkg/data/ws"
)

var logLevel string

func main() {
	flag.StringVar(&logLevel, "l", "info", "Logging level")
	flag.Parse()
	conf.InitLogger(logLevel)
	//data.Connect()
	ws.Connect(http.GetItemDetail)

}
