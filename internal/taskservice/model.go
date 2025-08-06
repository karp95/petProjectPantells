package taskservice

type Task struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type CreateTaskRequest struct {
	Task   string `json:"task" validate:"required"`
	Status string `json:"status" validate:"required"`
}
