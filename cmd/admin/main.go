package main

import (
	"flag"

	"github.com/ando9527/poe-live-trader/pkg/cloud"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)
var (poessid string
     migration bool
)
func init(){
	flag.BoolVar(&migration, "m", false, "execute migration?")
	flag.StringVar(&poessid, "p", "", "Insert POESSID")
}
func main(){
	e := godotenv.Load("admin.env")
	if e != nil {
		logrus.Panic(e)
	}
	flag.Parse()
	//if poessid==""{
	//	flag.Usage()
	//}
	if migration==true{
		s := cloud.NewServer()
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


