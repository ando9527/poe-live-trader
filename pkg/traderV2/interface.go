package traderV2



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

