package service

import (
	"backend/internal/entities"
	"backend/internal/lib/e"
	"backend/internal/service/mock_repository"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

//go:generate mockgen -source=../repository/repository.go -destination=./mock_repository/mock_repository.go

func TestUserService_GetUsers(t *testing.T) {
	type params struct {
		offset, limit int
	}

	testCases := []struct {
		name         string
		params       params
		wantError    bool
		wantUsersLen int
	}{
		{"normal case", params{0, 1}, false, 1},
		{"bigger limit", params{0, 3}, false, 2},
		{"too small limit", params{0, 0}, true, 0},
		{"offset and limit on the edge", params{2, 3}, false, 1},
		{"offset and limit outside the range", params{4, 5}, false, 0},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	db := NewMockRepository(controller)
	s := NewUserService(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			users, err := s.GetUsers(context.Background(), tc.params.offset, tc.params.limit)

			if (err != nil) != tc.wantError {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tc.wantError)
			}

			if len(users) != tc.wantUsersLen {
				t.Errorf("GetUsers() len(users) = %v, want %v", len(users), tc.wantUsersLen)
			}
		})
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	testCases := []struct {
		name    string
		userId  int
		wantErr bool
	}{
		{"existing user", 1, false},
		{"non-existing user", -1, true},
		{"deleted user", 3, true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	db := NewMockRepository(controller)
	s := NewUserService(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := s.GetUserByID(context.Background(), tc.userId)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	testCases := []struct {
		name    string
		user    entities.User
		wantErr bool
	}{
		{"normal case", entities.User{FirstName: "Li", LastName: "Chi", Email: "li.chi@asia.com", Password: "password"}, false},
		{"too short name", entities.User{FirstName: "L", LastName: "Chi", Email: "admin@example.com", Password: "password"}, true},
		{"too short last name", entities.User{FirstName: "Li", LastName: "C", Email: "admin@example.com", Password: "password"}, true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	db := NewMockRepository(controller)
	s := NewUserService(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := s.CreateUser(context.Background(), tc.user)
			if (err != nil) != tc.wantErr {
				t.Errorf("CreateUser(%v) error = %v, wantErr %v", tc.user, err, tc.wantErr)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	testCases := []struct {
		name    string
		userId  int
		wantErr bool
	}{
		{"normal case", 1, false},
		{"already deleted user", 3, true},
		{"non-existing user", 5, true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	db := NewMockRepository(controller)
	s := NewUserService(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.DeleteUser(context.Background(), tc.userId)
			if (err != nil) != tc.wantErr {
				t.Errorf("DeleteUser(%v) error = %v, wantErr %v", tc.userId, err, tc.wantErr)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	testCases := []struct {
		name    string
		user    entities.User
		wantErr bool
	}{
		{"normal case", entities.User{Id: 1, FirstName: "Some", LastName: "User", Email: "admin@example.com", Password: "password"}, false},
		{"deleted user", entities.User{Id: 3, FirstName: "Some", LastName: "User", Email: "admin@example.com", Password: "password"}, true},
		{"non-existing user", entities.User{Id: 5, FirstName: "Some", LastName: "User", Email: "admin@example.com", Password: "password"}, true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	db := NewMockRepository(controller)
	s := NewUserService(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.UpdateUser(context.Background(), tc.user)
			if (err != nil) != tc.wantErr {
				t.Errorf("UpdateUser(%v) error = %v, wantErr %v", tc.user, err, tc.wantErr)
			}
		})
	}
}

func NewMockRepository(controller *gomock.Controller) *mock_repository.MockUserRepository {
	mockDb := mock_repository.NewMockUserRepository(controller)

	var mockUsers = []entities.User{
		entities.User{Id: 1, FirstName: "Admin", LastName: "User", Email: "admin@example.com", Password: "password", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsDeleted: false},
		entities.User{Id: 2, FirstName: "John", LastName: "Doe", Email: "john@doe.com", Password: "string-password", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsDeleted: false},
		entities.User{Id: 3, FirstName: "Jennifer", LastName: "Lawrence", Email: "jen@star.com", Password: "secret", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsDeleted: true},
		entities.User{Id: 4, FirstName: "Rhaenyra", LastName: "Targaryen", Email: "dragon.stone@gmail.com", Password: "ice-and-fire", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsDeleted: false},
	}

	mockDb.EXPECT().GetByID(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, id int) (entities.User, error) {
		if id < 1 || id > 4 {
			return entities.User{}, e.WrapIfErr("failed to get user", errors.New("user not found"))
		}

		if user := mockUsers[id-1]; user.IsDeleted {
			return entities.User{}, e.WrapIfErr("failed to get user", errors.New("user not found"))
		}
		return mockUsers[id-1], nil
	}).AnyTimes()

	mockDb.EXPECT().Create(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, user entities.User) (int, error) {
		return 0, nil
	}).AnyTimes()

	mockDb.EXPECT().Delete(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, id int) error {
		return nil
	}).AnyTimes()

	mockDb.EXPECT().Update(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, user entities.User) error {
		if user.Id < 1 || user.Id > 4 {
			return e.WrapIfErr("failed to update user", errors.New("user not found"))
		}
		return nil
	}).AnyTimes()

	mockDb.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, offset, limit int) ([]entities.User, error) {
		if offset >= len(mockUsers) {
			return []entities.User{}, nil
		}

		if offset+limit >= len(mockUsers) {
			limit = len(mockUsers) - offset
		}

		return mockUsers[offset : offset+limit], nil
	}).AnyTimes()

	return mockDb
}
