package errs

import (
	"net/http"
)

var AlreadyExists = &AppError{
	Status:  http.StatusConflict,
	Code:    CodeAlreadyExists,
	Message: "User already exists.",
}

var Auth = &AppError{
	Status:  http.StatusBadRequest,
	Code:    CodeAuthFailed,
	Message: "Failded authorize user.",
}
