package validation

import (
	"errors"
	"inventaris/model"
)

func ValidateUser(user *model.User, isLogin bool) error {
    if user.Username == "" {
        return errors.New("username is required")
    }
    if user.Password == "" {
        return errors.New("password is required")
    }
    if !isLogin && user.Role == "" {
        return errors.New("role is required")
    } 
	if !isLogin && user.Role != "admin" && user.Role != "user" {
		return errors.New("role must be 'admin' or 'user'")
	}
	return nil
}