package follow

import (
	"context"
	dto "funtwitt/domain/follow/dto"
	model "funtwitt/domain/follow/model"
	repository "funtwitt/domain/follow/repository"
	"funtwitt/prisma/db"
)

type FollowServiceImpl struct {
	Repo repository.FollowRepository
}

func NewFollowService(repo repository.FollowRepository) FollowService {
	return &FollowServiceImpl{Repo: repo}
}

func (s *FollowServiceImpl) Create(ctx context.Context, request dto.FollowCreate) (dto.FollowResponse, error) {
	follow, err := s.Repo.Create(ctx, model.Follow{
		InnerFollow: db.InnerFollow{
			FollowerID: request.FollowerID,
			FolloweeID: request.FolloweeID,
		},
	})
	if err != nil {
		return dto.FollowResponse{}, err
	}

	return follow.Dto(), nil
}

func (s *FollowServiceImpl) FindByFollowerID(ctx context.Context, followerID string) ([]dto.FollowResponse, error) {
	follows, err := s.Repo.FindByFollowerID(ctx, followerID)
	if err != nil {
		return nil, err
	}

	var responses []dto.FollowResponse
	for _, follow := range follows {
		responses = append(responses, follow.Dto())
	}

	return responses, nil
}

func (s *FollowServiceImpl) Delete(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}
