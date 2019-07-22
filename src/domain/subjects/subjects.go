package domain

type MatricSubjects struct {
	SubjectCode string `json:"subjectCode"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Term        string `json:"term"`
}

type UniversityCourses struct {
	CourseCode  string `json:"courseCode"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Term        string `json:"term"`
}
