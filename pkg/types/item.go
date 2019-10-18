package types

type WsPool interface {
	GetBuilderChannel()<-chan ItemBuilder
	Run()
}

type HttpClient interface {
	RequestItemDetail(idList []string, filterID string) (ItemDetail, error)
}

type Item interface {
	GetNotification()string
	GetUserID()string
	SetUserID(string)
}

type ItemBuilder interface {

	SetWhisper(client HttpClient) (ItemBuilder, error)
	SetUserID() ItemBuilder
	Build() []Item
}