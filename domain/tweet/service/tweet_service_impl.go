package tweet

import (
	"context"
	dto "funtwitt/domain/tweet/dto"
	model "funtwitt/domain/tweet/model"
	repository "funtwitt/domain/tweet/repository"
	"funtwitt/prisma/db"
)

type TweetServiceImpl struct {
	Repo repository.TweetRepository
}

func NewTweetService(repo repository.TweetRepository) TweetService {
	return &TweetServiceImpl{Repo: repo}
}

func (s *TweetServiceImpl) Create(ctx context.Context, request dto.TweetCreate) (dto.TweetResponse, error) {
	tweet, err := s.Repo.Create(ctx, model.Tweet{
		InnerTweet: db.InnerTweet{
			Content: request.Content,
			UserID:  request.UserID,
		},
	})
	if err != nil {
		return dto.TweetResponse{}, err
	}

	return tweet.Dto(), nil
}

func (s *TweetServiceImpl) FindByUserID(ctx context.Context, userID string) ([]dto.TweetResponse, error) {
	tweets, err := s.Repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []dto.TweetResponse
	for _, tweet := range tweets {
		responses = append(responses, tweet.Dto())
	}

	return responses, nil
}

func (s *TweetServiceImpl) Delete(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}
