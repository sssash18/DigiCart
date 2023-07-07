package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/sssash18/Digicart/microservices/payments/routes"
)

func main() {
	godotenv.Load()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	routes.SetupRoutes(router)
	fmt.Println("Starting the payments microservice at :8070")
	http.ListenAndServe(":8070", router)
}
