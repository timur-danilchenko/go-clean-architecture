package service

import (
	"context"

	dtoservice "github.com/timur-danilchenko/project/internal/dto/service"
	"github.com/timur-danilchenko/project/internal/repository"
)

type UserService struct {
	Repository *repository.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, user *dtoservice.CreateUserRequest) (*dtoservice.CreateUserResponse, error) {
	userReq := mapCreateUserRequest(user)

	userRes, err := s.Repository.CreateUser(ctx, userReq)
	if err != nil {
		return nil, err
	}

	return mapCreateUserResponse(userRes), nil
}

func (s *UserService) GetUserByID(ctx context.Context, user *dtoservice.GetUserByIDRequest) (*dtoservice.GetUserByIDResponse, error) {
	userReq := mapGetUserByIDRequest(user)

	userRes, err := s.Repository.GetUserByID(ctx, userReq)
	if err != nil {
		return nil, err
	}

	return mapGetUserByIDResponse(userRes), nil
}
