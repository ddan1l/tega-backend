package errs

var ValidationFailed = &AppError{
	Status:  422,
	Code:    CodeValidation,
	Message: "Validation error.",
	Details: make(map[string]string),
}

var BadRequest = &AppError{
	Status:  400,
	Code:    CodeBadRequest,
	Message: "Bad request.",
	Details: make(map[string]string),
}

var NotFound = &AppError{
	Status:  404,
	Code:    CodeNotFound,
	Message: "Not found.",
	Details: make(map[string]string),
}
