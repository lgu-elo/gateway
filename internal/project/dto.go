package project

type (
	Project struct {
		ID          int    `json:"id" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}

	ProjectDto struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
)
