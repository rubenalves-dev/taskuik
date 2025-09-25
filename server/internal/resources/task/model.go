package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title  string
	Status TaskStatus
}

type TaskStatus int

const (
	TODO TaskStatus = iota
	IN_PROGRESS
	DONE
)
