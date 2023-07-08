package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sssash18/Digicart/pkg/common/models"
)

func TestGetMessage(t *testing.T) {
	result := GetMessage("ORDER_PLACED", "adam")
	expected := fmt.Sprintf("Hey %s\n Congratulations your order has been placed successfully. Please complete the payment.", "adam")
	//expected := "Rtrt"
	if result != expected {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expected, result)
	}
}

func TestNotifier(t *testing.T) {
	body, _ := json.Marshal(models.Message{
		MessageType: "ORDER_PLACED",
		UserID:      "1",
		FirstName:   "Adam",
		Email:       "suyashchoudhary42@gmail.com",
	})
	result := Notifier(body)
	
	if result != nil {
		t.Error()
	}
}
