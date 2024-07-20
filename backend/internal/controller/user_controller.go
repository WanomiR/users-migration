package controller

import (
	"backend/internal/entities"
	"backend/internal/lib/e"
	"backend/internal/lib/rr"
	"backend/internal/service"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	ListUsers(w http.ResponseWriter, r *http.Request)
}

type UserControl struct {
	service       service.UserServicer
	readResponder rr.ReadResponder
}

func NewUserControl(service service.UserServicer, readresponder rr.ReadResponder) *UserControl {
	return &UserControl{service: service, readResponder: readresponder}
}

// CreateUser godoc
// @Summary create user
// @Description Create new user
// @Tags users
// @Accept json
// @Produce json
// @Param input body entities.User true "user data"
// @Success 201 {object} rr.JSONResponse
// @Failure 400 {object} rr.JSONResponse
// @Router /api/users/0 [post]
func (u *UserControl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := u.readResponder.ReadJSON(w, r, &user); err != nil {
		u.readResponder.WriteJSONError(w, err)
		return
	}

	userID, err := u.service.CreateUser(r.Context(), user)
	if err != nil {
		u.readResponder.WriteJSONError(w, err)
		return
	}

	resp := rr.JSONResponse{
		Message: fmt.Sprintf("user created with id: %d", userID),
	}

	u.readResponder.WriteJSON(w, 201, resp)

}

// GetUserById godoc
// @Summary get user by id
// @Description Return user provided user id
// @Tags users
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} rr.JSONResponse
// @Failure 400 {object} rr.JSONResponse
// @Router /api/users/{id} [get]
func (u *UserControl) GetUserById(w http.ResponseWriter, r *http.Request) {
	// retrieve url param for user id
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	userId, err := strconv.Atoi(id)
	if err != nil {
		u.readResponder.WriteJSONError(w, e.WrapIfErr("error parsing url param", err))
		return
	}

	user, err := u.service.GetUserByID(r.Context(), userId)
	if err != nil {
		u.readResponder.WriteJSONError(w, err)
		return
	}

	resp := rr.JSONResponse{
		Message: "user found",
		Data:    user,
	}

	u.readResponder.WriteJSON(w, http.StatusOK, resp)
}

// UpdateUser godoc
// @Summary update user
// @Description Update user data provided user id
// @Tags users
// @Accept json
// @Produce json
// @Param input body entities.User true "user data"
// @Success 200 {object} rr.JSONResponse
// @Failure 400 {object} rr.JSONResponse
// @Router /api/users/{id} [post]
func (u *UserControl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	if err := u.readResponder.ReadJSON(w, r, &user); err != nil {
		u.readResponder.WriteJSONError(w, err)
		return
	}

	err := u.service.UpdateUser(r.Context(), user)
	if err != nil {
		u.readResponder.WriteJSONError(w, err)
		return
	}

	resp := rr.JSONResponse{
		Message: "user updated",
	}

	u.readResponder.WriteJSON(w, http.StatusOK, resp)
}

// DeleteUser godoc
// @Summary delete user
// @Description Delete user provided user id
// @Tags users
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} rr.JSONResponse
// @Failure 400 {object} rr.JSONResponse
// @Router /api/users/{id} [delete]
func (u *UserControl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// retrieve url param for user id
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	userId, err := strconv.Atoi(id)
	if err != nil {
		u.readResponder.WriteJSONError(w, e.WrapIfErr("error parsing url param", err))
		return
	}

	if err = u.service.DeleteUser(r.Context(), userId); err != nil {
		u.readResponder.WriteJSONError(w, err)
		return
	}

	resp := rr.JSONResponse{
		Message: "user deleted",
	}

	u.readResponder.WriteJSON(w, http.StatusOK, resp)
}

// ListUsers godoc
// @Summary list users
// @Description Return list of users provided limit and offset
// @Tags users
// @Produce json
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Success 200 {object} rr.JSONResponse
// @Failure 400 {object} rr.JSONResponse
// @Router /api/users/ [get]
func (u *UserControl) ListUsers(w http.ResponseWriter, r *http.Request) {
	requestUrl, _ := url.Parse(r.URL.String())

	limit, err := strconv.Atoi(requestUrl.Query().Get("limit"))
	if err != nil {
		u.readResponder.WriteJSONError(w, e.WrapIfErr("error parsing limit", err))
		return
	}

	offset, err := strconv.Atoi(requestUrl.Query().Get("offset"))
	if err != nil {
		u.readResponder.WriteJSONError(w, e.WrapIfErr("error parsing offset", err))
		return
	}

	users, err := u.service.GetUsers(r.Context(), limit, offset)
	if err != nil {
		u.readResponder.WriteJSONError(w, err)
		return
	}

	resp := rr.JSONResponse{
		Message: fmt.Sprintf("%d users found", len(users)),
		Data:    users,
	}

	u.readResponder.WriteJSON(w, 200, resp)
}
