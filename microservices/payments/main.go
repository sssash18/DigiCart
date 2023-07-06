package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sssash18/Digicart/microservices/auth/routes"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	routes.SetupRoutes(router)
	http.ListenAndServe(":8070", router)
}
