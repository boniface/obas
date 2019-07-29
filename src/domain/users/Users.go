package domain

import (
	"time"
)

type User struct {
	Email       string    `json:"email"`
	FirstName   string    `json:"firstName"`
	MiddleName  string    `json:"middleName"`
	LastName    string    `json:"lastName"`
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

type UserResults struct {
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
