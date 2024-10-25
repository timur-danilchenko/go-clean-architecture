package dtoservice

import "github.com/google/uuid"

type CreateUserRequest struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type CreateUserResponse struct {
	ID uuid.UUID
}

type GetUserByIDRequest struct {
	ID uuid.UUID
}

type GetUserByIDResponse struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
