package log

import (
	"os"
	"time"

	"github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

func InitLogger(debug bool) {
	logLevel := "info"
	if debug == true {
		logLevel = "debug"
	}
	// Setup logger format
	l, err := logrus.ParseLevel(logLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	if err != nil {
		logrus.Panic(err.Error())
		os.Exit(1)
	}
	logrus.SetLevel(l)
}

func InitCloudLogger(debug bool) {
	logLevel := "info"
	if debug == true {
		logLevel = "debug"
	}
	// Setup logger format
	l, err := logrus.ParseLevel(logLevel)
	logrus.SetFormatter(log.NewFormatter())
	if err != nil {
		logrus.Panic(err.Error())
		os.Exit(1)
	}
	logrus.SetLevel(l)
}