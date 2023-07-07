package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sssash18/Digicart/microservices/orders/controller"
	"github.com/sssash18/Digicart/pkg/common/middleware"
)

func SetupRoutes(router *chi.Mux) {
	router.Use(middleware.Authenticate)
	router.Get("/orders", controller.GetOrders)
	router.Get("/orders/{id}", controller.GetOrder)
	router.Post("/orders/new", controller.CreateOrder)
}
