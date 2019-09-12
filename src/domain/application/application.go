package domain

type ApplicationType struct {
	ApplicantTypeId string `json:"applicantTypeId"`
	Name            string `json:"name"`
}

type ApplicationResult struct {
	ApplicationResultId string `json:"applicationResultId"`
	Description         string `json:"description"`
	Date                string `json:"date"`
}

type ApplicationStatus struct {
	ApplicationStatusId string `json:"applicationStatusId"`
	Description         string `json:"description"`
	Date                string `json:"date"`
}
