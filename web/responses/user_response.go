package res

type UserResponse struct {
	ID       int    `json:"id" example:"1"`
	FullName string `json:"fullName" example:"John"`
	Email    string `json:"email" example:"john@john.com"`
}
