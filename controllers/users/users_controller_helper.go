package controllers

import (
	"net/http"
	"obas/config"
	academicsDomain "obas/domain/academics"
	institutionDomain "obas/domain/institutions"
	locationDomain "obas/domain/location"
	userDomain "obas/domain/users"
	academicsIO "obas/io/academics"
	demographyIO "obas/io/demographics"
	institutionIO "obas/io/institutions"
	usersIO "obas/io/users"
	"obas/util"
	"strings"
)

type ExtendedUserMatricSubject struct {
	userDomain.UserMatricSubject
	SubjectName string
}

func getProvinces(app *config.Env) ([]locationDomain.Location, PageToast) {
	var provinces []locationDomain.Location
	var alert PageToast
	provinces, err := util.GetProvinces()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = PageToast{dangerAlertStyle, "Could not retrieve provinces!"}
	}
	return provinces, alert
}

func getInstitutionTypes(app *config.Env) ([]institutionDomain.InstitutionTypes, PageToast) {
	var alert PageToast
	var institutionTypes []institutionDomain.InstitutionTypes
	institutionTypes, err := institutionIO.GetInstitutionTypes()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = PageToast{dangerAlertStyle, "Could not retrieve institution types!"}
	}
	return institutionTypes, alert
}

/**
Get user matric institution given user id
*/
func getUserMatricInstitution(app *config.Env, userId string) (userDomain.UserMatricInstitution, PageToast) {
	var userMatricInstitution userDomain.UserMatricInstitution
	var alert PageToast
	userMatricInstitution, err := usersIO.GetUserMatricInstitution(userId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = PageToast{dangerAlertStyle, "Could not retrieve user matric institution!"}
	}
	return userMatricInstitution, alert
}

/**
Get Matric institution name for user
*/
func getUserMatricInstitutionName(app *config.Env, institutionId string) (string, PageToast) {
	var institutionName string
	var institution institutionDomain.Institution
	var alert PageToast
	institution, alert = getInstitution(app, institutionId)
	if alert.AlertInfo == "" {
		if institution.Id == "" {
			institutionName = "<< NOT SET >>"
		} else {
			institutionName = institution.Name
		}
	}
	return institutionName, alert
}

/**
Get institution given institution id
*/
func getInstitution(app *config.Env, institutionId string) (institutionDomain.Institution, PageToast) {
	var institution institutionDomain.Institution
	var alert PageToast
	institution, err := institutionIO.GetInstitution(institutionId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = PageToast{dangerAlertStyle, "Could not retrieve institution!"}
	}
	return institution, alert
}

/**
Check if session has alert message
*/
func checkForSessionAlert(app *config.Env, r *http.Request) PageToast {
	message := app.Session.GetString(r.Context(), "message")
	messageType := app.Session.GetString(r.Context(), "message-type")
	var alert PageToast
	if message != "" && messageType != "" {
		alert = PageToast{messageType, message}
		app.Session.Remove(r.Context(), "message")
		app.Session.Remove(r.Context(), "message-type")
	}
	return alert
}

/**
Get date in YYYYMMDD format
*/
func getDate_YYYYMMDD(dateString string) string {
	return strings.Split(dateString, " ")[0]
}

/**
Filter user race from list of race given user demography
*/
func getUserRace(demography usersIO.UserDemography, races []demographyIO.Race) demographyIO.Race {
	for _, race := range races {
		if demography.RaceId == race.RaceId {
			return race
		}
	}
	return demographyIO.Race{}
}

/**
Filter user title form list of titles given user demography
*/
func getUserTitle(demography usersIO.UserDemography, titles []demographyIO.Title) demographyIO.Title {
	for _, title := range titles {
		if demography.TitleId == title.TitleId {
			return title
		}
	}
	return demographyIO.Title{}
}

/**
Filter user gender from list of genders given user demography
*/
func getUserGender(demography usersIO.UserDemography, genders []demographyIO.Gender) demographyIO.Gender {
	for _, gender := range genders {
		if demography.GenderId == gender.GenderId {
			return gender
		}
	}
	return demographyIO.Gender{}
}

/**
Get Matric subjects given institution id
*/
func getMatricSubjects(app *config.Env, institutionId string) ([]academicsDomain.Subject, PageToast) {
	var subjects []academicsDomain.Subject
	var courseSubjects []academicsDomain.CourseSubject
	var alert PageToast
	institutionCourses, alert := getInstitutionCourses(app, institutionId)
	if alert.AlertInfo == "" && len(institutionCourses) > 0 {
		institutionCourse := institutionCourses[0]
		courseSubjects, alert = getSubjectsForCourse(app, institutionCourse.CourseId)
	}
	if alert.AlertInfo == "" {
		for _, courseSubject := range courseSubjects {
			subject, alert := getSubject(app, courseSubject.SubjectId)
			if alert.AlertInfo == "" {
				subjects = append(subjects, subject)
			}
		}
	}
	return subjects, alert
}

/**
Get Subject details given subject id
 */
func getSubject(app *config.Env, subjectId string) (academicsDomain.Subject, PageToast) {
	var subject academicsDomain.Subject
	var alert PageToast
	subject, err := academicsIO.GetSubject(subjectId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = PageToast{dangerAlertStyle, "Could not retrieve subject!"}
	}
	return subject, alert
}

/**
Get subjects for course given course id
 */
func getSubjectsForCourse(app *config.Env, courseId string) ([]academicsDomain.CourseSubject, PageToast) {
	var courseSubjects []academicsDomain.CourseSubject
	var alert PageToast
	courseSubjects, err := academicsIO.GetCourseSubjects(courseId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = PageToast{dangerAlertStyle, "Could not retrieve subjects for course!"}
	}
	return courseSubjects, alert
}

/**
Get courses for institution given institution id
 */
func getInstitutionCourses(app *config.Env, institutionId string) ([]institutionDomain.InstitutionCourse, PageToast) {
	var institutionCourses []institutionDomain.InstitutionCourse
	var alert PageToast
	institutionCourses, err := institutionIO.GetInstitutionCourses(institutionId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = PageToast{dangerAlertStyle, "Could not retrieve institution courses!"}
	}
	return institutionCourses, alert
}

/**
Get (ETL) transformed user matric subjects given userId
 */
func getTransformedUserMatricSubjects(app *config.Env, userId string) ([]ExtendedUserMatricSubject, PageToast) {
	var eUserMatricSubjects []ExtendedUserMatricSubject
	userMatricSubjects, alert := getUserMatricSubjects(app, userId)
	if alert.AlertInfo == "" {
		for _, userMatricSubject := range userMatricSubjects {
			subject, alert := getSubject(app, userMatricSubject.SubjectId)
			if alert.AlertInfo == "" {
				eUserMatricSubject := ExtendedUserMatricSubject{userMatricSubject, subject.Name}
				eUserMatricSubjects = append(eUserMatricSubjects, eUserMatricSubject)
			} else {
				break
			}
		}
	}
	return eUserMatricSubjects, alert
}

/**
Get user matric subjects given user id.
 */
func getUserMatricSubjects(app *config.Env, userId string) ([]userDomain.UserMatricSubject, PageToast) {
	var userMatricSubjects []userDomain.UserMatricSubject
	var alert PageToast
	userMatricSubjects, err := usersIO.GetUserMatricSubjects(userId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = PageToast{dangerAlertStyle, "Could not retrieve user matric subjects!"}
	}
	return userMatricSubjects, alert
}
