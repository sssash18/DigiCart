package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID   string  `json:"orderID" gorm:"unique"`
	ProductID string  `json:"productID"`
	UserID    string  `json:"userID"`
	Payment   Payment `json:"payment" gorm:"foreignKey:OrderID;references:OrderID"`
}
