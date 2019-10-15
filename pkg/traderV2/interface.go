package traderV2

import (
	"github.com/ando9527/poe-live-trader/pkg/types"
)

type WsPool interface {
	GetBuilderChannel()<-chan types.ItemBuilder
	Run()
}

type AudioPlayer interface {
	Init()
	Play(string, float64)
}

type Database interface {
	isIgnored(string) bool
	Connect()
	Migration()
}

type IDCache interface {
	AllowSend(string) bool
	Run()
}

type Notifier interface {
	Run()
	SendToQueue(string)
}

