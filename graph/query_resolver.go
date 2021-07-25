package graph

import (
	"context"
	"github.com/Hudayberdyyev/gqlProject/graph/generated"
	"github.com/Hudayberdyyev/gqlProject/graph/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UsersRepo.GetUserByID(id)
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
