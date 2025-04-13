package errs

import (
	"net/http"
)

var AlreadyExists = &AppError{
	Status:  http.StatusConflict,
	Code:    CodeAlreadyExists,
	Message: "User already exists.",
}

var IncorrectPassword = &AppError{
	Status:  http.StatusUnauthorized,
	Code:    CodeIncorrectPassword,
	Message: "Incorrect password.",
}

var UserNotFound = &AppError{
	Status:  http.StatusBadRequest,
	Code:    CodeUserNotFound,
	Message: "User not found.",
}

var Unauthorized = &AppError{
	Status:  http.StatusUnauthorized,
	Code:    CodeUnauthorized,
	Message: "Unauthorized.",
}

var Forbidden = &AppError{
	Status:  http.StatusForbidden,
	Code:    CodeForbidden,
	Message: "Forbidden.",
}

var TokenExpired = &AppError{
	Status:  http.StatusForbidden,
	Code:    CodeTokenExpired,
	Message: "Token expired.",
}

var Auth = &AppError{
	Status:  http.StatusBadRequest,
	Code:    CodeAuthFailed,
	Message: "Failded authorize user.",
}
