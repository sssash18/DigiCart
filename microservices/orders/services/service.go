package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/sssash18/Digicart/pkg/common/database"
	"github.com/sssash18/Digicart/pkg/common/models"
	rabbitmq "github.com/sssash18/Digicart/pkg/common/rabbitmq/producer"
)

func GetOrders(userID string) ([]models.Order, error) {
	if userID == "" {
		return nil, errors.New("empty userID")
	}
	orders := &[]models.Order{}
	db := database.GetDB()
	db.Find(orders, "user_id=?", userID)
	return *orders, nil
}

func GetOrder(userID string, orderID string) (*models.Order, error) {
	if userID == "" || orderID == "" {
		return nil, errors.New("empty userID or orderID")
	}
	order := &models.Order{}
	db := database.GetDB()
	db.Find(order, "user_id=?", userID).Where("order_id=?", orderID)
	return order, nil
}

func CreateOrder(order *models.Order, token string) (*models.Order, error) {
	order.OrderID = uuid.New().String()
	db := database.GetDB()
	tx := db.Create(&order)
	if tx.Error != nil {
		return nil, tx.Error
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"userID":  order.UserID,
		"orderID": order.OrderID,
		"status":  "PENDING",
		"amount":  1000,
	})
	r, _ := http.NewRequest("POST", fmt.Sprintf("%s/payments/create", os.Getenv("PAYMENT_SERVICE_URL")), bytes.NewBuffer(postBody))
	bearer := "Bearer " + token
	r.Header.Set("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	user := models.User{}
	db.Find(&user, "user_id=?", order.UserID)
	rabbitmq.Publish(&models.Message{
		MessageType: "ORDER_PLACED",
		UserID:      order.UserID,
		FirstName:   user.FirstName,
		Email:       user.Email,
	})
	return order, nil
}
