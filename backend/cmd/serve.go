package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Server is running on http://localhost:8080")

	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
