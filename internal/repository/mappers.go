package repository

import (
	dtorepository "github.com/timur-danilchenko/project/internal/dto/repository"
	dtoservice "github.com/timur-danilchenko/project/internal/dto/service"
)

func mapCreateUserRequest(data dtoservice.CreateUserRequest) dtorepository.CreateUserRequest {
	return dtorepository.CreateUserRequest{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Phone:     data.Phone,
	}
}

func mapCreateUserResponse(data dtorepository.CreateUserResponse) dtoservice.CreateUserResponse {
	return dtoservice.CreateUserResponse{
		ID: data.ID,
	}
}

func mapGetUserByIDRequest(data dtoservice.GetUserByIDRequest) dtorepository.GetUserByIDRequest {
	return dtorepository.GetUserByIDRequest{
		ID: data.ID,
	}
}

func mapGetUserByIDResponse(data dtorepository.GetUserByIDResponse) dtoservice.GetUserByIDResponse {
	return dtoservice.GetUserByIDResponse{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Phone:     data.Phone,
	}
}
