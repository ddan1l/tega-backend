package user_dto

type (
	CreateUserDto struct {
		FullName     string
		Email        string
		PasswordHash string
	}

	FindByIdDto struct {
		Id int
	}

	FindByEmailDto struct {
		Email string
	}
)
