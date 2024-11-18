package user

import (
	"context"
	"errors"
	model "funtwitt/domain/user/model"
	helper "funtwitt/helper"
	"funtwitt/prisma/db"
)

type UserRepositoryImpl struct {
	Db *db.PrismaClient
}

func NewUserRepository(dbClient *db.PrismaClient) UserRepository {
	return &UserRepositoryImpl{Db: dbClient}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, data model.User) (model.User, error) {
	id, err := helper.GeneratePrefixedID("user_", 15)
	if err != nil {
		return model.User{}, err
	}

	record, err := r.Db.User.CreateOne(
		db.User.ID.Set(id),
		db.User.Email.Set(data.Email),
		db.User.Name.Set(data.Name),
	).Exec(ctx)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		InnerUser: record.InnerUser,
	}, nil
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context) ([]model.User, error) {
	records, err := r.Db.User.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]model.User, len(records))
	for i, record := range records {
		result[i] = model.User{
			InnerUser: record.InnerUser,
		}
	}

	return result, nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id string) (model.User, error) {
	record, err := r.Db.User.FindUnique(db.User.ID.Equals(id)).Exec(ctx)
	if err != nil {
		return model.User{}, err
	}
	if record == nil {
		return model.User{}, errors.New("user not found")
	}

	return model.User{
		InnerUser: record.InnerUser,
	}, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, input model.User) (model.User, error) {
	record, err := r.Db.User.FindUnique(db.User.ID.Equals(input.ID)).Update(
		db.User.Name.Set(input.Name),
	).Exec(ctx)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		InnerUser: record.InnerUser,
	}, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := r.Db.User.FindUnique(db.User.ID.Equals(id)).Delete().Exec(ctx)
	return err
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (model.User, error) {
	record, err := r.Db.User.FindFirst(
		db.User.Email.Equals(email),
	).Exec(ctx)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		InnerUser: record.InnerUser,
	}, nil
}
