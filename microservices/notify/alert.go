package main

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"

	"github.com/sssash18/Digicart/pkg/common/models"
)

func GetMessage(msgType string, firstName string) string {
	var msg string
	switch msgType {
	case "ORDER_PLACED":
		msg = fmt.Sprintf("Hey %s\n Congratulations your order has been placed successfully. Please complete the payment.", firstName)
	case "PAYMENT_DONE":
		msg = fmt.Sprintf("Hey %s\n Congratulations your order has been placed successfully. Payment has been received successfully.", firstName)
	}
	return msg
}

func Notifier(msg []byte) error {
	message := models.Message{}
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return err
	}

	from := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("SENDER_PASSWORD")

	to := []string{
		message.Email,
	}
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	messageBody := []byte(GetMessage(message.MessageType, message.FirstName))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, messageBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
