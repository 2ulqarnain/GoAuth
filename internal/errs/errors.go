package errs

import "errors"

var (
	ErrInvalidCreds        = errors.New("invalid credentials")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrUserNotFound        = errors.New("user does not exist")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrUsernameTaken       = errors.New("username not available")
	ErrEmailTaken          = errors.New("email already taken")
	ErrWeakPassword        = errors.New("password is too weak")
	ErrPasswordMismatch    = errors.New("passwords do not match")
	ErrAccountLocked       = errors.New("account is locked")
	ErrAccountDisabled     = errors.New("account is disabled")
	ErrUnverifiedEmail     = errors.New("email unverified")
	ErrUnverifiedPhone     = errors.New("phone number unverified")
	ErrMissingToken        = errors.New("authorization token is missing")
	ErrTokenExpired        = errors.New("authorization token is expired")
	ErrTokenInvalid        = errors.New("authorization token is invalid")
	ErrTokenMalformed      = errors.New("authorization token is malformed")
	ErrSessionExpired      = errors.New("session is expired")
	ErrSessionNotFound     = errors.New("session not found")
	ErrNoPermission        = errors.New("user does not have permission")
	ErrInternalServerError = errors.New("internal server error")
)

// App Errors

var (
	ErrInvalidStoredHash = errors.New("invalid stored hash")
)
