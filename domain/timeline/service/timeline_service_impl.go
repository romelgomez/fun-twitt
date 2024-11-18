package tweet

import (
	"context"
	repository "funtwitt/domain/timeline/repository"
	dto "funtwitt/domain/tweet/dto"
)

type TimelineServiceImpl struct {
	Repo repository.TimelineRepository
}

func NewTimelineService(repo repository.TimelineRepository) TimelineService {
	return &TimelineServiceImpl{Repo: repo}
}

func (s *TimelineServiceImpl) GetTimeline(ctx context.Context, id string) ([]dto.TweetResponse, error) {
	tweets, err := s.Repo.GetTimeline(ctx, id)
	if err != nil {
		return nil, err
	}

	var responses []dto.TweetResponse
	for _, tweet := range tweets {
		responses = append(responses, tweet.Dto())
	}

	return responses, nil
}
