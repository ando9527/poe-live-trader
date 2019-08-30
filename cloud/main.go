package main

import (
	"github.com/ando9527/poe-live-trader/cloud/conf"
	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/sirupsen/logrus"
)

var (
	version  string
)


func main(){
	cfg := conf.NewConfig()
	log.InitCloudLogger(cfg.LogLevel)
	logrus.Info("version ", version )
	s := cloud.NewServer()
	s.Run()
}