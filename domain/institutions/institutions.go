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
