package types

type HttpClient interface {
	RequestItemDetail(idList []string, filterID string) (ItemDetail, error)
}

type Item interface {
	Notify()
	GetNotification()string
	GetUserID()string
}

type ItemBuilder interface {
	//FilterID from cfg file
	//SetFilter(string) ItemBuilder
	//ID received from GGG websocket server
	//SetID([]string) ItemBuilder
	//Get detail from http api
	SetWhisper(client *HttpClient) ItemBuilder
	SetUserID() ItemBuilder
	Build() Item
}