package repository

import (
	"context"
	"database/sql"

	dtorepository "github.com/timur-danilchenko/project/internal/dto/repository"
)

type UserRepository struct {
	Conn *sql.Conn
}

func (r *UserRepository) CreateUser(ctx context.Context, user dtorepository.CreateUserRequest) (dtorepository.CreateUserResponse, error) {
	var result dtorepository.CreateUserResponse

	query := `insert into users(first_name, last_name, email, phone) values (
		$1, $2, $3, $4
	) returning id;`
	params := []interface{}{
		user.FirstName, user.LastName,
		user.Email, user.Phone,
	}

	err := r.Conn.QueryRowContext(ctx, query, params).Scan(&result.ID)
	if err != nil {
		return dtorepository.CreateUserResponse{}, err
	}

	return result, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, user dtorepository.GetUserByIDRequest) (dtorepository.GetUserByIDResponse, error) {
	var result dtorepository.GetUserByIDResponse

	query := `select first_name, last_name, email, phone from users where id=$1;`
	row := r.Conn.QueryRowContext(ctx, query, user.ID)

	err := row.Scan(&result.FirstName, &result.LastName, &result.Email, &result.Phone)
	if err != nil {
		return dtorepository.GetUserByIDResponse{}, err
	}

	result.ID = user.ID
	return result, nil
}
