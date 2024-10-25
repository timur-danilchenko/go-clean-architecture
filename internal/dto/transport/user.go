package dtotransport

import "github.com/google/uuid"

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetUserByIDRequest struct {
	ID uuid.UUID
}

type GetUserByIDResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}
