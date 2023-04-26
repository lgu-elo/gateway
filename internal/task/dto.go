package task

type (
	Task struct {
		ID          int    `json:"id" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Status      string `json:"status" validate:"required"`
		Description string `json:"description" validate:"required"`
		User_id     int    `json:"user_id" validate:"required"`
		Project_id  int    `json:"project_id" validate:"required"`
	}

	TaskDto struct {
		Name        string `json:"name" validate:"required"`
		Status      string `json:"status" validate:"required"`
		Description string `json:"description" validate:"required"`
		User_id     int    `json:"user_id" validate:"required"`
		Project_id  int    `json:"project_id" validate:"required"`
	}
)
