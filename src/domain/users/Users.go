package domain

import (
	"time"
)

type Student struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
	Identity  string `json:"identity"`
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

type Users struct {
	Email       string    `json:"email"`
	FirstName   string    `json:"firstName"`
	MiddleName  string    `json:"middleName"`
	lastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

// i found them in the obasapi domain so i added them but i'm not sure if i should

type UserAddress struct {
	UserAddressId   string `json:"userAddressId"`
	PhysicalAddress string `json:"physicalAddress"`
	PostalCode      string `json:"postalCode"`
}

type UserApplicationResult struct {
	UserApplicationResultId string `json:"userApplicationResultId"`
	Description             string `json:"description"`
}

type UserCommunication struct {
	CommunicationId string `json:"communicationId"`
	Description     string `json:"description"`
}

type UserContacts struct {
	UserContactId     string `json:"userContactId"`
	CellNumber        string `json:"cellNumber"`
	AlternativeNumber string `json:"alternativeNumber"`
	AlternativeEmail  string `json:"alternativeEmail"`
}

type UserDemographics struct {
	UserDemographicsId string `json:"userDemographicsId"`
	GenderId           string `json:"genderId"`
	RaceId             string `json:"raceId"`
}

type UserDocuments struct {
	UserDocumentsId string `json:"userDocumentsId"`
	DocumentId      string `json:"documentId"`
}

type UserInstitution struct {
	UserInstitutionId string `json:"userInstitutionId"`
	Name              string `json:"name"`
}

type UserPassword struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

type UserRelative struct {
	UserRelativeId string `json:"userRelativeId"`
	Name           string `json:"name"`
	Cellphone      string `json:"cellphone"`
	Relationship   string `json:"relationship"`
	Email          string `json:"email"`
}

type UserResult struct {
	UserResultId string `json:"userResultId"`
	Description  string `json:"description"`
}

type UserRole struct {
	UserId string `json:"userId"`
	RoleId string `json:"roleId"`
}

type UserSubjects struct {
	UserSubjectId string `json:"userSubjectId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Term          string `json:"term"`
}
