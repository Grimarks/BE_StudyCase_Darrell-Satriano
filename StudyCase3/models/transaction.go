package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	EventID  uint `json:"event_id"`
	Quantity int  `json:"quantity"`
}
