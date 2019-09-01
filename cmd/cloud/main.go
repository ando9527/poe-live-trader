package main

import (
	"github.com/ando9527/poe-live-trader/cmd/cloud/env"
	"github.com/ando9527/poe-live-trader/pkg/graphql"
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
	s := graphql.NewServer(cfg.Dsn, cfg.User, cfg.Pass, cfg.LogLevel)
	s.Run()
}