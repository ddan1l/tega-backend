package errs

var AlreadyExists = &AppError{
	Status:  409,
	Code:    CodeAlreadyExists,
	Message: "User already exists.",
}

var IncorrectPassword = &AppError{
	Status:  401,
	Code:    CodeIncorrectPassword,
	Message: "Incorrect password.",
}

var UserNotFound = &AppError{
	Status:  400,
	Code:    CodeUserNotFound,
	Message: "User not found.",
}

var Unauthorized = &AppError{
	Status:  401,
	Code:    CodeUnauthorized,
	Message: "Unauthorized.",
}

var Forbidden = &AppError{
	Status:  403,
	Code:    CodeForbidden,
	Message: "Forbidden.",
}

var TokenExpired = &AppError{
	Status:  403,
	Code:    CodeTokenExpired,
	Message: "Token expired.",
}

var Auth = &AppError{
	Status:  400,
	Code:    CodeAuthFailed,
	Message: "Failded authorize user.",
}
