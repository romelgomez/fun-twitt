package user

import (
	"context"
	model "funtwitt/domain/user/model"
)

type UserRepository interface {
	Create(ctx context.Context, data model.User) (model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	FindByID(ctx context.Context, id string) (model.User, error)
	FindBySortID(ctx context.Context, id string) (model.User, error)
	Update(ctx context.Context, input model.User) (model.User, error)
	Delete(ctx context.Context, id string) error
	FindByEmail(ctx context.Context, email string) (model.User, error)
}
