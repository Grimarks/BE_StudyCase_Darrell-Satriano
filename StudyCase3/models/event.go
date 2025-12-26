package models

type Event struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Capacity    int
	TicketsSold int
}
