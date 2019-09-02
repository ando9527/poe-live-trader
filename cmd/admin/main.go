package main

//import (
//	"flag"
//	"fmt"
//
//	"github.com/ando9527/poe-live-trader/cmd/admin/env"
//	"github.com/ando9527/poe-live-trader/pkg/log"
//	"github.com/joho/godotenv"
//	"github.com/sirupsen/logrus"
//)
//var (poessid string
//     migration bool
//)
//func init(){
//	flag.BoolVar(&migration, "m", false, "execute migration?")
//	flag.StringVar(&poessid, "i", "", "Insert POESSID")
//}
//func main(){
//	e := godotenv.Load("admin.env")
//	if e != nil {
//		logrus.Panic(e)
//	}
//	env := env.NewEnv()
//	log.InitLogger(env.LogLevel)
//	flag.Parse()
//	if poessid!=""{
//		fmt.Println("poessid", poessid)
//		fmt.Println(env.CloudUrl)
//		cloud.PostSSID(env.CloudUrl, poessid, env.User, env.User)
//	}
//
//	if migration==true{
//		s := cloud.NewServer(env.Dsn, env.User, env.Pass, env.LogLevel)
//		s.Connect()
//		s.Migration()
//	}
	//

	//conf.NewEnv()
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

//}


