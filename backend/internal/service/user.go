package service

import (
	"backend/internal/entities"
	"backend/internal/lib/e"
	"backend/internal/repository"
	"context"
	"errors"
)

type UserServicer interface {
	GetUsers(ctx context.Context, limit, offset int) ([]entities.User, error)
	GetUserByID(ctx context.Context, id int) (entities.User, error)
	CreateUser(ctx context.Context, user entities.User) (int, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, user entities.User) error
}

type UserService struct {
	DB repository.UserRepository
}

func NewUserService(db repository.UserRepository) *UserService {
	return &UserService{DB: db}
}

func (u *UserService) GetUsers(ctx context.Context, limit, offset int) ([]entities.User, error) {

	if limit < 1 {
		return nil, errors.New("limit must be greater than zero")
	}

	users, err := u.DB.List(ctx, limit, offset)
	if err != nil {
		return nil, e.WrapIfErr("failed to list users:", err)
	}

	return users, err
}

func (u *UserService) GetUserByID(ctx context.Context, id int) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) CreateUser(ctx context.Context, user entities.User) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UpdateUser(ctx context.Context, user entities.User) error {
	//TODO implement me
	panic("implement me")
}
