package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/sssash18/Digicart/microservices/orders/routes"
)

func main() {
	godotenv.Load()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	routes.SetupRoutes(router)
	fmt.Println("Starting the orders microservice at :8090")
	http.ListenAndServe(":8090", router)
}
