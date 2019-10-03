package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/ando9527/poe-live-trader/cmd/client/env"
	"github.com/ando9527/poe-live-trader/pkg/log"
	"github.com/ando9527/poe-live-trader/pkg/trader"
	"github.com/ando9527/poe-live-trader/pkg/ws"
	"github.com/briandowns/spinner"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	//_ "net/http/pprof"
)

var (
	version  string
)

func pause() {
	fmt.Print("Press 'Enter' to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func NotifyDC(cancel context.CancelFunc) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			select {
			case <-interrupt:
				cancel()
				time.Sleep(time.Second*2)
				return
			}
		}
	}()

}

func main() {
	//go func() {
	//	logrus.Info(http.ListenAndServe("localhost:6060", nil))
	//}()
	err := godotenv.Load("client.env")
	if err != nil {
		logrus.Error("Error loading client.env file")
		pause()
		return
	}
	cfg, err:= env.NewEnv()
	if err != nil {
		logrus.Error("Please setup .env file properly, ", err)
		pause()
		return
	}


	log.InitLogger(cfg.LogLevel, cfg.Logging)
	logrus.Infof("Poe Live Trader %s", version)

	config:=		ws.Config{
		POESSID: cfg.Poesessid,
		League:  cfg.League,
		Filter:  cfg.Filter,
	}

	ctx, cancel := context.WithCancel(context.Background())
	NotifyDC(cancel)
	client := trader.NewTrader(ctx, config)

	client.Launch()
	spinnerAnimation(ctx)

	for result:= range client.Whisper{
		logrus.Info(result)

		if cfg.Ignored[getName(result)]{
			logrus.Info("User in ignored list, ", getName(result))
			continue
		}




		client.Mutex.Lock()
		if client.IDCache[getName(result)]{
			logrus.Info("History duplicated user in cache, ",result)
			client.Mutex.Unlock()
			continue
		}
		client.KeySim.Message<-result
		client.IDCache[getName(result)]=true

		client.Mutex.Unlock()
	}
}



func getName(template string)(n string){
	tmp:=strings.Split(template, " ")[0]
	return strings.Replace(tmp,"@", "", 1)
}

func spinnerAnimation(ctx context.Context){
	go func(){
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()
		for{
			select {
				case <-ctx.Done():
				s.Stop()
				return
			}
		}
	}()
}