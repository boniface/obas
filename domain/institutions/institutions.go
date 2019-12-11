package domain

type School struct {
	SchoolId          string `json:"schoolId"`
	SchoolName        string `json:"schoolName"`
	SchoolProvince    string `json:"schoolProvince"`
	SchoolAddress     string `json:"schoolAddress"`
	SchoolPhonenumber string `json:"schoolPhonenumber"`
}

type University struct {
	UniversityId          string `json:"universityId"`
	UniversityName        string `json:"universityName"`
	UniversityProvince    string `json:"universityProvince"`
	UniversityPhoneNumber string `json:"universityPhoneNumber"`
	UniversityEmail       string `json:"universityEmail"`
}
type Institution struct {
	Id                string `json:"id"`
	InstitutionTypeId string `json:"institutionTypeId"`
	Name              string `json:"name"`
}
type InstitutionTypes struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type InstitutionCourse struct {
	InstitutionId string `json:"institutionId"`
	CourseId      string `json:"courseId"`
}
type InstitutionLocation struct {
	InstitutionId string `json:"institutionId"`
	LocationId    string `json:"locationId"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
}
type InstitutionAddress struct {
	InstitutionId string `json:"institutionId"`
	AddressTypeId string `json:"addressTypeId"`
	Address       string `json:"address"`
	PostalCode    string `json:"postalCode"`
}
