package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server is running on :8080")

	err := http.ListenAndServe(":8080", wrappedMux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
