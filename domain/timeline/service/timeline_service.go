package tweet

import (
	"context"
	dto "funtwitt/domain/tweet/dto"
)

type TimelineService interface {
	GetTimeline(ctx context.Context, userID string) ([]dto.TweetResponse, error)
}
