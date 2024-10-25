package repository

import (
	"context"
	"database/sql"

	dtorepository "github.com/timur-danilchenko/project/internal/dto/repository"
	dtoservice "github.com/timur-danilchenko/project/internal/dto/service"
)

type UserRepository struct {
	Conn *sql.Conn
}

func (r *UserRepository) CreateUser(ctx context.Context, data dtoservice.CreateUserRequest) (dtoservice.CreateUserResponse, error) {
	userReq := mapCreateUserRequest(data)
	var userRes dtorepository.CreateUserResponse

	query := `insert into users(first_name, last_name, email, phone) values (
		$1, $2, $3, $4
	) returning id;`

	if err := r.Conn.QueryRowContext(ctx, query, userReq.FirstName, userReq.LastName, userReq.Email, userReq.Phone).Scan(&userRes.ID); err != nil {
		return dtoservice.CreateUserResponse{}, err
	}

	return mapCreateUserResponse(userRes), nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, data dtoservice.GetUserByIDRequest) (dtoservice.GetUserByIDResponse, error) {
	userReq := mapGetUserByIDRequest(data)
	var userRes dtorepository.GetUserByIDResponse

	query := `select first_name, last_name, email, phone from users where id=$1;`
	if err := r.Conn.QueryRowContext(ctx, query, userReq.ID).Scan(&userRes.FirstName, &userRes.LastName, &userRes.Email, &userRes.Phone); err != nil {
		return dtoservice.GetUserByIDResponse{}, err
	}

	userRes.ID = userReq.ID
	return mapGetUserByIDResponse(userRes), nil
}
