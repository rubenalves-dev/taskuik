package task

import (
	"net/http"

	"github.com/rubenalves-dev/taskuik/internal/common"
)

var _ Service = (*service)(nil)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateTask(task *Task) common.StatusErr {
	if err := s.repo.Insert(task); err != nil {
		return common.ErrStatus(err, http.StatusInternalServerError)
	}
	return common.OkStatus(http.StatusCreated)
}

func (s *service) ListTasks(page, pageSize int) common.ListResult[Task] {
	tasks, total, err := s.repo.List(page, pageSize)
	if err != nil {
		return common.ErrListResult[Task](err, http.StatusInternalServerError)
	}
	return common.OkListResult(tasks, total, http.StatusOK)
}

func (s *service) GetTask(id uint) common.Result[*Task] {
	task, err := s.repo.Get(id)
	if err != nil {
		return common.ErrResult[*Task](err, http.StatusInternalServerError)
	}
	return common.OkResult(task, http.StatusOK)
}

func (s *service) UpdateTask(task *Task) common.StatusErr {
	if err := s.repo.Update(task.ID, task); err != nil {
		return common.ErrStatus(err, http.StatusInternalServerError)
	}
	return common.OkStatus(http.StatusOK)
}

func (s *service) DeleteTask(id uint) common.StatusErr {
	if err := s.repo.Delete(id); err != nil {
		return common.ErrStatus(err, http.StatusInternalServerError)
	}
	return common.OkStatus(http.StatusNoContent)
}
