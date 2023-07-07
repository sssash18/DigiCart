package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sssash18/Digicart/microservices/orders/services"
	"github.com/sssash18/Digicart/pkg/common/models"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	orders, err := services.GetOrders(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(models.Response{
			Status: "error",
			Err:    err.Error(),
		})
		w.Write(resp)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(models.Response{
		Status: "success",
		Data:   orders,
	})
	w.Write(resp)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	orderID := chi.URLParam(r, "id")
	fmt.Println(orderID)
	order, err := services.GetOrder(userID, orderID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(models.Response{
			Status: "error",
			Err:    err.Error(),
		})
		w.Write(resp)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(models.Response{
		Status: "success",
		Data:   order,
	})
	w.Write(resp)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := &models.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(models.Response{
			Status: "error",
			Err:    err.Error(),
		})
		w.Write(resp)
		return
	}
	order.UserID = r.Header.Get("userID")
	services.CreateOrder(order)
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(models.Response{
		Status: "success",
		Data:   order,
	})
	w.Write(resp)
}
