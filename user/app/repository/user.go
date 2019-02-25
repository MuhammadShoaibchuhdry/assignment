package repository

import (
	"net/http"

	"assignment/user/app/model"
)

type User interface {
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserHandler(w http.ResponseWriter, r *http.Request)
}

type UserDB interface {
	AddUser(user *model.UserSignUp) error
	GetUser(userName, password string, u *model.UserSignUp) error
	UpdateUser(userId string, u *model.UserUpdate) error
	DeleteUser(id string) error
}
