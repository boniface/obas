package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"obas/config"
	academicsDomain "obas/domain/academics"
	applicationDomain "obas/domain/application"
	institutionDomain "obas/domain/institutions"
	locationDomain "obas/domain/location"
	userDomain "obas/domain/users"
	addressIO "obas/io/address"
	applicationIO "obas/io/applications"
	demographyIO "obas/io/demographics"
	documentIO "obas/io/documents"
	locationIO "obas/io/location"
	loginIO "obas/io/login"
	storageIO "obas/io/storage"
	usersIO "obas/io/users"
	utilIO "obas/io/util"
	"obas/util"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	layoutOBAS        = "2006-01-02"
	dangerAlertStyle  = "alert-danger"
	successAlertStyle = "alert-success"
)

type AddressPlaceHolder struct {
	AddressName string
	Address     string
	PostalCode  string
}

type PageToast struct {
	AlertType string
	AlertInfo string
}

type ContactPlaceHolder struct {
	ContactName   string
	ContactDetail string
}

type DistrictData struct {
	Student   usersIO.User
	Provinces []locationDomain.Location
	TownName  string
	Alert     PageToast
	Menu      string
	SubMenu   string
}

func Users(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", UsersHandler(app))
	r.Get("/admin", AdminHandler(app))
	r.Get("/student", StudentHandler(app))

	r.Get("/processingStatus", ProcessingStatusTypeHandler(app))
	r.Get("/student/bursary", StudentApplicationStatusHandler(app))
	r.Get("/studentContact", StudentContactsHandler(app))
	r.Get("/student/documents", StudentDocumentsHandler(app))
	r.Get("/studentResults", StudentResultsHandler(app))

	r.Get("/student/profile/personal", StudentProfilePersonalHandler(app))
	r.Get("/student/profile/demography", StudentProfileDemographyHandler(app))
	r.Get("/student/profile/address", StudentProfileAddressHandler(app))
	r.Get("/student/profile/relative", StudentProfileRelativeHandler(app))
	r.Get("/student/profile/settings", StudentProfileSettingsHandler(app))
	r.Get("/student/profile/courses", StudentProfileCourseHandler(app))
	r.Get("/student/profile/academics", StudentProfileSubjectHandler(app))
	r.Get("/student/profile/districts", StudentProfileDistrictHandler(app))
	r.Get("/student/profile/contacts", StudentProfileContactsHandler(app))
	r.Get("/student/profile/application_process", StudentProfileApplicationProcessHandler(app))

	r.Post("/student/profile/personal/update", UpdateStudentProfilePersonalHandler(app))
	r.Post("/student/profile/address/addresstype", StudentProfileAddressTypeHandler(app))
	r.Post("/student/profile/address/update", StudentProfileAddressUpdateHandler(app))
	r.Post("/student-profile-relative-upate", StudentProfileRelativeUpdateHandler(app))
	r.Post("/student-profile-demography-update", StudentProfileDemographyUpdateHandler(app))
	r.Post("/student-profile-password-update", StudentProfilePasswordUpdate(app))
	r.Post("/student/profile/contact/contacttype", StudentProfileContactTypeHandler(app))
	r.Post("/student-profile-contact-update", StudentProfileContactUpdateHandler(app))
	r.Post("/student-profile-town-update", StudentProfileTownUpdateHandler(app))
	r.Post("/student-document-file-upload", StudentDocumentsUploadHandler(app))

	r.Get("/student/bursary/application", StudentBursaryApplicationHandler(app))
	r.Post("/student/bursary/application/start", StudentBursaryApplicationStartHandler(app))
	r.Post("/student/bursary/application/institution/matric/update", StudentBursaryApplicationMatricHandler(app))
	r.Post("/student/bursary/application/type/update", StudentBursaryApplicationTypeHandler(app))
	r.Post("/student/bursary/application/matric/subject/update", StudentBursaryApplicationMatricSubjectHandler(app))

	return r
}

func StudentBursaryApplicationMatricSubjectHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Subject score NOT saved!"
		successMessage := "Subject scored saved!"
		var userMatricSubject userDomain.UserMatricSubject

		r.ParseForm()
		subjectId := r.PostFormValue("subject")
		scoreStr := r.PostFormValue("score")
		if subjectId == "" || scoreStr == "" {
			errorMsg := " Subject and/or score can't be empty!"
			app.ErrorLog.Println(errorMsg)
			setSessionMessage(app, r, dangerAlertStyle, failureMessage+errorMsg)
		} else {
			score, err := strconv.ParseFloat(scoreStr, 64)
			if err != nil {
				errorMsg := " ~ Possible incorrect subject score value."
				app.ErrorLog.Println(err.Error())
				setSessionMessage(app, r, dangerAlertStyle, failureMessage+errorMsg)
			} else {
				userMatricSubject = userDomain.UserMatricSubject{user.Email, subjectId, score}
				app.InfoLog.Println("UserMatricSubject to save: ", userMatricSubject)
				userMatricSubject, err = usersIO.CreateUserMatricSubject(userMatricSubject)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					setSessionMessage(app, r, dangerAlertStyle, failureMessage)
				} else {
					setSessionMessage(app, r, successAlertStyle, successMessage)
				}
			}
		}
		app.InfoLog.Println("Matic Subject saved: ", userMatricSubject)
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Applicant Type NOT updated!"
		successMessage := "Applicant Type updated!"

		r.ParseForm()
		applicantTypeId := r.PostFormValue("applicantType")
		applicationId := r.PostFormValue("applicationId")
		application, err := applicationIO.GetApplication(applicationId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			setSessionMessage(app, r, dangerAlertStyle, failureMessage)
		} else {
			updatedApplication := application
			updatedApplication.ApplicantTypeId = applicantTypeId
			saved, err := applicationIO.UpdateApplication(updatedApplication, token)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				setSessionMessage(app, r, dangerAlertStyle, failureMessage)
			} else {
				app.InfoLog.Println("update application response is ", saved)
				setSessionMessage(app, r, successAlertStyle, successMessage)
			}
		}
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationMatricHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Matric Institution NOT saved!"
		successMessage := "Matric Institution saved!"

		r.ParseForm()
		institution := r.PostFormValue("institution")
		userMatricInstitution := userDomain.UserMatricInstitution{user.Email, institution}
		app.InfoLog.Println("User (matric) institution to save: ", userMatricInstitution)

		saved, err := usersIO.CreateUserMatricInstitution(userMatricInstitution)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			setSessionMessage(app, r, dangerAlertStyle, failureMessage)
		} else {
			setSessionMessage(app, r, successAlertStyle, successMessage)
		}
		app.InfoLog.Println("Matric Institution saved: ", saved)
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		var alert PageToast
		var applicationTypes []applicationDomain.ApplicationType
		var applicantTypes []applicationDomain.ApplicantType
		var application applicationDomain.Application
		var provinces []locationDomain.Location
		var institutionTypes []institutionDomain.InstitutionTypes
		var userMatricInstitution userDomain.UserMatricInstitution
		var userMatricInstitutionName string
		var matricSubjects []academicsDomain.Subject
		var eUserMatricSubjects []ExtendedUserMatricSubject
		isComplete := true
		latestUserApplication, err := usersIO.GetLatestUserApplication(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = PageToast{dangerAlertStyle, "Could not retrieve your latest application!"}
		} else {
			if latestUserApplication.ApplicationId != "" {
				isComplete, err = applicationIO.IsApplicationCompleted(latestUserApplication.ApplicationId)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					alert = PageToast{dangerAlertStyle, "Could not retrieve status of latest application"}
				} else {
					application, err = applicationIO.GetApplication(latestUserApplication.ApplicationId)
					if err != nil {
						app.ErrorLog.Println(err.Error())
						alert = PageToast{dangerAlertStyle, "Could not retrieve application!"}
					}
				}
			}
			applicantTypes, err = applicationIO.GetApplicantTypes()
			if err != nil {
				app.ErrorLog.Println(err.Error())
				alert = PageToast{dangerAlertStyle, "Could not retrieve applicant types!"}
			} else {
				if isComplete {
					applicationTypes, err = applicationIO.GetApplicationTypes()
					if err != nil {
						app.ErrorLog.Println(err.Error())
						alert = PageToast{dangerAlertStyle, "Could not retrieve application types!"}
					} else {
						alert = checkForSessionAlert(app, r)
					}
				} else {
					provinces, alert = getProvinces(app)
					proceed := alert.AlertInfo == ""
					if proceed {
						institutionTypes, alert = getInstitutionTypes(app)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						userMatricInstitution, alert = getUserMatricInstitution(app, user.Email)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						userMatricInstitutionName, alert = getUserMatricInstitutionName(app, userMatricInstitution.InstitutionId)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						matricSubjects, alert = getMatricSubjects(app, userMatricInstitution.InstitutionId)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						eUserMatricSubjects, alert = getTransformedUserMatricSubjects(app, user.Email)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						alert = checkForSessionAlert(app, r)
					}
				}
			}
		}

		type PageData struct {
			Student                   usersIO.User
			Menu                      string
			SubMenu                   string
			ApplicationTypes          []applicationDomain.ApplicationType
			ApplicantTypes            []applicationDomain.ApplicantType
			Application               applicationDomain.Application
			IsComplete                bool
			Provinces                 []locationDomain.Location
			InstitutionTypes          []institutionDomain.InstitutionTypes
			UserMatricInstitutionName string
			MatricSubjects            []academicsDomain.Subject
			UserMatricSubjects        []ExtendedUserMatricSubject
			Alert                     PageToast
		}

		data := PageData{user, "bursary", "application", applicationTypes, applicantTypes, application, isComplete, provinces, institutionTypes, userMatricInstitutionName, matricSubjects, eUserMatricSubjects, alert}
		files := []string{
			app.Path + "content/student/bursary/application.html",
			app.Path + "content/student/template/sidebar.template.html",
			app.Path + "content/student/template/application/matric-form.template.html",
			app.Path + "content/student/template/application/current-institution-form.template.html",
			app.Path + "base/template/form/institution-form.template.html",
			app.Path + "base/template/form/location-form.template.html",
			app.Path + "content/student/template/application/prospective-institution-form.template.html",
			app.Path + "content/student/template/application/institution-course.template.html",
			app.Path + "content/student/template/application/document.template.html",
			app.Path + "content/student/template/application/applicant-type-form.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}

}

func StudentBursaryApplicationStartHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		var newApplication = applicationDomain.Application{}
		failureMessage := "Application NOT created!"
		successMessage := "Application created!"
		isSuccessful := false

		r.ParseForm()
		applicationTypeId := r.PostFormValue("applicationType")
		applicantTypeId := r.PostFormValue("applicantType")

		if applicationTypeId != "" && applicantTypeId != "" {
			application := applicationDomain.Application{"", applicationTypeId, applicantTypeId}
			app.InfoLog.Println("Application to create: ", application)
			newApplication, err = applicationIO.CreateApplication(application)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				setSessionMessage(app, r, dangerAlertStyle, failureMessage)
			} else {
				if newApplication.Id != "" {
					userApplication := userDomain.UserApplication{email, newApplication.Id, time.Now()}
					app.InfoLog.Println("User Application to create: ", userApplication)
					_, err := usersIO.CreateUserApplication(userApplication)
					if err != nil {
						app.ErrorLog.Println(err.Error())
						setSessionMessage(app, r, dangerAlertStyle, failureMessage)
					} else {
						incompleteStatus, err := utilIO.GetIncompleteStatus()
						if err != nil {
							app.ErrorLog.Println(err.Error())
							setSessionMessage(app, r, dangerAlertStyle, failureMessage)
						} else {
							if incompleteStatus.Id != "" {
								userApplicationStatus := applicationIO.ApplicationStatus{newApplication.Id, incompleteStatus.Id, user.Email, "Starting Application", time.Now()}
								app.InfoLog.Println("Application Status to create: ", userApplicationStatus)
								_, err = applicationIO.CreateApplicationStatus(userApplicationStatus)
								if err != nil {
									app.ErrorLog.Println(err.Error() + " ~ User Application Status NOT created!")
									setSessionMessage(app, r, dangerAlertStyle, failureMessage)
								} else {
									isSuccessful = true
									setSessionMessage(app, r, successAlertStyle, successMessage)
								}
							} else {
								app.ErrorLog.Println("No status id found!")
								setSessionMessage(app, r, dangerAlertStyle, failureMessage)
							}
						}
					}
				} else {
					app.ErrorLog.Println("No application id!")
					setSessionMessage(app, r, dangerAlertStyle, failureMessage)
				}
			}
		} else {
			error := "Application type and/or applicant type is null!"
			app.ErrorLog.Println(error)
			setSessionMessage(app, r, dangerAlertStyle, failureMessage+" Reason: "+error)
		}
		app.InfoLog.Println("application response is ", isSuccessful)
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentProfileApplicationProcessHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/bursary/student_application_process.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentDocumentsUploadHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseMultipartForm(10 << 20)
		documentTypeId := r.PostFormValue("documenttype")
		file, handler, err := r.FormFile("file")
		if err != nil {
			app.ErrorLog.Println(err.Error())
			setSessionMessage(app, r, dangerAlertStyle, "Internal Server Error!")
		} else {
			defer file.Close()
			tempFilePath := "temp_file/" + handler.Filename
			tempFile, err := os.OpenFile(tempFilePath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				setSessionMessage(app, r, dangerAlertStyle, "Internal Server Error!")
			} else {
				defer tempFile.Close()
				io.Copy(tempFile, file)
				fileName := handler.Filename
				successMessage := "File: " + fileName + " updated!"
				failureMessage := "File: " + fileName + " NOT updated!"
				app.InfoLog.Println("File Name: ", fileName)
				app.InfoLog.Println("File Size: ", handler.Size)
				app.InfoLog.Println("File Header: ", handler.Header)

				fileData, err := storageIO.UploadFile(tempFilePath, token)
				os.Remove(tempFilePath)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					setSessionMessage(app, r, dangerAlertStyle, failureMessage)
				} else {
					if strings.Contains(fileData.Id, "error") {
						app.ErrorLog.Println(fileData.Id)
						setSessionMessage(app, r, dangerAlertStyle, "Internal Server Error!")
					} else {
						document := documentIO.Document{
							fileData.Id,
							documentTypeId,
							fileName,
							fileData.Url,
							handler.Header.Get("Content-Type"),
							time.Now(),
							"", ""}
						app.InfoLog.Println("Document to create: ", document)
						documentSaved, err := documentIO.CreateDocument(document, token)
						if err != nil {
							app.ErrorLog.Println(err.Error())
							setSessionMessage(app, r, dangerAlertStyle, "Document NOT saved!")
						} else {
							if documentSaved {
								userDocument := usersIO.UserDocument{email, fileData.Id}
								app.InfoLog.Println("User document to save: ", userDocument)
								userDocumentSaved, err := usersIO.UpdateUserDocument(userDocument, token)
								if err != nil {
									app.ErrorLog.Println(err.Error())
									setSessionMessage(app, r, dangerAlertStyle, "User document saved!")
								} else {
									if userDocumentSaved {
										setSessionMessage(app, r, successAlertStyle, successMessage)
									} else {
										setSessionMessage(app, r, dangerAlertStyle, failureMessage)
									}
								}
							} else {
								setSessionMessage(app, r, dangerAlertStyle, "User document NOT saved!")
							}
						}
					}
				}
			}
		}
		http.Redirect(w, r, "/users/student/documents", 301)
	}
}

func StudentDocumentsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		type DocumentData struct {
			Document            documentIO.Document
			DocumentType        string
			DocumentDate        string
			DocumentStatusBadge string
		}
		var alert PageToast
		var documentTypes []documentIO.DocumentType
		var userDocuments []DocumentData
		documentTypes, err = documentIO.GetDocumentTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = PageToast{dangerAlertStyle, "Could not retrieve document types!"}
		}

		userDocumentsObj, err := usersIO.GetUserDocuments(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = PageToast{dangerAlertStyle, "Could not retrieve user documents!"}
		} else {
			for _, userDocument := range userDocumentsObj {
				documentId := userDocument.DocumentId
				document, err := documentIO.GetDocument(documentId)
				if err != nil {
					app.ErrorLog.Println(err.Error() + " - Could not retrieve document for id: " + documentId)
				} else {
					documentType, err := documentIO.GetDocumentType(document.DocumentTypeId)
					if err != nil {
						app.ErrorLog.Println(err.Error() + " - Could not retrieve document type for document!")
					} else {
						date := getDate_YYYYMMDD(document.Date.String())
						var progressBadge string
						documentStatus := document.DocumentStatus
						if documentStatus == "Approved" {
							progressBadge = "badge-success"
						} else if documentStatus == "Not Approved" {
							progressBadge = "badge-danger"
						} else {
							progressBadge = "badge-warning"
						}
						documentData := DocumentData{document, documentType.DocumentTypename, date, progressBadge}
						userDocuments = append(userDocuments, documentData)
					}
				}
			}
			alert = checkForSessionAlert(app, r)
		}

		type PageData struct {
			Student       usersIO.User
			DocumentTypes []documentIO.DocumentType
			UserDocuments []DocumentData
			Alert         PageToast
		}
		data := PageData{
			user,
			documentTypes,
			userDocuments,
			alert,
		}
		files := []string{
			app.Path + "content/student/bursary/student_documents.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentProfileTownUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		townCode := r.PostFormValue("town")
		userTown := userDomain.UserTown{email, townCode}
		app.InfoLog.Println("UserTown to update: ", userTown)
		updated, err := usersIO.UpdateUserTown(userTown, token)
		successMessage := "User town updated!"
		failureMessage := "User town NOT updated!"
		if err != nil {
			app.ErrorLog.Println(err.Error())
			setSessionMessage(app, r, dangerAlertStyle, failureMessage)
		} else {
			if updated {
				setSessionMessage(app, r, successAlertStyle, successMessage)
			} else {
				setSessionMessage(app, r, dangerAlertStyle, failureMessage)
			}
		}
		app.InfoLog.Println("Update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/districts", 301)
	}
}

func StudentProfileDistrictHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		var alert PageToast
		var provinces []locationDomain.Location
		var townName string

		provinces, err = util.GetProvinces()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = PageToast{dangerAlertStyle, "Could not retrieve provinces!"}
		} else {
			userTown, err := usersIO.GetUserTown(email)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				alert = PageToast{dangerAlertStyle, "Could not retrieve user town!"}
			} else {
				if userTown.LocationId != "" {
					location, err := locationIO.GetLocation(userTown.LocationId)
					if err != nil {
						app.ErrorLog.Println(err.Error())
						alert = PageToast{dangerAlertStyle, "Could not retrieve location!"}
					} else {
						townName = location.Name
						alert = checkForSessionAlert(app, r)
					}
				} else {
					townName = "<<not set>>"
					alert = checkForSessionAlert(app, r)
				}
			}
		}

		data := DistrictData{
			user,
			provinces,
			townName,
			alert, "profile", "districts"}

		files := []string{
			app.Path + "content/student/profile/district_and_municipality.html",
			app.Path + "content/student/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentProfileContactUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		contactTypeId := r.PostFormValue("contactTypeId")
		contact := r.PostFormValue("contact")
		userContact := usersIO.UserContact{email, contactTypeId, contact}
		app.InfoLog.Println("UserContact to update: ", userContact)
		updated, err := usersIO.UpdateUserContact(userContact, token)
		successMessage := "User contact updated!"
		failureMessage := "User contact NOT updated!"
		if err != nil {
			app.ErrorLog.Println(err.Error())
			setSessionMessage(app, r, dangerAlertStyle, failureMessage)
		} else {
			if updated {
				setSessionMessage(app, r, successAlertStyle, successMessage)
			} else {
				setSessionMessage(app, r, dangerAlertStyle, failureMessage)
			}
		}
		app.InfoLog.Println("Update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/contacts", 301)
	}
}

func StudentProfileContactTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		var alert PageToast
		var userContact usersIO.UserContact
		var contactName string
		var contactTypes []addressIO.ContactType
		var contacts []ContactPlaceHolder
		r.ParseForm()
		contactTypeId := r.PostFormValue("contacttypes")
		if contactTypeId == "" {
			errMsg := "No contact type selected!"
			app.ErrorLog.Println(errMsg)
			alert = PageToast{dangerAlertStyle, errMsg}
		} else {
			contactTypes, err := addressIO.GetContactTypes()
			if err != nil {
				errMsg := "Could not retrieve contact types!"
				app.ErrorLog.Println(err.Error() + " - " + errMsg)
				alert = PageToast{dangerAlertStyle, errMsg}
			} else {
				for _, contactType := range contactTypes {
					if contactTypeId == contactType.ContactTypeId {
						contactName = contactType.Name
					}
					userContact, err := usersIO.GetUserContact(email, contactType.ContactTypeId)
					if err != nil {
						errMsg := "Could not retrieve user contact for " + contactType.Name
						app.ErrorLog.Println(err.Error() + " - " + errMsg)
					} else {
						contacts = append(contacts, ContactPlaceHolder{contactType.Name, userContact.Contact})
					}
				}
				userContact, err = usersIO.GetUserContact(email, contactTypeId)
				if err != nil {
					errMsg := "Could not retrieve user contact for " + contactName
					app.ErrorLog.Println(err.Error() + " - " + errMsg)
					alert = PageToast{dangerAlertStyle, errMsg}
				}
			}
		}

		type PageData struct {
			Student       usersIO.User
			ContactTypes  []addressIO.ContactType
			Contacts      []ContactPlaceHolder
			Contact       usersIO.UserContact
			ContactTypeId string
			ContactName   string
			Alert         PageToast
		}
		data := PageData{user, contactTypes, contacts, userContact, contactTypeId, contactName, alert}
		files := []string{
			app.Path + "content/student/profile/contacts.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentProfileContactsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		var alert PageToast
		var contacts []ContactPlaceHolder
		contactTypes, err := addressIO.GetContactTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = PageToast{dangerAlertStyle, "Could not retrieve contact types!"}
		} else {
			for _, contactType := range contactTypes {
				userContact, err := usersIO.GetUserContact(email, contactType.ContactTypeId)
				if err != nil {
					errMsg := "Could not retrieve user contact for " + contactType.Name
					app.ErrorLog.Println(err.Error() + " - " + errMsg)
				} else {
					contacts = append(contacts, ContactPlaceHolder{contactType.Name, userContact.Contact})
				}
			}
			alert = checkForSessionAlert(app, r)
		}

		type PageData struct {
			Student       usersIO.User
			ContactTypes  []addressIO.ContactType
			Contacts      []ContactPlaceHolder
			Contact       usersIO.UserContact
			ContactTypeId string
			ContactName   string
			Alert         PageToast
			Menu          string
			SubMenu       string
		}

		data := PageData{user, contactTypes, contacts, usersIO.UserContact{}, "", "", alert, "profile", "contacts"}
		files := []string{
			app.Path + "content/student/profile/contacts.html",
			app.Path + "content/student/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentProfilePasswordUpdate(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		currentPassword := r.PostFormValue("current_password")
		newPasswordOne := r.PostFormValue("new_password_one")
		newPasswordTwo := r.PostFormValue("new_password_two")

		successMessage := "User password updated!"
		failureMessage := "User password NOT Updated!"

		if newPasswordOne != newPasswordTwo {
			errMsg := "New password mismatch."
			app.ErrorLog.Println(errMsg)
			setSessionMessage(app, r, dangerAlertStyle, failureMessage+" - "+errMsg)
		} else {
			userChangePassword := loginIO.ChangePassword{email, currentPassword, newPasswordOne, time.Now()}
			app.InfoLog.Println("User password to update: ", userChangePassword)
			loginToken, err := loginIO.DoChangePassword(userChangePassword, token)

			if err != nil {
				app.ErrorLog.Println(err.Error())
				setSessionMessage(app, r, dangerAlertStyle, failureMessage)
			}
			app.Session.Put(r.Context(), "userId", loginToken.Email)
			app.Session.Put(r.Context(), "token", loginToken.Token)
			setSessionMessage(app, r, successAlertStyle, successMessage)
		}
		http.Redirect(w, r, "/users/student/profile/settings", 301)
	}
}

func StudentProfileSettingsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		alert := checkForSessionAlert(app, r)

		type PageData struct {
			Student usersIO.User
			Alert   PageToast
			Menu    string
			SubMenu string
		}

		data := PageData{user, alert, "profile", "settings"}
		files := []string{
			app.Path + "content/student/profile/settings.html",
			app.Path + "content/student/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentProfileDemographyUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		title := r.PostFormValue("title")
		gender := r.PostFormValue("gender")
		race := r.PostFormValue("race")
		userDemograpgy := usersIO.UserDemography{email, title, gender, race}
		app.InfoLog.Println("userDemography to update: ", userDemograpgy)

		updated, err := usersIO.UpdateUserDemographics(userDemograpgy, token)
		successMessage := "User demography updated!"
		failureMessage := "User demography NOT Updated!"

		if err != nil {
			app.ErrorLog.Println(err.Error())
			setSessionMessage(app, r, dangerAlertStyle, failureMessage)
		} else {
			if updated {
				setSessionMessage(app, r, successAlertStyle, successMessage)
			} else {
				setSessionMessage(app, r, dangerAlertStyle, failureMessage)
			}
		}
		app.InfoLog.Println("UserDemography update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/demography", 301)
	}
}

func StudentProfileDemographyHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		var alert PageToast
		var genders []demographyIO.Gender
		var races []demographyIO.Race
		var gender demographyIO.Gender
		var race demographyIO.Race
		var title demographyIO.Title
		titles, err := demographyIO.GetTitles()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = PageToast{dangerAlertStyle, "Could not retrieve titles!"}
		} else {
			genders, err = demographyIO.GetGenders()
			if err != nil {
				app.ErrorLog.Println(err.Error())
				alert = PageToast{dangerAlertStyle, "Could not retrieve genders!"}
			} else {
				races, err = demographyIO.GetRaces()
				if err != nil {
					app.ErrorLog.Println(err.Error())
					alert = PageToast{dangerAlertStyle, "Could not retrieve races!"}
				} else {
					userDemography, err := usersIO.GetUserDemographic(email)
					if err != nil {
						app.ErrorLog.Println(err.Error())
						alert = PageToast{dangerAlertStyle, "Could not retrieve student demography!"}
					} else {
						title = getUserTitle(userDemography, titles)
						gender = getUserGender(userDemography, genders)
						race = getUserRace(userDemography, races)
						alert = checkForSessionAlert(app, r)
					}
				}
			}
		}

		type PageData struct {
			Student       usersIO.User
			Titles        []demographyIO.Title
			Genders       []demographyIO.Gender
			Races         []demographyIO.Race
			Alert         PageToast
			StudentTitle  demographyIO.Title
			StudentGender demographyIO.Gender
			StudentRace   demographyIO.Race
			Menu          string
			SubMenu       string
		}

		data := PageData{
			user,
			titles,
			genders,
			races,
			alert,
			title,
			gender,
			race,
			"profile",
			"demography"}
		app.InfoLog.Println("DistrictData: ", data.Alert)
		files := []string{
			app.Path + "content/student/profile/demography.html",
			app.Path + "content/student/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentProfileRelativeUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		relativeName := r.PostFormValue("relative_name")
		relationship := r.PostFormValue("relationship")
		cellphone := r.PostFormValue("relative_cellphone")
		relativeEmail := r.PostFormValue("relative_email")
		userRelative := usersIO.UserRelative{email, relativeName, cellphone, relativeEmail, relationship}
		app.InfoLog.Println("UserRelative to update: ", userRelative)
		updated, err := usersIO.UpdateUserRelative(userRelative, token)

		successMessage := "User relative updated!"
		failureMessage := "User relative NOT Updated!"

		if err != nil {
			app.ErrorLog.Println(err.Error())
			setSessionMessage(app, r, dangerAlertStyle, failureMessage)
		} else {
			if updated {
				setSessionMessage(app, r, successAlertStyle, successMessage)
			} else {
				setSessionMessage(app, r, dangerAlertStyle, failureMessage)
			}
		}
		app.InfoLog.Println("UserRelative update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/relative", 301)

	}
}

func setSessionMessage(app *config.Env, r *http.Request, messageType string, message string) {
	app.Session.Put(r.Context(), "message-type", messageType)
	app.Session.Put(r.Context(), "message", message)
}

func StudentProfileSubjectHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/academics.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentProfileCourseHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/courses.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}

}

func StudentProfileRelativeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		var alert PageToast

		userRelative, err := usersIO.GetUserRelative(user.Email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = PageToast{dangerAlertStyle, "Could not retrieve student relative!"}
		} else {
			alert = checkForSessionAlert(app, r)
		}

		type PageData struct {
			Student         usersIO.User
			StudentRelative usersIO.UserRelative
			Alert           PageToast
			Menu            string
			SubMenu         string
		}

		data := PageData{user, userRelative, alert, "profile", "relative"}
		files := []string{
			app.Path + "content/student/profile/relative.html",
			app.Path + "content/student/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentProfileAddressUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		addressTypeId := r.PostFormValue("addressTypeId")
		address := r.PostFormValue("address")
		postalCode := r.PostFormValue("postalCode")
		userAddress := usersIO.UserAddress{email, addressTypeId, address, postalCode}
		app.InfoLog.Println("UserAddress to update: ", userAddress)
		updated, err := usersIO.UpdateUserAddress(userAddress, token)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/address", 301)
	}
}

func StudentProfileAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		addressTypes, err := addressIO.GetAddressTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error(), addressTypes)
		}

		addresses := []AddressPlaceHolder{}

		for _, addressType := range addressTypes {
			userAddress, err := usersIO.GetUserAddress(email, addressType.AddressTypeID)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				addresses = append(addresses, AddressPlaceHolder{addressType.AddressName, userAddress.Address, userAddress.PostalCode})
			}
		}

		type PageData struct {
			Student       usersIO.User
			AddressTypes  []addressIO.AddressType
			Addresses     []AddressPlaceHolder
			Address       usersIO.UserAddress
			AddressTypeID string
			AddressName   string
			Menu          string
			SubMenu       string
		}

		data := PageData{user, addressTypes, addresses, usersIO.UserAddress{}, "", "", "profile", "address"}
		files := []string{
			app.Path + "content/student/profile/address.html",
			app.Path + "content/student/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentProfileAddressTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		addressTypeId := r.PostFormValue("addresstypes")
		userAddress, err := usersIO.GetUserAddress(email, addressTypeId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

		addressTypes, err := addressIO.GetAddressTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error(), addressTypes)
		}

		addresses := []AddressPlaceHolder{}
		var addressName string

		for _, addressType := range addressTypes {
			if addressTypeId == addressType.AddressTypeID {
				addressName = addressType.AddressName
			}
			userAddress, err := usersIO.GetUserAddress(email, addressType.AddressTypeID)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				addresses = append(addresses, AddressPlaceHolder{addressType.AddressName, userAddress.Address, userAddress.PostalCode})
			}
		}

		type PageData struct {
			Student       usersIO.User
			AddressTypes  []addressIO.AddressType
			Addresses     []AddressPlaceHolder
			Address       usersIO.UserAddress
			AddressTypeID string
			AddressName   string
		}

		data := PageData{user, addressTypes, addresses, userAddress, addressTypeId, addressName}
		files := []string{
			app.Path + "content/student/profile/address.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func StudentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		type PageData struct {
			Student usersIO.User
			Menu    string
			SubMenu string
		}
		data := PageData{user, "", ""}
		files := []string{
			app.Path + "content/student/student_dashboard.page.html",
			app.Path + "content/student/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}

}

func StudentProfilePersonalHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		dobString := getDate_YYYYMMDD(user.DateOfBirth.String()) // split date and get in format: yyy-mm-dd

		type PageData struct {
			Student     usersIO.User
			DateOfBirth string
			Menu        string
			SubMenu     string
		}

		data := PageData{user, dobString, "profile", "personal"}
		files := []string{
			app.Path + "content/student/profile/personal.html",
			app.Path + "content/student/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func UpdateStudentProfilePersonalHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		idNumber := r.PostFormValue("id_number")
		firstName := r.PostFormValue("first_name")
		lastName := r.PostFormValue("last_name")
		dateOfBirthStr := r.PostFormValue("dateOfBirth")
		dateOfBirth, _ := time.Parse(layoutOBAS, dateOfBirthStr)
		user := usersIO.User{email, idNumber, firstName, "", lastName, dateOfBirth}
		app.InfoLog.Println("User to update: ", user)
		updated, err := usersIO.UpdateUser(user, token)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/personal", 301)
	}
}

func UsersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			//courses []io.Users
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "base/register/register.page.html",
			/**app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.template.html",*/
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}

func AdminHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allAdmin, err := io.GetAdmins()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//courses []io.Admin
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}

func ProcessingStatusTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allProcess, err := io.GetProcessingStatusTypes()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//academics []io.ProcessingStatusType
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}

func StudentApplicationStatusHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "content/student/bursary/bursary.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())

		}
	}
}

func StudentContactsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allStudentContacts, err := io.GetStudentContacts()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//academics []io.StudentContacts
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}

func StudentResultsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allStudentResults, err := io.GetStudentResults()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//academics []io.StudentResults
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}
