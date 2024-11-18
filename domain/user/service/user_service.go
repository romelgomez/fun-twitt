package user

import (
	"context"
	dto "funtwitt/domain/user/dto"
)

type UserService interface {
	Create(ctx context.Context, request dto.UserCreate) (dto.UserResponse, error)
	FindByID(ctx context.Context, id string) (dto.UserResponse, error)
	Update(ctx context.Context, request dto.UserUpdate) (dto.UserResponse, error)
	Delete(ctx context.Context, id string) (dto.UserDelete, error)
	FindByEmail(ctx context.Context, email string) (dto.UserResponse, error)
	FindAll(ctx context.Context) ([]dto.UserResponse, error)
}
