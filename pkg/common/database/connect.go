package database

import (
	"github.com/sssash18/Digicart/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "host=localhost user=suyashchoudhary  dbname=digicart  sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = conn
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.Order{})
}

func GetDB() *gorm.DB {
	return db
}
