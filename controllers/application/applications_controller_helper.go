package controllers

import (
	"obas/config"
	genericHelper "obas/controllers/misc"
	domain "obas/domain/application"
	applicationIO "obas/io/applications"
)

func GetApplicantTypes(app *config.Env)([]domain.ApplicantType, genericHelper.PageToast){
	var applicantTypes []domain.ApplicantType
	var alert genericHelper.PageToast
	applicantTypes, err := applicationIO.GetApplicantTypes()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve applicant types!"}
	}
	return applicantTypes, alert
}

func GetApplicationTypes(app *config.Env)([]domain.ApplicationType, genericHelper.PageToast) {
	var applicationTypes []domain.ApplicationType
	var alert genericHelper.PageToast
	applicationTypes, err := applicationIO.GetApplicationTypes()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve application types!"}
	}
	return applicationTypes, alert
}

func GetMatricApplicantType(app *config.Env) (domain.ApplicantType, genericHelper.PageToast) {
	var matricApplicantType domain.ApplicantType
	var alert genericHelper.PageToast
	matricApplicantType, err := applicationIO.GetMatricApplicantType()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve matric applicant type!"}
	}
	return matricApplicantType, alert

}

/**
Get application given application id
 */
func GetApplication(app *config.Env, applicationId string)(domain.Application, genericHelper.PageToast) {
	var application domain.Application
	var alert genericHelper.PageToast
	application, err := applicationIO.GetApplication(applicationId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve application types!"}
	}
	return application, alert
}

/**
Check if an application is completed given the application id
 */
func IsApplicationCompleted(app *config.Env, applicationId string) (bool, genericHelper.PageToast) {
	var isComplete bool
	var alert genericHelper.PageToast
	isComplete, err := applicationIO.IsApplicationCompleted(applicationId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve status of latest application!"}
	}
	return isComplete, alert
}

/**
Get application status given application id
 */
func GetApplicationStatus(app *config.Env, applicationId string) (domain.ApplicationStatus, genericHelper.PageToast) {
	var applicationStatus domain.ApplicationStatus
	var alert genericHelper.PageToast
	applicationStatus, err := applicationIO.GetApplicationStatus(applicationId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve application status!"}
	}
	return applicationStatus, alert
}
