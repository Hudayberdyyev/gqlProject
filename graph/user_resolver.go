package graph

import (
	"context"
	"github.com/Hudayberdyyev/gqlProject/graph/generated"
	"github.com/Hudayberdyyev/gqlProject/graph/models"
)

type userResolver struct{ *Resolver }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetupsForUser(obj)
}
