package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID        string `json:"userID"`
	PaymentID     string `json:"paymentID"`
	OrderID       string `json:"orderID"`
	ModeOfPayment string `json:"modeOfPayment"`
	Status        string `json:"status"`
}
