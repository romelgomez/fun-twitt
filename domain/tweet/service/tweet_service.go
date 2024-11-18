package tweet

import (
	"context"
	dto "funtwitt/domain/tweet/dto"
)

type TweetService interface {
	Create(ctx context.Context, request dto.TweetCreate) (dto.TweetResponse, error)
	FindByUserID(ctx context.Context, id string) ([]dto.TweetResponse, error)
	Delete(ctx context.Context, id string) error
}
