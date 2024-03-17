package models

import (
	"errors"
	"net/mail"
	"net/url"

	"github.com/go-passwd/validator"
)

type User struct {
	ID       uint
	Username string
	Password string
	Address  string
	Phone    string
	Email    string
	UUID     string
}

func NewSignUpUser(form url.Values) (*User, error) {
	// parsing form data
	var (
		name     = form.Get("name")
		email    = form.Get("email")
		password = form.Get("password")
		address  = form.Get("address")
		phone    = form.Get("phone")
	)

	// checking form data existance
	switch {
	case name == "":
		return nil, errors.New("no name provided")
	case email == "":
		return nil, errors.New("no email provided")
	case password == "":
		return nil, errors.New("no password provided")
	case address == "":
		return nil, errors.New("no address provided")
	case phone == "":
		return nil, errors.New("no phone provided")
	}

	// validating email
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, err
	}

	// validating password
	var (
		ERRnotApplicable = errors.New("simple password")
		ERRnotLong       = errors.New("short password")
	)
	validate := validator.New(
		validator.MinLength(8, ERRnotLong),
		validator.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, ERRnotApplicable),
		validator.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 1, ERRnotApplicable),
		validator.ContainsAtLeast("1234567890", 1, ERRnotApplicable),
		validator.ContainsAtLeast("/+-.:;\"'`<>,{}[]()^%$#@!~", 1, ERRnotApplicable),
	)

	err := validate.Validate(password)
	if err != nil {
		return nil, err
	}

	// creating the instance
	user := new(User)
	user.Email = email
	user.Username = name
	user.Password = password
	user.Address = address
	user.Phone = phone

	return user, nil
}
