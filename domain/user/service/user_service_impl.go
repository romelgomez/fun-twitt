package user

import (
	"context"
	dto "funtwitt/domain/user/dto"
	model "funtwitt/domain/user/model"
	repository "funtwitt/domain/user/repository"
	"funtwitt/prisma/db"
)

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func NewUserServiceImpl(repo repository.UserRepository) UserService {
	return &UserServiceImpl{Repo: repo}
}

func (s *UserServiceImpl) Create(ctx context.Context, request dto.UserCreate) (dto.UserResponse, error) {
	user, err := s.Repo.Create(ctx, model.User{
		InnerUser: db.InnerUser{
			Email: request.Email,
			Name:  request.Name,
		},
	})
	if err != nil {
		return dto.UserResponse{}, err
	}

	return user.Dto(), nil
}

func (s *UserServiceImpl) FindByID(ctx context.Context, id string) (dto.UserResponse, error) {
	result, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return result.Dto(), nil
}

func (s *UserServiceImpl) Update(ctx context.Context, request dto.UserUpdate) (dto.UserResponse, error) {
	user, err := s.Repo.Update(ctx, model.User{
		InnerUser: db.InnerUser{
			ID:   request.ID,
			Name: request.Name,
		},
	})
	if err != nil {
		return dto.UserResponse{}, err
	}

	return user.Dto(), nil
}

func (s *UserServiceImpl) Delete(ctx context.Context, id string) (dto.UserDelete, error) {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return dto.UserDelete{}, err
	}

	return dto.UserDelete{
		Status:  "ok",
		Message: "deleted successfully",
	}, nil
}

func (s *UserServiceImpl) FindByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	user, err := s.Repo.FindByEmail(ctx, email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return user.Dto(), nil
}

func (s *UserServiceImpl) FindAll(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, user.Dto())
	}

	return responses, nil
}
