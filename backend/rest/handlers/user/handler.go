package user

import (
	"ecommerce/config"
	"ecommerce/repo"
	middleware "ecommerce/rest/middlewares"
)

type Handler struct {
	cnf         *config.Config
	middlewares *middleware.Middlewares
	userRepo    repo.UserRepo
}

func NewHandler(middlewares *middleware.Middlewares,
	userRepo repo.UserRepo,
	cnf *config.Config) *Handler {
	return &Handler{
		cnf:         cnf,
		middlewares: middlewares,
		userRepo:    userRepo,
	}
}
