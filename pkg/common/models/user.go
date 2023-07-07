package models

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    string  `json:"userID" gorm:"unique"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Phone     string  `json:"phone"`
	Address   Address `json:"address" gorm:"foreignKey:UserID;references:UserID"`
	Orders    []Order `json:"orders" gorm:"foreignKey:UserID;references:UserID"`
}

type AuthClaims struct {
	Data *User `json:"data"`
	jwt.StandardClaims
}
