package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf            config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(cnf config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler) *Server {

	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}

}

func (server *Server) Start() error {
	manager := middleware.NewManager()

	mux := http.NewServeMux()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	address := ":" + strconv.Itoa(int(server.cnf.HttpPort))

	fmt.Println("Server is running on Port", address)

	err := http.ListenAndServe(address, wrappedMux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	return http.ListenAndServe(address, wrappedMux)
}
