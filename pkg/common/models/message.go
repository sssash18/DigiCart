package models

type Message struct {
	MessageType string `json:"messageType"`
	UserID      string `json:"userID"`
	FirstName   string `json:"firstName"`
	Email       string `json:"email"`
}
