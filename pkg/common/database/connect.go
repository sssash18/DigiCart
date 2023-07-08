package database

import (
	"fmt"
	"os"

	"github.com/sssash18/Digicart/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	EnvLoad()
	user := os.Getenv("DB_USER")
	dsn := fmt.Sprintf("host=localhost user=%s  dbname=digicart  sslmode=disable TimeZone=Asia/Shanghai", user)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = conn
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Payment{})

}

func GetDB() *gorm.DB {
	return db
}
