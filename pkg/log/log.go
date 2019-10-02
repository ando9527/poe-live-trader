package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

func InitLogger(level string) {
	// Setup logger format
	l, err := logrus.ParseLevel(level)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     false,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	if err != nil {
		logrus.Panic(err.Error())
		os.Exit(1)
	}
	logrus.SetLevel(l)

	f, err := os.OpenFile("logrus.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	mw := io.MultiWriter(os.Stdout, f)
	logrus.SetOutput(mw)
}

func InitCloudLogger(level string) {
	// Setup logger format
	l, err := logrus.ParseLevel(level)
	logrus.SetFormatter(log.NewFormatter())
	if err != nil {
		logrus.Panic(err.Error())
		os.Exit(1)
	}
	logrus.SetLevel(l)
}