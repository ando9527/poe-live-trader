package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ando9527/poe-live-trader/cmd/admin/env"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/ando9527/poe-live-trader/pkg/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)
var (poessid string
    migration bool
)
func init(){
	flag.BoolVar(&migration, "m", false, "Execute migration")
	flag.StringVar(&poessid, "i", "", "Insert POESSID")
}
func main(){
	e := godotenv.Load("admin.env")
	if e != nil {
		logrus.Panic(e)
	}
	env := env.NewEnv()
	log.InitLogger(env.LogLevel)
	flag.Parse()
	if len(os.Args)<=1{
		flag.Usage()
	}

	//insert poessid
	if poessid!=""{
		fmt.Println("poessid", poessid)
		fmt.Println(env.CloudUrl)
		e := server.UpdateSSID(env.CloudUrl, poessid, env.User, env.Pass)
		if e != nil {
			panic(e)
		}
	}
	//migration
	if migration==true{
		s := server.NewServer(env.Dsn, env.User, env.Pass, env.LogLevel, true)
		s.Connect()
		s.Migration()
	}

}


