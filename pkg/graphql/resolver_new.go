package graphql

import (
	"context"

	"github.com/jinzhu/gorm"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	db *gorm.DB
}

func New()(r *Resolver)  {
	r=&Resolver{
		db: nil,
	}
	return r
}


func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateSsid(ctx context.Context, input NewSsid) (*Ssid, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ssid(ctx context.Context) (*Ssid, error) {
	panic("not implemented")
}
