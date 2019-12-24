package controller

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	domain "obas/domain/application"
	domain2 "obas/domain/users"
	domain4 "obas/domain/util"
	"obas/io/academics"
	applicationIO "obas/io/applications"
	"obas/io/documents"
	"obas/io/institutions"
	"obas/io/users"
	"obas/io/util"
	"time"
	//usersIO "obas/io/users"
)

func Admin(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", AdminHandler(app))
	r.Get("/applicant", AdminApplicantHandler(app))
	r.Get("/applicant/application/{userId}/{applicationId}", AdminApplicantApplicationHandler(app))
	r.Get("/application", AdminApplicationHandler(app))

	r.Post("/email", AdminEmailHandler(app))
	r.Post("/change/document-status", ChangeDocumentStatusHandler(app))
	r.Post("/change/application-status", ChangeApplicationStatusHandler(app))

	return r
}

func ChangeDocumentStatusHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		documentStatusId := r.PostFormValue("documentStatusId")
		documentId := r.PostFormValue("documentId")

		if documentStatusId != "" || documentId != "" {
			document, err := documents.GetDocumentStatus(documentId)
			if err != nil {
				fmt.Println("error reading applicationStatues in getSearchResult")
			} else {
				_, err := documents.updateDocument
			}
		}

		http.Redirect(w, r, "/support/management/academics", 301)
	}
}

func ChangeApplicationStatusHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		applicationStatus := r.PostFormValue("applicationStatus")
		applicationId := r.PostFormValue("applicationId")
		if applicationStatus != "" || applicationId != "" {

		}

		http.Redirect(w, r, "/support/management/academics", 301)
	}
}

func AdminEmailHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		studentEmail := r.PostFormValue("studentEmail")
		message := r.PostFormValue("message")
		fmt.Println("Sending mail to  " + studentEmail + "\nThe message is:\n" + message)

		http.Redirect(w, r, "/support/management/academics", 301)
	}
}
func getDocument(docId string) documents.Document {
	entity := documents.Document{}
	document, err := documents.GetDocument(docId)
	if err != nil {
		return entity
	}
	return document

}

type applicantDetails struct {
	ApplicantionId string
	Email          string
	Name           string
	LatName        string
	ApplicantType  string
	institution    string
	Course         string
	DateTime       time.Time
}
type applicantsearch struct {
	ApplicantDetails   applicantDetails
	ApplicationStatus  string
	ApplicationStatues []domain4.GenericStatus
	Modifier           users.User
	ModificationDate   time.Time
	Comment            string
}
type documentDetails struct {
	DocumentType   string
	Status         string
	Document       users.UserDocument
	DocumentStatus []domain4.GenericStatus
}

func getUser(userId string) users.User {
	var entity = users.User{}
	user, err := users.GetUser(userId)
	if err != nil {
		fmt.Println("error reading applicationStatus in getSearchResult")
		return entity
	}
	return user
}
func getStatus(statusId string) domain4.GenericStatus {
	var entity = domain4.GenericStatus{}
	status, err := util.GetStatus(statusId)
	if err != nil {
		fmt.Println("error reading status in getSearchResult")
		return entity
	}
	return status
}

func getSearchResult(applicationId string) applicantsearch {
	var entity = applicantsearch{}
	for _, applicantDetails := range getApplicants() {
		if applicantDetails.ApplicantionId == applicationId {
			//we get the application detail for the user here
			applicationStatus, err := applicationIO.GetApplicationStatus(applicantDetails.ApplicantionId)
			if err != nil {
				fmt.Println("error reading applicationStatus in getSearchResult")
			}
			Statues, err := util.GetStatuses()
			if err != nil {
				fmt.Println("error reading applicationStatues in getSearchResult")
			}
			return applicantsearch{applicantDetails, getStatus(applicationStatus.StatusId).Name, Statues, getUser(applicationStatus.ModifiedBy), applicationStatus.DateTime, applicationStatus.Comment}
		}
	}
	return entity
}

