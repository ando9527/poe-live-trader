package conf

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger(level string) {
	// Setup logger format
	l, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Fatal(err.Error())
		os.Exit(1)
	}
	logrus.SetLevel(l)
}
