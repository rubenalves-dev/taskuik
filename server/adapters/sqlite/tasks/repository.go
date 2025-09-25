package tasks

import (
	"github.com/rubenalves-dev/taskuik/adapters"
	"github.com/rubenalves-dev/taskuik/internal/resources/task"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client adapters.Database) task.Repository {
	repo := &repository{
		db: client.GetDB(),
	}

	err := repo.db.AutoMigrate(&task.Task{})
	if err != nil {
		panic("Failed to auto-migrate Task model: " + err.Error())
	}

	return repo
}

func (r *repository) Get(id uint) (*task.Task, error) {
	var t task.Task
	if err := r.db.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *repository) List(page, pageSize int) ([]task.Task, int, error) {
	var tasks []task.Task
	var total int64

	if err := r.db.Model(&task.Task{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, int(total), nil
}

func (r *repository) Insert(t *task.Task) error {
	return r.db.Create(t).Error
}

func (r *repository) Update(id uint, updatedTask *task.Task) error {
	var t task.Task
	if err := r.db.First(&t, id).Error; err != nil {
		return err
	}
	t.Title = updatedTask.Title
	t.Status = updatedTask.Status
	return r.db.Save(&t).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&task.Task{}, id).Error
}
