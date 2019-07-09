package ws

type Client struct {
}

func (client *Client) GetItemID() (itemID chan []string) {
	itemID = make(chan []string)
	itemID <- []string{"6bf0738f765b4d364fc65105910493c13b3d89ded2797cbcca32b99ca0579825"}
	return itemID
}

func NewWebsocketClient() (client *Client) {
	return &Client{}
}
