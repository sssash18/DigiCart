package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID        string `json:"userID"`
	PaymentID     string `json:"paymentID" gorm:"unique"`
	OrderID       string `json:"orderID"`
	ModeOfPayment string `json:"modeOfPayment"`
	Status        string `json:"status"`
	Amount        int64  `json:"amount"`
}
