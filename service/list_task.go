package service

import (
	"context"
	"fmt"

	"github.com/senioris/go_todo_app/entity"
	"github.com/senioris/go_todo_app/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %v", err)
	}
	return ts, nil
}
