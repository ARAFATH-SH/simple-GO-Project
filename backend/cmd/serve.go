package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()
	middlewares := middleware.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo()
	productHandler := product.NewHandler(middlewares, productRepo)

	userRepo := repo.NewUserRepo()
	userHandler := user.NewHandler(middlewares, userRepo, cnf)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()

}
