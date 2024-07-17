package service

import (
	"backend/internal/entities"
	"backend/internal/repository"
	"context"
)

type UserServicer interface {
	GetAllUsers(ctx context.Context) ([]entities.User, error)
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

func (u *UserService) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	//TODO implement me
	panic("implement me")
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
