package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID  string `json:"userID"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}
