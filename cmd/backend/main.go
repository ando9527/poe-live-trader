package main

import (
	"github.com/ando9527/poe-live-trader/cmd/client/env"
	"github.com/ando9527/poe-live-trader/pkg/backend"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	version  string
)



func main() {
	err := godotenv.Load("client.env")
	if err != nil {
		logrus.Error("Error loading client.env file")
		return
	}
	cfg := env.NewClient()
	cfg.Init()

	log.InitLogger(cfg.LogLevel, true)
	logrus.Infof("Poe Live Trader %s", version)
	logrus.Debug("Debug mode on")

	t:= backend.NewClient(cfg)
	t.Run()
}



