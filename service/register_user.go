package service

import (
	"context"
	"fmt"

	"github.com/senioris/go_todo_app/entity"
	"github.com/senioris/go_todo_app/store"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUesr struct {
	DB   store.Execer
	Repo UserRegister
}

func (r *RegisterUesr) RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	u := &entity.User{
		Name:     name,
		Password: string(pw),
		Role:     role,
	}

	if err := r.Repo.RegisterUesr(ctx, r.DB, u); err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	return u, nil
}
