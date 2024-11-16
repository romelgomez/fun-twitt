package follow

import (
	"context"
	dto "funtwitt/domain/follow/dto"
)

type FollowService interface {
	Create(ctx context.Context, request dto.FollowCreate) (dto.FollowResponse, error)
	FindByFollowerID(ctx context.Context, followerID string) ([]dto.FollowResponse, error)
	Delete(ctx context.Context, id string) error
}
