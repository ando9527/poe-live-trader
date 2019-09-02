package server

import (
	"context"

	"github.com/ando9527/poe-live-trader/pkg/graphql/graph"
	"github.com/ando9527/poe-live-trader/pkg/graphql/models"
	"github.com/ando9527/poe-live-trader/pkg/types"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{ db *gorm.DB}

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateOrUpdateSsid(ctx context.Context, input models.NewSsid) (*models.Ssid, error) {
	ssid:=models.Ssid{Content:input.Content}
	ssid.Anchor=types.ANCHOR
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

func (r *queryResolver) Ssid(ctx context.Context) (*models.Ssid, error) {
	ssid:=models.Ssid{}
	e:=r.db.Where(models.Ssid{Anchor: types.ANCHOR}).First(&ssid).Error
	if e != nil {
		return nil, e
	}
	logrus.Debugf("Query %s", ssid.Content)
	return &ssid, nil
}
