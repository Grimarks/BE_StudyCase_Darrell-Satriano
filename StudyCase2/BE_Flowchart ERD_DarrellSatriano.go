package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Capacity    int    `json:"capacity"`
	TicketsSold int    `json:"tickets_sold"`
}

type Transaction struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	EventID  int `json:"event_id"`
	Quantity int `json:"quantity"`
}

var users []User
var events []Event
var transactions []Transaction

// ===== USER =====
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = len(users) + 1
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// ===== EVENT =====
func createEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	json.NewDecoder(r.Body).Decode(&event)
	event.ID = len(events) + 1
	event.TicketsSold = 0
	events = append(events, event)
	json.NewEncoder(w).Encode(event)
}

// ===== TRANSACTION =====
func createTransaction(w http.ResponseWriter, r *http.Request) {
	var trx Transaction
	json.NewDecoder(r.Body).Decode(&trx)

	for i := range events {
		if events[i].ID == trx.EventID {
			// CEK OVERSELLING
			if events[i].TicketsSold+trx.Quantity > events[i].Capacity {
				http.Error(w, "Tiket tidak mencukupi", http.StatusBadRequest)
				return
			}

			events[i].TicketsSold += trx.Quantity
			trx.ID = len(transactions) + 1
			transactions = append(transactions, trx)

			json.NewEncoder(w).Encode(trx)
			return
		}
	}

	http.Error(w, "Event tidak ditemukan", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/users", createUser)
	http.HandleFunc("/events", createEvent)
	http.HandleFunc("/transactions", createTransaction)

	fmt.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)
}
