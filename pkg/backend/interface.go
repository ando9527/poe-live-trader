package backend



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
	AllowSend(id string) bool
	Run()
}

type Notifier interface {
	Run()
	SendToQueue(m string)
}

type WsServer interface {
	Run()
	SendToServer(m string)
}

type GrpcGateway interface {
	Run()
}