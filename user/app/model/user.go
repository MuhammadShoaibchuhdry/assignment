package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// UserSignUp represents the structure of our resource
	UserSignUp struct {
		Id               bson.ObjectId `json:"id" bson:"_id"`
		Name             string        `json:"fullname"`
		Email            string        `json:"email"`
		Password         string        `json:"password"`
		PhoneNumber      string        `json:"phonenumber"`
	}
	// User Update
	UserUpdate struct {
		Name             string        `json:"fullname"`
		PhoneNumber      string        `json:"phonenumber"`
	}
	// UserSignIn
	UserSignIn struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// UserResponse
	UserResponse struct {
		AccessToken   string   `json:"accessToken"`
		Name          string   `json:"name"`
		Email         string   `json:"email"`
	}
	Error struct {
		Error string `json:"error"`
	}
	Success struct {
		Message string `json:"message"`
	}
)
