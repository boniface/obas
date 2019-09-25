package domain

import "time"

type ApiKeys struct {
	Id     string    `json:"id"`
	Value  string    `json:"value"`
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

type ResetToken struct {
	ResetTokenValue string `json:"resetokenvalue"`
	Email           string `json:"email"`
	Status          string `json:"status"`
}
