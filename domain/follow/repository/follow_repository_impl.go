package follow

import (
	"context"
	model "funtwitt/domain/follow/model"
	"funtwitt/helper"
	"funtwitt/prisma/db"
)

type FollowRepositoryImpl struct {
	Db *db.PrismaClient
}

func NewFollowRepository(dbClient *db.PrismaClient) FollowRepository {
	return &FollowRepositoryImpl{Db: dbClient}
}

func (r *FollowRepositoryImpl) Create(ctx context.Context, data model.Follow) (model.Follow, error) {
	id, err := helper.GeneratePrefixedID("fw_", 20)
	if err != nil {
		return model.Follow{}, err
	}

	record, err := r.Db.Follow.CreateOne(
		db.Follow.ID.Set(id),
		db.Follow.Follower.Link(db.User.ID.Equals(data.FollowerID)),
		db.Follow.Followee.Link(db.User.ID.Equals(data.FolloweeID)),
	).Exec(ctx)
	if err != nil {
		return model.Follow{}, err
	}

	return model.Follow{
		InnerFollow: record.InnerFollow,
	}, nil
}

func (r *FollowRepositoryImpl) FindByFollowerID(ctx context.Context, followerID string) ([]model.Follow, error) {
	records, err := r.Db.Follow.FindMany(db.Follow.FollowerID.Equals(followerID)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]model.Follow, len(records))
	for i, record := range records {
		result[i] = model.Follow{
			InnerFollow: record.InnerFollow,
		}
	}

	return result, nil
}

func (r *FollowRepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := r.Db.Follow.FindUnique(db.Follow.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
