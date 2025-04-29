package req

type (
	CreateProjectRequest struct {
		Name        string `json:"name" validate:"required,min=3"`
		Slug        string `json:"slug" validate:"required,min=3"`
		Description string `json:"description"`
	}
)
