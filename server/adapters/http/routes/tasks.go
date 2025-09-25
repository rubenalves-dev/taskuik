package routes

import (
	"github.com/rubenalves-dev/taskuik/adapters/sqlite/tasks"
	"github.com/rubenalves-dev/taskuik/internal/resources/task"
)

func (r *Router) RegisterTaskRoutes() {
	if r.appDB == nil {
		panic("Database client is not set for Task routes")
	}

	repository := tasks.NewRepository(r.appDB)
	service := task.NewService(repository)
	handler := task.NewHandler(service)

	handler.Handle(r.Engine)
}
