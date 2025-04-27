package user_dto

type (
	CreateUserDto struct {
		FullName     string
		Email        string
		PasswordHash string
	}

	FindByIdDto struct {
		ID int
	}

	FindByEmailDto struct {
		Email string
	}

	ProjectDto struct {
		ID   int
		Name string
		Slug string
	}

	ProjectsDto struct {
		Projects []ProjectDto
	}
)
