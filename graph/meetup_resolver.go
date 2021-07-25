package graph

import (
	"context"
	"github.com/Hudayberdyyev/gqlProject/graph/generated"
	"github.com/Hudayberdyyev/gqlProject/graph/models"
)

type meetupResolver struct{ *Resolver }

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return getUserLoader(ctx).Load(obj.UserID)
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }
