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
	UserId   string `json:"userId"`
	TitleId  string `json:"titleId"`
	GenderId string `json:"genderId"`
	RaceId   string `json:"raceId"`
}

type UserDocument struct {
	UserId     string `json:"userId"`
	DocumentId string `json:"documentId"`
}

type UserInstitution struct {
	InstitutionId string `json:"institutionId"`
	UserId        string `json:"userId"`
	Current       bool   `json:"current"`
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
	UserId        string `json:"userId"`
	SubjectId     string `json:"subjectId"`
	Marks         string `json:"marks"`
	InstitutionId string `json:"institutionId"`
}

type UserTown struct {
	UserId   string `json:"userId"`
	TownCode string `json:"townCode"`
}
type UserCourse struct {
	UserId        string `json:"UserId"`
	CourseId      string `json:"courseId"`
	InstitutionId string `json:"institutionId"`
}
type UserApplication struct {
	UserId        string `json:"userId"`
	ApplicationId string `json:"applicationId"`
}
type UserApplicationStatus struct {
	ApplicationId string    `json:"applicationId"`
	StatusId      string    `json:"statusId"`
	DateTime      time.Time `json:"dateTime"`
	Modifier      string    `json:"modifier"`
}
