package rest

import (
	"ecommerce/config"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) error {
	manager := middleware.NewManager()

	mux := http.NewServeMux()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	address := ":" + strconv.Itoa(int(cnf.HttpPort))

	fmt.Println("Server is running on Port", address)

	err := http.ListenAndServe(address, wrappedMux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	return http.ListenAndServe(address, wrappedMux)
}
