package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sssash18/Digicart/pkg/common/database"
	"github.com/sssash18/Digicart/pkg/common/models"
	rabbitmq "github.com/sssash18/Digicart/pkg/common/rabbitmq/producer"
)

func PayOrder(paymentID string, userID string) error {
	db := database.GetDB()
	payment := &models.Payment{}
	var count int64
	db.Find(payment, "payment_id=?", paymentID).Count(&count)
	if count == 0 {
		return errors.New("no such payment is pending")
	}
	if payment.UserID != userID {
		return errors.New("user not allowed for this payment")
	}
	payment.Status = "PAID"
	payment.ModeOfPayment = "CASH"
	db.Save(payment)
	user := models.User{}
	db.Find(&user, "user_id=?", payment.UserID)
	rabbitmq.Publish(&models.Message{
		MessageType: "PAYMENT_DONE",
		UserID:      payment.UserID,
		FirstName:   user.FirstName,
		Email:       user.Email,
	})
	return nil
}

func Payments(userID string) ([]models.Payment, error) {
	db := database.GetDB()
	payments := &[]models.Payment{}
	tx := db.Find(payments, "user_id=?", userID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return *payments, nil
}

func CreatePayment(payment *models.Payment) error {
	payment.PaymentID = uuid.New().String()
	db := database.GetDB()
	tx := db.Create(payment)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
