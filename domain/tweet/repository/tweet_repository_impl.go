package tweet

import (
	"context"
	model "funtwitt/domain/tweet/model"
	"funtwitt/helper"
	"funtwitt/prisma/db"
)

type TweetRepositoryImpl struct {
	Db *db.PrismaClient
}

func NewTweetRepository(dbClient *db.PrismaClient) TweetRepository {
	return &TweetRepositoryImpl{Db: dbClient}
}

func (r *TweetRepositoryImpl) Create(ctx context.Context, data model.Tweet) (model.Tweet, error) {
	shortID, err := helper.GeneratePrefixedID("tweet_", 20)
	if err != nil {
		return model.Tweet{}, err
	}

	record, err := r.Db.Tweet.CreateOne(
		db.Tweet.Content.Set(data.Content),
		db.Tweet.User.Link(db.User.ID.Equals(data.UserID)),
		db.Tweet.SortID.Set(shortID),
	).Exec(ctx)
	if err != nil {
		return model.Tweet{}, err
	}

	return model.Tweet{
		InnerTweet: record.InnerTweet,
	}, nil
}

func (r *TweetRepositoryImpl) FindByUserID(ctx context.Context, userID string) ([]model.Tweet, error) {
	records, err := r.Db.Tweet.FindMany(db.Tweet.UserID.Equals(userID)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]model.Tweet, len(records))
	for i, record := range records {
		result[i] = model.Tweet{
			InnerTweet: record.InnerTweet,
		}
	}

	return result, nil
}

func (r *TweetRepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := r.Db.Tweet.FindUnique(db.Tweet.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
