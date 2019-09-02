package resolver
//
//import (
//	"context"
//
//	"github.com/ando9527/poe-live-trader/pkg/graphql/graph"
//	"github.com/ando9527/poe-live-trader/pkg/graphql/models"
//)
//
//// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
//
//type Resolver struct{}
//
//func (r *Resolver) Mutation() graph.MutationResolver {
//	return &mutationResolver{r}
//}
//func (r *Resolver) Query() graph.QueryResolver {
//	return &queryResolver{r}
//}
//
//type mutationResolver struct{ *Resolver }
//
//func (r *mutationResolver) CreateOrUpdateSsid(ctx context.Context, input models.NewSsid) (*models.Ssid, error) {
//	panic("not implemented")
//}
//
//type queryResolver struct{ *Resolver }
//
//func (r *queryResolver) Ssid(ctx context.Context) (*models.Ssid, error) {
//	panic("not implemented")
//}
