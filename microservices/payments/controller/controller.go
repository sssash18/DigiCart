package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sssash18/Digicart/microservices/payments/services"
	"github.com/sssash18/Digicart/pkg/common/models"
)

func PayOrder(w http.ResponseWriter, r *http.Request) {
	paymentID := chi.URLParam(r, "id")
	userID := r.Header.Get("userID")
	if paymentID == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(models.Response{
			Status: "error",
			Err:    "Empty Payment ID",
		})
		w.Write(resp)
		return
	}

	err := services.PayOrder(paymentID, userID)
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
	})

	w.Write(resp)

}

func Payments(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	payments, err := services.Payments(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(models.Response{
			Status: "err",
			Err:    err.Error(),
		})
		w.Write(resp)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(models.Response{
		Status: "success",
		Data:   payments,
	})
	w.Write(resp)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	payment := &models.Payment{}
	err := json.NewDecoder(r.Body).Decode(payment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(models.Response{
			Status: "err",
			Err:    err.Error(),
		})
		w.Write(resp)
		return
	}
	err = services.CreatePayment(payment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(models.Response{
			Status: "err",
			Err:    err.Error(),
		})
		w.Write(resp)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(models.Response{
		Status: "success",
	})
	w.Write(resp)
}
