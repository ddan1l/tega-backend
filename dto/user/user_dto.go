package user_dto

type (
	CreateUserDto struct {
		FullName     string
		Email        string
		PasswordHash string
	}

	UserrDto struct {
		ID       int    `json:"id" example:"1"`
		FullName string `json:"fullName" example:"John"`
		Email    string `json:"email" example:"john@john.com"`
	}

	FindByIdDto struct {
		ID int
	}

	FindByEmailDto struct {
		Email string
	}
)
