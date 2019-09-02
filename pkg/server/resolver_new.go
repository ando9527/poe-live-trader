package server

import (
	"context"

	"github.com/ando9527/poe-live-trader/pkg/graphql/graph"
	"github.com/ando9527/poe-live-trader/pkg/graphql/models"
	"github.com/jinzhu/gorm"
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
	ssid:=models.Ssid{Anchor:ANCHOR}
	e:=r.db.First(&ssid).Error
	if e != nil {
		e := r.db.Create(&models.Ssid{
			Content: input.Content,
			Anchor:  ANCHOR,
		}).Error
		if e != nil {
			return nil, e
		}
	}

	e = r.db.Model(&ssid).Update("content", input.Content ).Error
	if e != nil {
		return nil,e
	}

	return &ssid, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ssid(ctx context.Context) (*models.Ssid, error) {
	ssid:=models.Ssid{Anchor:ANCHOR}
	e:=r.db.First(&ssid).Error
	if e != nil {
		return nil, e
	}
	return &ssid, nil
}
