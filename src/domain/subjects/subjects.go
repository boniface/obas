package domain

type MatricSubjects struct {
	SubjectCode string `json:"subjectCode"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Term        string `json:"Term"`
}

type UniversityCourses struct {
	CourseCode  string `json:"courseCode"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"Type"`
	Term        string `json:"Term"`
}
