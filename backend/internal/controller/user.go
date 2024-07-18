package controller

import (
	"backend/internal/lib/e"
	"backend/internal/lib/rr"
	"backend/internal/service"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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

func (u *UserControl) CreateUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControl) GetUserById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

// ListUsers godoc
// @Summary list users
// @Description list users provided limit and offset
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
