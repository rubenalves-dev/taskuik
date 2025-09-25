package task

import "github.com/rubenalves-dev/taskuik/internal/common"

type Service interface {
	CreateTask(task *Task) common.StatusErr
	ListTasks(page, pageSize int) common.ListResult[Task]
	GetTask(id uint) common.Result[*Task]
	UpdateTask(task *Task) common.StatusErr
	DeleteTask(id uint) common.StatusErr
}

type Repository interface {
	Get(id uint) (*Task, error)
	List(page, pageSize int) ([]Task, int, error)
	Insert(task *Task) error
	Update(id uint, updatedTask *Task) error
	Delete(id uint) error
}
