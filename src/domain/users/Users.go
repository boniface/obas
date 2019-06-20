package domain

import (
	"obas/src/domain/location"
	"time"
)

type Student struct {
	Email     string `json:"Email"`
	Firstname string `json:"Firstname"`
	Surname   string `json:"Surname"`
	Identity  string `json:"Identity"`
}

type StudentAddress struct {
	Email           string            `json:"Email"`
	PhysicalAddress string            `json:"PhysicalAddress"`
	City            location.Location `json:"City"`
}

type StudentContacts struct {
}

type StudentDemographics struct {
}

type StudentResults struct {
}

type StudentProfile struct {
}
type StudentApplicationStatus struct {
	Email  string    `json:"Email"`
	Date   time.Time `json:"Date"`
	Status string    `json:"Status"`
}

type ProcessingStatusType struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type StudentDocuments struct {
	Email        string       `json:"Email"`
	DocumentType DocumentType `json:"DocumentType"`
	Description  string       `json:"Description"`
	DocumentUrl  string       `json:"DocumentUrl"`
}

type DocumentType struct {
	Id   string `json:"Id"`
	Type string `json:"Type"`
}

type Admin struct {
	Email string    `json:"Email"`
	Role  AdminRole `json:"Role"`
}

type AdminRole struct {
}
