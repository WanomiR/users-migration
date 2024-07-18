package service

import (
	"backend/internal/entities"
	"backend/internal/lib/e"
	"backend/internal/repository"
	"context"
	"errors"
)

type UserServicer interface {
	GetUsers(ctx context.Context, limit, offset int) ([]*entities.User, error)
	GetUserByID(ctx context.Context, id int) (*entities.User, error)
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

func (u *UserService) GetUsers(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	if limit < 1 {
		return nil, errors.New("limit must be greater than zero")
	}

	usersAll, err := u.DB.List(ctx, limit, offset)
	if err != nil {
		return nil, e.WrapIfErr("failed to list users", err)
	}

	// exclude deleted users
	users := make([]*entities.User, 0, len(usersAll))
	for _, user := range usersAll {
		if !user.IsDeleted {
			users = append(users, user)
		}
	}

	return users, err
}

func (u *UserService) GetUserByID(ctx context.Context, id int) (*entities.User, error) {
	user, err := u.DB.GetByID(ctx, id)
	if err != nil {
		return nil, e.WrapIfErr("failed to get user", err)
	}

	if user.IsDeleted {
		return nil, e.WrapIfErr("failed to get user", errors.New("user not found"))
	}

	return user, nil
}

func (u *UserService) CreateUser(ctx context.Context, user entities.User) (int, error) {
	userId, err := u.DB.Create(ctx, user)
	if err != nil {
		return 0, e.WrapIfErr("failed to create user", err)
	}

	return userId, nil
}

func (u *UserService) DeleteUser(ctx context.Context, userId int) (err error) {
	defer func() { err = e.WrapIfErr("failed to delete user", err) }()

	var user *entities.User
	user, err = u.DB.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	if user.IsDeleted {
		return errors.New("user not found")
	}

	if err = u.DB.Delete(ctx, userId); err != nil {
		return err
	}

	return nil
}

func (u *UserService) UpdateUser(ctx context.Context, userInput entities.User) (err error) {
	defer func() { err = e.WrapIfErr("failed to update user", err) }()

	var user *entities.User
	user, err = u.DB.GetByID(ctx, userInput.Id)
	if err != nil {
		return err
	}

	if user.IsDeleted {
		return errors.New("user not found")
	}

	if err = u.DB.Update(ctx, userInput); err != nil {
		return err
	}

	return nil
}
