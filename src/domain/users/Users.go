package users

import (
	"obas/src/domain/location"
	"time"
)

type Student struct {
	Email     string
	Firstname string
	Surname   string
	Identity  string
}

type StudentAddress struct {
	Email           string
	PhysicalAddress string
	City            location.Location
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
	Email  string
	Date   time.Time
	Status string
}

type ProcessingStatusType struct {
	Id   string
	Name string
}

type StudentDocuments struct {
	Email        string
	DocumentType DocumentType
	Description  string
	DocumentUrl  string
}

type DocumentType struct {
	Id   string
	Type string
}

type Admin struct {
	Email string
	Role  AdminRole
}

type AdminRole struct {
}
