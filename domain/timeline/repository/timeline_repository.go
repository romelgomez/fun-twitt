package tweet

import (
	"context"
	model "funtwitt/domain/tweet/model"
)

type TimelineRepository interface {
	GetTimeline(ctx context.Context, userID string) ([]model.Tweet, error)
}
