package taskservice

type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type CreateTaskRequest struct {
	Task   string `json:"task" validate:"required"`
	IsDone bool   `json:"is_done"`
}
