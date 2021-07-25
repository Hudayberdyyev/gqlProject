package graph

import (
	"context"
	"errors"
	"fmt"
	"github.com/Hudayberdyyev/gqlProject/graph/generated"
	"github.com/Hudayberdyyev/gqlProject/graph/model"
	"github.com/Hudayberdyyev/gqlProject/graph/models"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetups, err := r.MeetupsRepo.GetById(id)

	if err != nil || meetups == nil {
		return false, errors.New("meetups does not exists")
	}

	err = r.MeetupsRepo.Delete(meetups)

	if err != nil {
		return false, fmt.Errorf("error while deleting meetup: %v", err)
	}

	return true, nil
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := r.MeetupsRepo.GetById(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exists")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is not long enough")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}

	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("description is not long enough")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	meetup, err = r.MeetupsRepo.Update(meetup)

	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}

	return meetup, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("the name not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("the description not long enough")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}

	return r.MeetupsRepo.CreateMeetup(meetup)
}
