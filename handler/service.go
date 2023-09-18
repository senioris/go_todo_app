package handler

import (
	"context"

	"github.com/senioris/go_todo_app/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTaskService AddTaskService
type ListTaskService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type RegisterUesrService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}
