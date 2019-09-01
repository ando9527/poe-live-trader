package graphql

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	db *gorm.DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateOrUpdateSsid(ctx context.Context, input NewSsid) (*Ssid, error) {
	ssid:=Ssid{Content:input.Content}
	ssid.Anchor=ANCHOR
	e := r.db.Model(&ssid).Update("content", input.Content ).Error
	if e!=nil {
		e := r.db.Create(&ssid).Error
		if e != nil {
			return nil, e
		}
	}
	return &ssid, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ssid(ctx context.Context) (*Ssid, error) {
	ssid:=Ssid{}
	e:=r.db.Where(Ssid{Anchor: ANCHOR}).First(&ssid).Error
	if e != nil {
		return nil, e
	}
	logrus.Debugf("Query %s", ssid.Content)
	return &ssid, nil
}

