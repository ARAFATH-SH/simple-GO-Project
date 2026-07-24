package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	//domains
	userSvc := user.NewService(userRepo)
	prdSvc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	//handlers
	productHandler := productHandler.NewHandler(middlewares, prdSvc)
	userHandler := userHandler.NewHandler(cnf, userSvc)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()

}
