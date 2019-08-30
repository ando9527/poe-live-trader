package request

import (
	"os"

	"github.com/ando9527/poe-live-trader/pkg/log"
)

func init(){
	log.InitLogger(os.Getenv("APP_LOG_LEVEL"))
}