func getApplicants() []applicantDetails {
	var applicantD []applicantDetails
	var applicant applicantDetails
	userApplications, err := users.GetAllUserApplications()
	if err != nil {
		fmt.Println("error reading userApplications in AdminApplicantApplicationHandler")
	}
	applications, err := applicationIO.GetApplications()
	if err != nil {
		fmt.Println("error reading applications in AdminApplicantApplicationHandler")
	}

	for _, application := range applications {
		userApplications, err := users.GetUserApplications("")
		if err != nil {
			fmt.Println("error reading userApplications in AdminApplicantApplicationHandler")
		}
		user, err := users.GetUser(userApplication.UserId)
		if err != nil {
			fmt.Println("error reading user in AdminApplicantApplicationHandler")
		}
		institution, err := institutions.GetInstitutionType(application.InstitutionId)
		if err != nil {
			fmt.Println("error reading institution in AdminApplicantApplicationHandler")
		}
		course, err := academics.GetCourse(application.CourseId)
		if err != nil {
			fmt.Println("error reading course in AdminApplicantApplicationHandler")
		}
		applicantType, err := applicationIO.GetApplicantType(application.ApplicantTypeId)
		if err != nil {
			fmt.Println("error reading applicantType in AdminApplicantApplicationHandler")
		}

		applicant = applicantDetails{userApplication.ApplicationId, user.Email, user.FirstName, user.LastName, applicantType.Name, institution.Name, course.CourseName, userApplication.DateTime}

		applicantD = append(applicantD, applicant)
	}
	return applicantD
}

func AdminApplicantApplicationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userId := chi.URLParam(r, "userId")
		applicationId := chi.URLParam(r, "applicationId")
		var myDocumentList []documentDetails

		if userId != "" || applicationId != "" {

			userDocuments, err := users.GetUserDocuments(userId)
			if err != nil {
				fmt.Println("error reading userDocuments in AdminApplicantApplicationHandler")
			}
			documentsStatus, err := util.GetStatuses()
			if err != nil {
				fmt.Println("error reading userDocuments in AdminApplicantApplicationHandler")
			}

			for _, document := range userDocuments {
				doc, err := documents.GetDocument(document.DocumentId)
				documentType, err := documents.GetDocumentType(document.DocumentId)
				if err != nil {
					fmt.Println("error reading documentType in AdminApplicantApplicationHandler")
				}
				myDocumentList = append(myDocumentList, documentDetails{documentType.DocumentTypename, doc.DocumentStatus, document, documentsStatus})
			}
		}

		type PageData struct {
			Applicant   []applicantDetails
			Document    []documentDetails
			Application applicantsearch
		}
		Data := PageData{getApplicants(), myDocumentList, getSearchResult(applicationId)}
		files := []string{
			app.Path + "content/admin/admin_applicant.html",
			app.Path + "content/admin/template/sidebar.template.html",
			app.Path + "content/admin/template/navbar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, Data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func AdminApplicationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "content/admin/admin_application.html",
			app.Path + "content/admin/template/sidebar.template.html",
			app.Path + "content/admin/template/navbar.template.html",
			app.Path + "base/template/footer.template.html",
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

type MyUserApplication struct {
	Application     domain.Application
	User            domain2.User
	UserApplication domain2.UserApplication
}

//func getUserAplication(applicationId string)domain2.UserApplication{
//	//return usersIO.
//	return func(w http.ResponseWriter, r *http.Request) {
//		return
//	}
//}

func AdminApplicantHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Pagedate struct {
			Applicant []applicantDetails
		}
		Data := Pagedate{getApplicants()}
		files := []string{
			app.Path + "content/admin/admin_applicant.html",
			app.Path + "content/admin/template/sidebar.template.html",
			app.Path + "content/admin/template/navbar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, Data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func AdminHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/**email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := users_io.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		fmt.Println("User ", user)*/

		files := []string{
			app.Path + "content/admin/admin_dashboard.page.html",
			app.Path + "content/admin/template/sidebar.template.html",
			app.Path + "content/admin/template/navbar.template.html",
			app.Path + "base/template/footer.template.html",
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
