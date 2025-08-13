package taskservice

import (
	"errors"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("record not found")

type TaskRepository interface {
	GetAll() ([]Task, error)
	Create(task *Task) error
	Delete(id string) error
	Update(id string, updates map[string]interface{}) error
	GetByID(id string) (Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db}
}

func (taskRepo *taskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	err := taskRepo.db.Find(&tasks).Error
	return tasks, err
}

func (taskRepo *taskRepository) Create(task *Task) error {
	return taskRepo.db.Create(task).Error
}

func (taskRepo *taskRepository) Delete(id string) error {
	return taskRepo.db.Delete(&Task{}, "id = ?", id).Error
}

func (taskRepo *taskRepository) Update(id string, updates map[string]interface{}) error {
	return taskRepo.db.Model(&Task{}).Where("id = ?", id).Updates(updates).Error
}

func (taskRepo *taskRepository) GetByID(id string) (Task, error) {
	var task Task
	err := taskRepo.db.First(&task, "id = ?", id).Error
	return task, err
}
