package entity

import "time"

// Fruit is the representation of a Fruit object used to hold the business logic.
type Fruit struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Color          string    `json:"color"`
	Country        string    `json:"country"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Fruits is a list of Fruit elements
type Fruits []Fruit
