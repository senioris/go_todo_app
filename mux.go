package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/senioris/go_todo_app/auth"
	"github.com/senioris/go_todo_app/clock"
	"github.com/senioris/go_todo_app/config"
	"github.com/senioris/go_todo_app/handler"
	"github.com/senioris/go_todo_app/service"
	"github.com/senioris/go_todo_app/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{
		Service:   &service.AddTask{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/tasks", at.ServeHttp)
	lt := &handler.ListTask{
		Service: &service.ListTask{DB: db, Repo: &r},
	}
	mux.Get("/tasks", lt.ServeHttp)
	ru := &handler.RegisterUser{
		Service:   &service.RegisterUesr{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/register", ru.ServeHttp)

	rcli, err := store.NewKVS(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	jwter, err := auth.NewJWTer(rcli, clock.RealClocker{})
	if err != nil {
		return nil, cleanup, err
	}
	l := &handler.Login{
		Service: &service.Login{
			DB: db, Repo: &r, TokenGenerator: jwter,
		},
		Validator: v,
	}
	mux.Post("/login", l.ServeHttp)
	return mux, cleanup, nil
}
