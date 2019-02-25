package model

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
)

// InputValidation - an interface for all "input submission" structs used for
// deserialization.  We pass in the request so that we can potentially get the
// context by the request from our context manager
type InputValidation interface {
	Validate(r *http.Request) error
}

var (
	// ErrInvalidName - error when we have an invalid name
	UserNameReq     = errors.New("user name required")
	UserEmailReq    = errors.New("user email required")
	UserPasswordReq = errors.New("user password required")
)

func (t UserSignUp) Validate(r *http.Request) error {
	// validate the name is not empty or missing
	if !govalidator.IsEmail(t.Email) {
		return UserEmailReq
	}
	if govalidator.IsNull(t.Name) {
		return UserNameReq
	}
	if govalidator.IsNull(t.Password) {
		return UserPasswordReq
	}
	return nil
}

func (t UserSignIn) Validate(r *http.Request) error {
	// validate the name is not empty or missing
	if !govalidator.IsEmail(t.Email) {
		return UserEmailReq
	}
	// validate the password is not empty or missing
	if govalidator.IsNull(t.Password) {
		return UserEmailReq
	}
	return nil
}

func (t UserUpdate) Validate(r *http.Request) error {
	// validate the name is not empty or missing
	if !govalidator.IsEmail(t.Name) {
		return UserEmailReq
	}
	// validate the phone is not empty or missing
	if govalidator.IsNull(t.PhoneNumber) {
		return UserEmailReq
	}
	return nil
}
