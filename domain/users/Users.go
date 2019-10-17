package domain

import (
	"time"
)

type User struct {
	Email       string    `json:"email"`
	IdNumber    string    `json:"idNumber"`
	FirstName   string    `json:"firstName"`
	MiddleName  string    `json:"middleName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type UserAddress struct {
	UserId        string `json:"userId"`
	AddressTypeId string `json:"addressTypeId"`
	Address       string `json:"address"`
	PostalCode    string `json:"postalCode"`
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
	UserId        string `json:"userId"`
	ContactTypeId string `json:"contactTypeId"`
	Contact       string `json:"contact"`
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
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRelative struct {
	UserId       string `json:"userId"`
	Name         string `json:"name"`
	Cellphone    string `json:"cellphone"`
	Email        string `json:"email"`
	Relationship string `json:"relationship"`
}

type UserResults struct {
	UserResultsId string `json:"userResultsId"`
	Description   string `json:"description"`
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
