package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sssash18/Digicart/microservices/auth/routes"
	"github.com/sssash18/Digicart/pkg/common/database"
)

func main() {
	database.EnvLoad()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	routes.SetupRoutes(router)
	fmt.Println("Starting the auth microservice at :8080")
	http.ListenAndServe(":8080", router)
}
