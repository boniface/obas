package domain

type ApplicantType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ApplicationResult struct {
	ApplicationResultId string `json:"applicationResultId"`
	Description         string `json:"description"`
	Date                string `json:"date"`
}

type ApplicationStatus struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Application struct {
	Id                string `json:"id"`
	ApplicationTypeId string `json:"applicationTypeId"`
	ApplicantTypeId   string `json:"applicantTypeId"`
	InstitutionId     string `json:"institutionId"`
	CourseId          string `json:"courseId"`
}
type ApplicationType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
