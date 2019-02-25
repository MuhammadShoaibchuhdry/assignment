package controller

import (
	"fmt"
	"net/http"
	"time"

	"assignment/user/app/model"
	"assignment/user/app/repository"
	"assignment/user/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

type user struct {
	db repository.UserDB
}

//NewCouchRepository create new repository
func NewUserRepository(ub repository.UserDB) repository.User {
	return &user{db: ub}
}

// CreateUserHandler will create new user in couchbase db_conn
func (u *user) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var data model.UserSignUp
	if err := model.DecodeAndValidate(r, &data); err != nil {
		utils.WriteJsonErr(w, err)
		return
	}
	// trim email
	data.Email = utils.TrimEmailAddress(data.Email)
	if err := u.db.AddUser(&data); err != nil {
		utils.WriteJsonErr(w, err)
		return
	}

	utils.WriteJsonRes(w, data)

}

// GetUserHandler will return user name and token id
func (u *user) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var data model.UserSignIn
	var user model.UserSignUp

	if err := model.DecodeAndValidate(r, &data); err != nil {
		utils.WriteJsonErr(w, err)
		return
	}
	data.Email = utils.TrimEmailAddress(data.Email)
	if err := u.db.GetUser(data.Email, data.Password, &user); err != nil {
		utils.WriteJsonErr(w, err)
		return
	}

	tokenString, err := utils.GenerateUserToken(jwt.MapClaims{
		"user_id": user.Id.Hex(),
		"exp":     time.Now().Add(time.Hour * 3).Unix(),
	})
	if err != nil {
		utils.WriteJsonErr(w, err)
		return
	}

	res := model.UserResponse{
		AccessToken:   tokenString,
		Name:          user.Name,
		Email:         user.Email,
	}

	utils.WriteJsonRes(w, res)

}

// UpdateUserHandler will update user name and phone
func (u *user) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var data model.UserUpdate
	_, claims, _ := jwtauth.FromContext(r.Context())
	id := claims["user_id"]
	if id == nil {
		utils.WriteJsonErr(w, fmt.Errorf("token id not found in request"))
		return
	}
	if err := model.DecodeAndValidate(r, &data); err != nil {
		utils.WriteJsonErr(w, err)
		return
	}
	if err := u.db.UpdateUser(id.(string), &data); err != nil {
		utils.WriteJsonErr(w, err)
		return
	}
	utils.WriteJsonRes(w, model.Success{Message: "success"})
}

// DeleteUserHandler will delete user from db
func (u *user) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	id := claims["user_id"]
	if id == nil {
		utils.WriteJsonErr(w, fmt.Errorf("token id not found in request"))
		return
	}

	if err := u.db.DeleteUser(id.(string)); err != nil {
		utils.WriteJsonErr(w, err)
		return
	}
	utils.WriteJsonRes(w, model.Success{Message: "success"})
}
