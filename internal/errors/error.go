package errors

import "errors"

// Create your custom error
var (
	// ErrSampleCustomError example custom error
	ErrSampleCustomError = errors.New("example custom error")

	// ErrWrongUsernameOrPassword when user try to login with wrong username or password
	ErrWrongUsernameOrPassword = errors.New("wrong username or password")

	// ErrUsernameAlreadyExist when user register with registered username
	ErrUsernameAlreadyExist = errors.New("username is already exist")
)
