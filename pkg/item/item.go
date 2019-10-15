package item

import (
	"strings"

	"github.com/ando9527/poe-live-trader/pkg/types"
)

type Item struct{
	Notification string
	userID       string
}


func (i *Item) GetNotification() string {
	return i.Notification
}

func (i *Item) GetUserID() string {
	return i.userID
}

func (i *Item) SetUserID(id string)  {
	i.userID = id
}

type Builder struct{
	ItemList []types.Item
	IdList   []string
	FilterID string
}

func (b *Builder) SetWhisper(client types.HttpClient) (types.ItemBuilder, error) {
	detail, e := client.RequestItemDetail(b.IdList, b.FilterID)
	if e != nil {
		return nil, e
	}
	for _,v:=range detail.Result{
		b.ItemList = append(b.ItemList, &Item{
			Notification: v.Listing.Whisper,
			userID:       "",
		} )
	}
	return b, nil
}

func (b *Builder) SetUserID() types.ItemBuilder {
	for _,v:=range b.ItemList {
		v.SetUserID(getName(v.GetNotification()))
	}
	return b
}


func (b *Builder) Build() []types.Item {
	return b.ItemList
}

func NewBuilder() *Builder {
	return &Builder{
		ItemList: []types.Item{},
	}
}


func getName(template string)(n string){
	tmp:=strings.Split(template, " ")[0]
	return strings.Replace(tmp,"@", "", 1)
}