package main

import (
	"flag"

	"github.com/ando9527/poe-live-trader/pkg/conf"
)

var logLevel string

func main() {
	flag.StringVar(&logLevel, "l", "info", "Logging level")
	flag.PrintDefaults()
	conf.InitLogger(logLevel)
}
