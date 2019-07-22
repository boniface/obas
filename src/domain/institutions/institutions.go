package domain

type School struct {
	SchoolId      string `json:"schoolId"`
	SchoolName    string `json:"schoolName"`
	SchoolDetails string `json:"schoolDetails"`
	SchoolState   string `json:"schoolState"`
}

type University struct {
	UniversityId      string `json:"universityId"`
	UniversityName    string `json:"universityName"`
	UniversityDetails string `json:"universityDetails"`
	UniversityState   string `json:"universityState"`
}
