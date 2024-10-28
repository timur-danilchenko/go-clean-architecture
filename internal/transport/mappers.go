package transport

import (
	dtoservice "github.com/timur-danilchenko/project/internal/dto/service"
	dtotransport "github.com/timur-danilchenko/project/internal/dto/transport"
)

func mapCreateUserRequest(data *dtotransport.CreateUserRequest) *dtoservice.CreateUserRequest {
	return &dtoservice.CreateUserRequest{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Phone:     data.Phone,
	}
}

func mapCreateUserResponse(data *dtoservice.CreateUserResponse) *dtotransport.CreateUserResponse {
	return &dtotransport.CreateUserResponse{
		ID: data.ID,
	}
}

func mapGetUserByIDRequest(data *dtotransport.GetUserByIDRequest) *dtoservice.GetUserByIDRequest {
	return &dtoservice.GetUserByIDRequest{
		ID: data.ID,
	}
}

func mapGetUserByIDResponse(data *dtoservice.GetUserByIDResponse) *dtotransport.GetUserByIDResponse {
	return &dtotransport.GetUserByIDResponse{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Phone:     data.Phone,
	}
}
