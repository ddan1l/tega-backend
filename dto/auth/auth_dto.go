package auth_dto

type (
	CreateTokenDto struct {
		UserId int
		Token  string
	}

	TokensPairDto struct {
		AccessToken  string
		RefreshToken string
	}

	RegisterUserDto struct {
		FullName string
		Email    string
		Password string
	}
	LoginUserDto struct {
		Email    string
		Password string
	}
)
