package traderV2



type AudioPlayer interface {
	Init()
	Play(string, float64)
}

type Database interface {
	IsIgnored(id string) bool
	Connect(name string)
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

