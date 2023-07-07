package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sssash18/Digicart/pkg/common/database"
	"github.com/sssash18/Digicart/pkg/common/models"
)

func CreateUser(user *models.User) error {
	db := database.GetDB()
	var count int64
	db.Find(&[]models.User{}).Where("email=?", user.Email).Count(&count)
	db.Find(&[]models.User{}).Where("phone=?", user.Phone).Count(&count)
	if count > 0 {
		return errors.New("user with the given email or phone already exists")
	}
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GenerateJWTToken(user *models.User) (string, error) {
	authClaims := &models.AuthClaims{
		Data: user,
	}
	authClaims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}
