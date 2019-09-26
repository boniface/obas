package domain

import "time"

type LogEvent struct {
	Id        string        `json:"id"`
	EventName string        `json:"eventName"`
	EventType string        `json:"eventType"`
	Message   string        `json:"message"`
	Date      time.Location `json:"date"`
}
