package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sssash18/Digicart/microservices/payments/controller"
)

func SetupRoutes(router *chi.Mux) {
	router.Get("/pay",controller.PayOrder)
	router.Get("/payments",controller.Payments)
}
