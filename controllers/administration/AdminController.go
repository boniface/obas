package controller

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"html/template"
	"net/http"
	"obas/config"
	domain "obas/domain/application"
	domain3 "obas/domain/documents"
	domain2 "obas/domain/users"
	domain4 "obas/domain/util"
	"obas/io/academics"
	applicationIO "obas/io/applications"
	"obas/io/demographics"
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
	r.Get("/applicant/application/{userId}/{applicationId}", AdminApplicationDocumentsHandler(app))
	r.Get("/applicant/document/{applicationId}/{userId}", AdminApplicantDocumentsHandler(app))
	r.Get("/applicant/applications", AdminApplicationsDocumentsHandler(app))
	r.Get("/application", AdminApplicationHandler(app))

	r.Post("/email", AdminEmailHandler(app))
	r.Post("/change/document-status", ChangeDocumentStatusHandler(app))
	r.Post("/change/application-status", ChangeApplicationStatusHandler(app))

	return r
}

func AdminApplicantDocumentsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		applicationId := chi.URLParam(r, "applicationId")
		user := chi.URLParam(r, "userId")

		if user != "" || applicationId != "" {
			app.Session.Destroy(r.Context())
			app.Session.Put(r.Context(), "userId", email) //this should be admin's email
			app.Session.Put(r.Context(), "token", token)

			app.Session.Put(r.Context(), "applicationId", applicationId)
			app.Session.Put(r.Context(), "user", user) //this user is the student's email

			app.Session.Put(r.Context(), "Admin_message", "")
		}
		http.Redirect(w, r, "/admin/applicant/applications", 301)
		return
	}
}

func ChangeDocumentStatusHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		app.Session.Put(r.Context(), "Admin_message", "Fail to Update")

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_, err := users.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		/***This method still not existing in the backend
		userRole, err := users.GetUserRoleWithUserId(email)
		if err != nil {
			fmt.Println("error reading userRole in ChangeDocumentStatusHandler")
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		role, err := demographics.GetRole(userRole.RoleId)
		if err != nil {
			fmt.Println("error reading role in ChangeDocumentStatusHandler")
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		if role.RoleName != "admin" {
			fmt.Println("error Not an Admin in ChangeDocumentStatusHandler")
			http.Redirect(w, r, "/login", 301)
			return
		}**/

		r.ParseForm()
		documentStatusId := r.PostFormValue("documentStatusId")
		documentId := r.PostFormValue("documentId")
		comment := r.PostFormValue("comment")
		UserId := r.PostFormValue("UserId")
		applicationId := r.PostFormValue("applicationId")

		if documentStatusId != "" || documentId != "" {
			documentStatus := domain3.DocumentStatus{documentId, documentStatusId, email, comment, time.Now()}
			_, err := documents.CreateDocumentStatus(documentStatus)
			if err != nil {
				fmt.Println("error reading document in ChangeDocumentStatusHandler")
			}

			app.Session.Destroy(r.Context())
			app.Session.Put(r.Context(), "userId", email)
			app.Session.Put(r.Context(), "token", token)
			app.Session.Put(r.Context(), "user", UserId)
			app.Session.Put(r.Context(), "applicationId", applicationId)
			app.Session.Put(r.Context(), "Admin_message", "Successfully Updated")
		}
		http.Redirect(w, r, "/admin/applicant/applications", 301)
	}
}

func ChangeApplicationStatusHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_, err := users.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		/***This method still not existing in the backend
		userRole, err := users.GetUserRoleWithUserId(email)
		if err != nil {
			fmt.Println("error reading userRole in ChangeApplicationStatusHandler")
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		role, err := demographics.GetRole(userRole.RoleId)
		if err != nil {
			fmt.Println("error reading role in ChangeApplicationStatusHandler")
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		if role.RoleName != "admin" {
			fmt.Println("error Not an Admin in ChangeApplicationStatusHandler")
			http.Redirect(w, r, "/login", 301)
			return
		}**/

		r.ParseForm()

		applicationStatus := r.PostFormValue("applicationStatus")
		applicationId := r.PostFormValue("applicationId")
		UserId := r.PostFormValue("UserId")
		comment := r.PostFormValue("comment")
		if applicationStatus != "" || applicationId != "" {
			newApplicationStatus := applicationIO.ApplicationStatus{applicationId, applicationStatus, email, comment, time.Now()}

			_, err := applicationIO.CreateApplicationStatus(newApplicationStatus)
			if err != nil {
				fmt.Println("error creating ApplicationStatus in ChangeApplicationStatusHandler")
			}
			app.Session.Destroy(r.Context())
			app.Session.Put(r.Context(), "userId", email)
			app.Session.Put(r.Context(), "token", token)
			app.Session.Put(r.Context(), "user", UserId)
			app.Session.Put(r.Context(), "Admin_message", "Successfully Updated")
		}
		http.Redirect(w, r, "/admin/applicant/applications", 301)
	}
}

func AdminEmailHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_, err := users.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		/***This method still not existing in the backend**/
		userRole, err := users.GetUserRoleWithUserId(email)
		if err != nil {
			fmt.Println("error reading userRole in ChangeApplicationStatusHandler")
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		role, err := demographics.GetRole(userRole.RoleId)
		if err != nil {
			fmt.Println("error reading role in ChangeApplicationStatusHandler")
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		if role.RoleName != "admin" {
			fmt.Println("error Not an Admin in ChangeApplicationStatusHandler")
			http.Redirect(w, r, "/login", 301)
			return
		}

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
	ApplicantionId  string
	Email           string
	Name            string
	LatName         string
	ApplicantType   string
	Institution     string
	Course          string
	DateTime        time.Time
	ApplicationStat applicationIO.ApplicationStatus
	Stat            string
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
	Id             string
	DocumentType   string
	DocumentStatu  domain3.DocumentStatus
	Document       users.UserDocument
	DocumentStatus []domain4.GenericStatus
	Doc            documents.Document
	UserId         string
	ApplicationId  string
	Stat           string
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

	applications, err := applicationIO.GetApplications()
	fmt.Println(" applications>>>> ", applications)
	if err != nil {
		fmt.Println("error reading applications in getApplicants")
	}

	for _, application := range applications {
		/**Getting the userApplication so that we can know the user,and date of that application**/
		userApplication, err := users.GetUserApplicationWithAppId(application.Id)
		if err != nil {
			fmt.Println("error reading userApplications in getApplicants")
		}
		//***Getting the user here**/
		user, err := users.GetUser(userApplication.UserId)
		if err != nil {
			fmt.Println("error reading user in getApplicants")
		}
		institution, err := institutions.GetInstitutionType(application.ApplicationTypeId)
		if err != nil {
			fmt.Println("error reading institution in getApplicants")
		}

		userApplicationCourse, err := users.GetUserApplicationCourseForAppl(user.Email, application.Id)
		if err != nil {
			fmt.Println("error reading userApplicationCourse in getApplicants")
		}

		course, err := academics.GetCourse(userApplicationCourse.UserId)
		if err != nil {
			fmt.Println("error reading course in getApplicants")
		}
		applicationStat, err := applicationIO.GetApplicationStatus(application.ApplicantTypeId)
		if err != nil {
			fmt.Println("error reading applicationStat in getApplicants")
		}
		//Getting the status from the ApplicationStatusId
		status, err := util.GetStatus(applicationStat.StatusId)
		if err != nil {
			fmt.Println("error reading applicationStat in getApplicants")
		}

		applicantType, err := applicationIO.GetApplicantType(application.ApplicantTypeId)
		if err != nil {
			fmt.Println("error reading applicantType in getApplicants")
		}

		applicant = applicantDetails{userApplication.ApplicationId, user.Email, user.FirstName, user.LastName, applicantType.Name, institution.Name, course.CourseName, userApplication.DateTime, applicationStat, status.Name}

		applicantD = append(applicantD, applicant)
	}
	return applicantD
}

//this method takes UserId and applicationId, so that it can find user's documents
func AdminApplicationsDocumentsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//userId := chi.URLParam(r, "applicationId")
		//applicationId := chi.URLParam(r, "userId")
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		applicationId := app.Session.GetString(r.Context(), "applicationId")
		userId := app.Session.GetString(r.Context(), "user")
		Admin_message := app.Session.GetString(r.Context(), "Admin_message")

		fmt.Println(userId, "<<<<<userId||applicationId>>>>", applicationId, "<<<<<")
		var myDocumentList []documentDetails

		documentsStatus, err := util.GetStatuses()
		if err != nil {
			fmt.Println("error reading userDocuments in AdminApplicationDocumentsHandler")
		}

		if userId != "" || applicationId != "" {

			userDocuments, err := users.GetUserDocuments(userId)
			fmt.Println("userDocuments>>>>", userDocuments)
			if err != nil {
				fmt.Println("error reading userDocuments in AdminApplicationDocumentsHandler")
			}

			for _, document := range userDocuments {
				doc, err := documents.GetDocument(document.DocumentId)
				if err != nil {
					fmt.Println("error reading doc in AdminApplicationDocumentsHandler")
				}
				documentType, err := documents.GetDocumentType(doc.DocumentTypeId)
				fmt.Println("documentType>>>>>>>   ", documentType)
				if err != nil {
					fmt.Println("error reading documentType in AdminApplicationDocumentsHandler")
				}
				myDocumentList = append(myDocumentList, documentDetails{doc.DocumentId, documentType.DocumentTypename, getDocumentStat(doc.DocumentId), document, documentsStatus, doc, userId, applicationId, getstat(getDocumentStat(doc.DocumentId).StatusId)})
			}
		}

		type PageData struct {
			Applicant   []applicantDetails
			Document    []documentDetails
			Application applicantsearch
			Message     string
		}
		Data := PageData{getApplicants(), myDocumentList, getSearchResult(applicationId), Admin_message}

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

func getDocumentStat(documentId string) domain3.DocumentStatus {
	entity := domain3.DocumentStatus{}
	resp, err := documents.GetDocumentStatus(documentId)
	if err != nil {
		fmt.Println("error reading GetDocumentStatus getDocumentStat", resp)
		return entity
	}
	return resp
}
func getstat(statId string) string {
	result := ""
	resp, err := util.GetStatus(statId)
	if err != nil {
		fmt.Println("error reading GetStatus in getstat", resp)
		return result
	}
	return resp.Name
}

//this method takes UserId and applicationId, so that it can find user's documents
func AdminApplicationDocumentsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userId := chi.URLParam(r, "userId")
		applicationId := chi.URLParam(r, "applicationId")
		fmt.Println(userId, "<<<<<userId||applicationId>>>>", applicationId)
		var myDocumentList []documentDetails

		documentsStatus, err := util.GetStatuses()
		if err != nil {
			fmt.Println("error reading userDocuments in AdminApplicationDocumentsHandler")
		}

		if userId != "" || applicationId != "" {

			userDocuments, err := users.GetUserDocuments(userId)
			if err != nil {
				fmt.Println("error reading userDocuments in AdminApplicationDocumentsHandler")
			}

			for _, document := range userDocuments {
				doc, err := documents.GetDocument(document.DocumentId)
				if err != nil {
					fmt.Println("error reading doc in AdminApplicationDocumentsHandler")
				}
				documentType, err := documents.GetDocumentType(doc.DocumentTypeId)
				fmt.Println("documentType>>>>>>>   ", documentType)
				if err != nil {
					fmt.Println("error reading documentType in AdminApplicationDocumentsHandler")
				}
				myDocumentList = append(myDocumentList, documentDetails{doc.DocumentId, documentType.DocumentTypename, getDocumentStat(doc.DocumentId), document, documentsStatus, doc, userId, applicationId, getstat(getDocumentStat(doc.DocumentId).StatusId)})
			}
		}
		render.JSON(w, r, myDocumentList)
		/***
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
		}*/
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

		documents := []documentDetails{}
		application := applicantsearch{}
		type Pagedate struct {
			Applicant   []applicantDetails
			Document    []documentDetails
			Application applicantsearch
		}
		Data := Pagedate{getApplicants(), documents, application}
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
