package taskservice

type TaskService interface {
	GetTasks() ([]Task, error)
	AddTask(req CreateTaskRequest) (Task, error)
	DeleteTask(id string) error
	PatchTask(id string, updates map[string]interface{}) (Task, error)
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	return &taskService{
		repo: repo}
}

func (service *taskService) GetTasks() ([]Task, error) {
	return service.repo.GetAll()
}

func (service *taskService) AddTask(req CreateTaskRequest) (Task, error) {
	task := Task{
		Task:   req.Task,
		IsDone: req.IsDone,
	}
	err := service.repo.Create(&task)
	return task, err
}

func (service *taskService) DeleteTask(id string) error {
	return service.repo.Delete(id)
}

func (service *taskService) PatchTask(id string, updates map[string]interface{}) (Task, error) {
	err := service.repo.Update(id, updates)
	if err != nil {
		return Task{}, err
	}
	return service.repo.GetByID(id)
}
