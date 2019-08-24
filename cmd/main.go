package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/ando9527/poe-live-trader/cmd/conf"
	"github.com/ando9527/poe-live-trader/pkg/audio"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/ando9527/poe-live-trader/pkg/trader"
	"github.com/atotto/clipboard"
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
	log.InitLogger(logLevel)
	logrus.Infof("Poe Live Trader %s", version)

	client := trader.NewTrader()
	client.Launch()
	whisper := client.Whisper
	for {
		select {
		case result := <-whisper:
			fmt.Println(result)
			audio.Play()
			err := clipboard.WriteAll(result)
			if err != nil {
				logrus.Warn("failed copy whisper to clipboard.")
			}
		}
	}

}
