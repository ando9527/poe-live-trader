package main

import (
	"flag"

	"github.com/sirupsen/logrus"

	"github.com/ando9527/poe-live-trader/pkg/conf"
	"github.com/ando9527/poe-live-trader/pkg/data/http"
	"github.com/ando9527/poe-live-trader/pkg/data/ws"
)

var (
	logLevel string
	version  string
)

func main() {
	flag.StringVar(&logLevel, "l", "info", "Logging level")
	flag.Parse()
	conf.InitLogger(logLevel)
	logrus.Infof("Poe Live Trader %s", version)
	ws.Connect(http.GetItemDetail)

}
