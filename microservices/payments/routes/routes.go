package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sssash18/Digicart/microservices/payments/controller"
	"github.com/sssash18/Digicart/pkg/common/middleware"
)

func SetupRoutes(router *chi.Mux) {
	router.Use(middleware.Authenticate)
	router.Get("/pay/{id}", controller.PayOrder)
	router.Get("/payments", controller.Payments)
	router.Post("/payments/create", controller.CreatePayment)
}
