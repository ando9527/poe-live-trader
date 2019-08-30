package main

import (
	"github.com/ando9527/poe-live-trader/cmd/cloud/env"
	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/sirupsen/logrus"
)

var (
	version  string
)


func main(){
	cfg := env.NewEnv()
	log.InitCloudLogger(cfg.LogLevel)
	logrus.Info("version ", version )
	s := cloud.NewServer(cfg.Dsn, cfg.User, cfg.Pass)
	s.Run()
}