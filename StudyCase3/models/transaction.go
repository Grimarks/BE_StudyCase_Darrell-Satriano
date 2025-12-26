package models

type Transaction struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	EventID  uint
	Quantity int
}
