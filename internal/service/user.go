package service

import (
	"context"

	dtotransport "github.com/timur-danilchenko/project/internal/dto/transport"
	"github.com/timur-danilchenko/project/internal/repository"
)

type UserService struct {
	Repository *repository.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, data dtotransport.CreateUserRequest) (dtotransport.CreateUserResponse, error) {
	userReq := mapCreateUserRequest(data)

	userRes, err := s.Repository.CreateUser(ctx, userReq)
	if err != nil {
		return dtotransport.CreateUserResponse{}, err
	}

	return mapCreateUserResponse(userRes), nil
}

func (s *UserService) GetUserByID(ctx context.Context, data dtotransport.GetUserByIDRequest) (dtotransport.GetUserByIDResponse, error) {
	userReq := mapGetUserByIDRequest(data)

	userRes, err := s.Repository.GetUserByID(ctx, userReq)
	if err != nil {
		return dtotransport.GetUserByIDResponse{}, err
	}

	return mapGetUserByIDResponse(userRes), nil
}
