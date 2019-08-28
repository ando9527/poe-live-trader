package main

import (
	"os"
	"strconv"

	"github.com/ando9527/poe-live-trader/cloud/env"
	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/sirupsen/logrus"
)

var (
	version  string
)


func main(){
	env.Verify()
	b, e := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if e != nil {
		panic(e)
	}
	log.InitCloudLogger(b)
	logrus.Info("version ", version )
	s := cloud.NewServer()
	s.Run()
}