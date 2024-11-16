package follow

import (
	"context"
	model "funtwitt/domain/follow/model"
)

type FollowRepository interface {
	Create(ctx context.Context, data model.Follow) (model.Follow, error)
	FindByFollowerID(ctx context.Context, followerID string) ([]model.Follow, error)
	Delete(ctx context.Context, id string) error
}
