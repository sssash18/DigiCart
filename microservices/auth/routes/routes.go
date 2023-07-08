package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sssash18/Digicart/microservices/auth/controller"
)

func SetupRoutes(router *chi.Mux) {
	router.Post("/signup", controller.SignUp)
	router.Post("/login", controller.Login)
}
