package models

import (
	"strings"
)

func newClientValidator(um UserModel, servs *Services) *userValidator {
	return &userValidator{
		servs:     servs,
		UserModel: um,
	}
}

type userValidator struct {
	UserModel
	servs      *Services
	validators map[string]map[string][]UserValFunc
}

// defineClientValidators is a function that return a slice of
// validator functions to each field of Client model
func defineClientValidators(uv *userValidator) {
	uv.validators = map[string]map[string][]UserValFunc{}
}

// Validate is used to validate data before save into database
func (uv *userValidator) Validate(user *User, action string) error {
	err := runClientValidation(user, uv, action)
	if err != nil {
		return err
	}
	return nil
}

func runClientValidation(user *User, uv *userValidator, action string) error {
	for _, valueField := range uv.validators {
		for keyAction, valueAction := range valueField {
			if strings.Contains(keyAction, action) {
				err := runClientValFuncs(user, valueAction...)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UserValFunc func(*User) error

func runClientValFuncs(user *User, fns ...UserValFunc) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}
