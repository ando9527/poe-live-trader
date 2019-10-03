package log

import (
	"os"
	"time"

	"github.com/joonix/log"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

func InitLogger(level string, logging bool) {
	// Setup logger format
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Panic(err)
		os.Exit(1)
	}

	logrus.SetLevel(logLevel)
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	if !logging{
		return
	}

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logrus.log",
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Level:      logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC822,
		},
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	logrus.AddHook(rotateFileHook)

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