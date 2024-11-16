package tweet

import (
	"context"
	"errors"
	model "funtwitt/domain/tweet/model"
	"funtwitt/prisma/db"
)

type TimelineRepositoryImpl struct {
	Db *db.PrismaClient
}

func NewTimelineRepository(dbClient *db.PrismaClient) TimelineRepository {
	return &TimelineRepositoryImpl{Db: dbClient}
}

func (r *TimelineRepositoryImpl) GetTimeline(ctx context.Context, userID string) ([]model.Tweet, error) {
	followees, err := r.Db.Follow.FindMany(
		db.Follow.FollowerID.Equals(userID),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var followeeIDs []string
	for _, followee := range followees {
		followeeIDs = append(followeeIDs, followee.FolloweeID)
	}

	if len(followeeIDs) == 0 {
		return nil, errors.New("no followees found for the user")
	}

	tweets, err := r.Db.Tweet.FindMany(
		db.Tweet.UserID.In(followeeIDs),
	).OrderBy(
		db.Tweet.CreatedAt.Order(db.SortOrderDesc),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]model.Tweet, len(tweets))
	for i, tweet := range tweets {
		result[i] = model.Tweet{
			InnerTweet: tweet.InnerTweet,
		}
	}

	return result, nil
}
