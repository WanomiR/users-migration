package controller

import (
	"backend/internal/service"
	"backend/internal/utils/readresponder"
	"net/http"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	ListAllUsers(w http.ResponseWriter, r *http.Request)
}

type UserControl struct {
	service       service.UserServicer
	readResponder readresponder.ReadResponder
}

func NewUserControl(service service.UserServicer, readresponder readresponder.ReadResponder) *UserControl {
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

func (u *UserControl) ListAllUsers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
