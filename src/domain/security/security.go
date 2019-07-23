package domain

import "time"

type ApiKeys struct {
	Id     string    `json:"id"`
	Value  string    `json:"value"`
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

type ResetToken struct {
	ResetokenValue string `json:"resetoken_value"`
	Email          string `json:"email"`
}
