package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sssash18/Digicart/pkg/common/database"
	"github.com/sssash18/Digicart/pkg/common/models"
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

func GetOrder(userID string, orderID string) (*models.Order,error){
	if userID == "" || orderID == "" {
		return nil, errors.New("empty userID or orderID")
	}
	order := &models.Order{}
	db := database.GetDB()
	db.Find(order, "user_id=?", userID).Where("order_id=?",orderID)
	return order, nil
}

func CreateOrder(order *models.Order) (*models.Order,error){
	order.OrderID = uuid.New().String()
	db := database.GetDB()
	tx := db.Create(&order)
	if tx.Error != nil {
		return nil,tx.Error
	}
	
	return order,nil
}