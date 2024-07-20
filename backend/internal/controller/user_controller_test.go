package controller

import (
	"backend/internal/controller/mock_service"
	"backend/internal/entities"
	"backend/internal/lib/rr"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//go:generate mockgen -source=../service/user_service.go -destination=./mock_service/mock_service.go

func TestUserControl_CreateUser(t *testing.T) {
	testCases := []struct {
		name       string
		body       any
		wantStatus int
	}{
		{"normal case", entities.User{Id: 1, FirstName: "Rhaenyra", LastName: "Targaryen", Email: "dragon.stone@gmail.com", Password: "ice-and-fire"}, 201},
		{"empty body", nil, 400},
		{"too short user name", entities.User{FirstName: "L", LastName: "Cxi"}, 400},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockService(controller)
	tgc := NewUserControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var body bytes.Buffer
			json.NewEncoder(&body).Encode(tc.body)

			req := httptest.NewRequest(http.MethodPost, "/api/users/0", &body)
			wr := httptest.NewRecorder()

			tgc.CreateUser(wr, req)

			r := wr.Result()

			var resp rr.JSONResponse
			json.NewDecoder(r.Body).Decode(&resp)
			defer r.Body.Close()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("CreateUser(), expected status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestUserControl_GetUserById(t *testing.T) {
	testCases := []struct {
		name       string
		userId     any
		wantStatus int
	}{
		{"existing user", 1, 200},
		{"non-existing user", 5, 400},
		{"bad url param", "aklf", 400},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockService(controller)
	tgc := NewUserControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/users/"+fmt.Sprintf("%v", tc.userId), nil)
			wr := httptest.NewRecorder()

			tgc.GetUserById(wr, req)

			r := wr.Result()

			var resp rr.JSONResponse
			json.NewDecoder(r.Body).Decode(&resp)
			defer r.Body.Close()

			if r.StatusCode != tc.wantStatus {
				fmt.Println(resp.Message)
				t.Errorf("GetUserById(), expected status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestUserControl_UpdateUser(t *testing.T) {
	testCases := []struct {
		name       string
		body       any
		wantStatus int
	}{
		{"normal case", entities.User{Id: 1, FirstName: "Rhaenyra", LastName: "Targaryen", Email: "dragon.stone@gmail.com", Password: "ice-and-fire"}, 200},
		{"empty body", nil, 400},
		{"non-existing user", entities.User{Id: 5, FirstName: "Rhaenyra", LastName: "Targaryen", Email: "dragon.stone@gmail.com", Password: "ice-and-fire"}, 400},
		{"too short user name", entities.User{FirstName: "L", LastName: "Cxi"}, 400},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockService(controller)
	tgc := NewUserControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var body bytes.Buffer
			json.NewEncoder(&body).Encode(tc.body)

			req := httptest.NewRequest(http.MethodPost, "/api/users/1", &body)
			wr := httptest.NewRecorder()

			tgc.UpdateUser(wr, req)

			r := wr.Result()

			var resp rr.JSONResponse
			json.NewDecoder(r.Body).Decode(&resp)
			defer r.Body.Close()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("UpdateUser(), expected status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestUserControl_DeleteUser(t *testing.T) {
	testCases := []struct {
		name       string
		userId     any
		wantStatus int
	}{
		{"existing user", 1, 200},
		{"non-existing user", 5, 400},
		{"bad url param", "aog", 400},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockService(controller)
	tgc := NewUserControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/api/users/"+fmt.Sprintf("%v", tc.userId), nil)
			wr := httptest.NewRecorder()

			tgc.DeleteUser(wr, req)

			r := wr.Result()

			var resp rr.JSONResponse
			json.NewDecoder(r.Body).Decode(&resp)
			defer r.Body.Close()

			if r.StatusCode != tc.wantStatus {
				fmt.Println(resp.Message)
				t.Errorf("DeleteUser(), expected status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestUserControl_ListUsers(t *testing.T) {
	type params struct {
		offset, limit any
	}
	testCases := []struct {
		name       string
		params     params
		wantStatus int
	}{
		{"normal case", params{0, 1}, 200},
		{"bad offset param", params{"q068", 1}, 400},
		{"bad limit param", params{1, "?"}, 400},
		{"bigger limit", params{0, 3}, 200},
		{"too small limit", params{0, 0}, 400},
		{"offset and limit on the edge", params{2, 3}, 200},
		{"offset and limit outside the range", params{4, 5}, 200},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockService(controller)
	tgc := NewUserControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			urlPath := "/api/users/" + "?limit=" + fmt.Sprintf("%v", tc.params.limit) + "&offset=" + fmt.Sprintf("%v", tc.params.offset)

			req := httptest.NewRequest(http.MethodGet, urlPath, nil)
			wr := httptest.NewRecorder()

			tgc.ListUsers(wr, req)

			r := wr.Result()

			var resp rr.JSONResponse
			json.NewDecoder(r.Body).Decode(&resp)
			defer r.Body.Close()

			if r.StatusCode != tc.wantStatus {
				fmt.Println(resp.Message)
				t.Errorf("ListUsers(), expected status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}

}

func NewMockService(controller *gomock.Controller) *mock_service.MockUserServicer {
	mockService := mock_service.NewMockUserServicer(controller)

	var mockUsers = []entities.User{
		entities.User{Id: 1, FirstName: "Admin", LastName: "User", Email: "admin@example.com", Password: "password", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsDeleted: false},
		entities.User{Id: 2, FirstName: "John", LastName: "Doe", Email: "john@doe.com", Password: "string-password", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsDeleted: false},
		entities.User{Id: 3, FirstName: "Jennifer", LastName: "Lawrence", Email: "jen@star.com", Password: "secret", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsDeleted: true},
		entities.User{Id: 4, FirstName: "Rhaenyra", LastName: "Targaryen", Email: "dragon.stone@gmail.com", Password: "ice-and-fire", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsDeleted: false},
	}

	mockService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, user entities.User) (int, error) {
		if len(user.FirstName) < 2 || len(user.LastName) < 2 {
			return 0, errors.New("first and last name must be at least 2 characters")
		}
		return 1, nil
	}).AnyTimes()

	mockService.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, id int) (entities.User, error) {
		if id < 1 || id > 4 {
			return entities.User{}, errors.New("user not found")
		}
		return mockUsers[id-1], nil
	}).AnyTimes()

	mockService.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, user entities.User) error {
		if user.Id < 1 || user.Id > 4 {
			return errors.New("user not found")
		} else if len(user.FirstName) < 2 || len(user.LastName) < 2 {
			return errors.New("first and last name must be at least 2 characters")
		}
		return nil
	}).AnyTimes()

	mockService.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, id int) error {
		if id < 1 || id > 4 {
			return errors.New("user not found")
		}
		return nil
	}).AnyTimes()

	mockService.EXPECT().GetUsers(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, offset, limit int) ([]entities.User, error) {
		if offset < 0 || limit < 1 {
			return []entities.User{}, errors.New("limit must be greater than zero")
		}

		if offset >= len(mockUsers) {
			return []entities.User{}, nil
		}

		if offset+limit >= len(mockUsers) {
			limit = len(mockUsers) - offset
		}

		return mockUsers[offset : offset+limit], nil
	}).AnyTimes()

	return mockService
}
