package tweet

import (
	"context"
	model "funtwitt/domain/tweet/model"
)

type TweetRepository interface {
	Create(ctx context.Context, data model.Tweet) (model.Tweet, error)
	FindByUserID(ctx context.Context, userID string) ([]model.Tweet, error)
	Delete(ctx context.Context, id string) error
}
