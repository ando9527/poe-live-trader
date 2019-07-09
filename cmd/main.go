package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	client2 "github.com/ando9527/poe-live-trader/pkg/client/v2/"

	"github.com/ando9527/poe-live-trader/conf"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	logLevel string
	version  string
)

func pause() {
	fmt.Print("Press 'Enter' to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file")
		pause()
		return
	}
	err = conf.InitConfig()
	if err != nil {
		logrus.Error(err)
		logrus.Error("Please setup .env file properly")
		pause()
		return
	}

	flag.StringVar(&logLevel, "l", "info", "Logging level")
	flag.Parse()
	conf.InitLogger(logLevel)
	logrus.Infof("Poe Live Trader %s", version)

	client := client2.NewClient()

	whisper := client.GetWhisper()
	for {
		fmt.Println(<-whisper)
	}
	//client := client2.NewClient()
	//client.ReConnect()
	//
	//h := &client2.Handler{}
	//client.ReadMessage(h)
	//client.NotifyDC()
}
