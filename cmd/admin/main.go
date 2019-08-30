package main

import (
	"flag"
	"fmt"

	"github.com/ando9527/poe-live-trader/cmd/admin/conf"
	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)
var (poessid string
     migration bool
)
func init(){
	flag.BoolVar(&migration, "m", false, "execute migration?")
	flag.StringVar(&poessid, "i", "", "Insert POESSID")
}
func main(){
	e := godotenv.Load("admin.env")
	if e != nil {
		logrus.Panic(e)
	}
	cfg := conf.NewConfig()
	log.InitLogger(cfg.LogLevel)
	flag.Parse()
	if poessid!=""{
		fmt.Println("poessid", poessid)
		fmt.Println(cfg.CloudUrl)
		cloud.PostSSID(cfg.CloudUrl, poessid, cfg.User, cfg.User)
	}

	if migration==true{
		s := cloud.NewServer(cfg.Dsn, cfg.User,cfg.Pass)
		s.Connect()
		s.Migration()
	}
	//

	//conf.NewConfig()
	//ctx:=context.Background()
	//
	//
	//if e != nil {
	//	logrus.Panic(e)
	//}
	//c, err := cloud.NewClient(ctx)
	//if err != nil {
	//	logrus.Panic(err)
	//}
	//err = c.UpdateInsert(poessid)
	//if err != nil {
	//	logrus.Panic(err)
	//}
	//fmt.Println("success!")

}


