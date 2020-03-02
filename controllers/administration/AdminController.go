package controller

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	domain2 "obas/domain/academics"
	applicationDomain "obas/domain/application"
	domain3 "obas/domain/demographics"
	documentDomain "obas/domain/documents"
	domain "obas/domain/institutions"
	userDomain "obas/domain/users"
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
	r.Get("/application", AdminApplicationHandler(app))

	r.Post("/email", AdminEmailHandler(app))
	r.Post("/change/document-status", ChangeDocumentStatusHandler(app))
	r.Post("/change/application-status", ChangeApplicationStatusHandler(app))

	r.Get("/applicant/document/{applicationId}/{userId}", AdminApplicationDocumentHandler(app))

	return r
}

func AdminApplicationDocumentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clearCash(w)
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		fmt.Println(email, "<<<<<email||token>>>>>", token)

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		userId := chi.URLParam(r, "userId")
		applicationId := chi.URLParam(r, "applicationId")
		//var documents []documentDomain.Document
		app.Session.Put(r.Context(), "Admin_message", "Fail to Update")

		//user, err := users.GetUser(userId)
		//if err != nil {
		//	app.ErrorLog.Println(err.Error())
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		//application := ApplicationDetail{}
		//statues,err:=util.GetStatuses()
		//if err!=nil{
		//
		//}

		type Pagedate struct {
			Applicant         []ApplicantDetails
			Document          []documentDetails
			Application       ApplicationDetail
			Tab               string
			SubMenu           string
			Accordion         string
			ApplicationStatus []ApplicationHolder
		}
		//application:=getSearchResult(applicationId)

		data := Pagedate{getApplicants(), getDocumentDetails(userId, applicationId), getSearchResult(userId, applicationId), "applicant", "", "document", getApplicationStatus(applicationId)}

		//fmt.Println(getDocumentDetails(userId), "<<<<getDocumentDetails(userId)")

		files := []string{ //views/html/content/admin/admin_application3.html
			app.Path + "/content/admin/admin_application3.html",
			app.Path + "/content/admin/template/sidebar.template.html",
			app.Path + "/content/admin/template/navbar.template.html",
			app.Path + "/base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			//return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

type ApplicationHolder struct {
	Modifier userDomain.User
	Comment  string
	Date     time.Time
	Stat     string
}

//this Method help to get the application status with all the necessary data along
func getApplicationStatus(applicationId string) []ApplicationHolder {
	var applicationStatHolder []ApplicationHolder
	applicationStatues, err := applicationIO.GetAllStatusesForApplication(applicationId)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, applicationStatus := range applicationStatues {
		status, err := util.GetStatus(applicationStatus.StatusId)
		if err != nil {
			fmt.Println(err.Error())
		}
		user, err := users.GetUser(applicationStatus.ModifiedBy)
		if err != nil {
			fmt.Println(err.Error())
		}
		applicationStatHolder = append(applicationStatHolder, ApplicationHolder{user, applicationStatus.Comment, applicationStatus.DateTime, status.Name})
	}
	return applicationStatHolder
}

func ChangeDocumentStatusHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		clearCash(w)
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		app.Session.Put(r.Context(), "Admin_message", "Fail to Update")

		fmt.Println(email, "email || token", token)
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
		//userRole, err := users.GetUserRoleWithUserId(email)
		//if err != nil {
		//	fmt.Println("error reading userRole in ChangeDocumentStatusHandler")
		//	app.ErrorLog.Println(err.Error())
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		//role, err := demographics.GetRole(userRole.RoleId)
		//if err != nil {
		//	fmt.Println("error reading role in ChangeDocumentStatusHandler")
		//	app.ErrorLog.Println(err.Error())
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}

		//if role.RoleName != "admin" {
		//	fmt.Println("error Not an Admin in ChangeDocumentStatusHandler")
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}

		r.ParseForm()
		documentStatusId := r.PostFormValue("documentStatusId")
		documentId := r.PostFormValue("docId")
		comment := r.PostFormValue("comment")
		applicationId := r.PostFormValue("applicationId")
		UserId := r.PostFormValue("UserId")

		if documentStatusId != "" || documentId != "" {
			documentStatus := documentDomain.DocumentStatus{documentId, documentStatusId, email, comment, time.Now()}
			_, err := documents.CreateDocumentStatus(documentStatus)
			if err != nil {
				fmt.Println("error reading document in ChangeDocumentStatusHandler")
			}
			//app.Session.Destroy(r.Context())
			//app.Session.Put(r.Context(), "userId", email)
			//app.Session.Put(r.Context(), "token", token)
			app.Session.Put(r.Context(), "Admin_message", "Successfully Updated")
		}
		app.Session.Put(r.Context(), "tab", "document")
		http.Redirect(w, r, "/admin/applicant/document/"+applicationId+"/"+UserId+"", 301)
	}
}

func ChangeApplicationStatusHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clearCash(w)
		var result = "success"
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		//_, err := users.GetUser(email)
		//if err != nil {
		//	app.ErrorLog.Println(err.Error())
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		///***This method still not existing in the backend**/
		//userRole, err := users.GetUserRoleWithUserId(email)
		//if err != nil {
		//	fmt.Println("error reading userRole in ChangeApplicationStatusHandler")
		//	app.ErrorLog.Println(err.Error())
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		//role, err := demographics.GetRole(userRole.RoleId)
		//if err != nil {
		//	fmt.Println("error reading role in ChangeApplicationStatusHandler")
		//	app.ErrorLog.Println(err.Error())
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		//
		//if role.RoleName != "admin" {
		//	fmt.Println("error Not an Admin in ChangeApplicationStatusHandler")
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}

		r.ParseForm()

		applicationStatus := r.PostFormValue("applicationStatus")
		applicationId := r.PostFormValue("applicationId")
		comment := r.PostFormValue("comment")
		userId := r.PostFormValue("UserId")

		fmt.Println(applicationStatus, "<<<applicationStatus>>>"+comment+"<<<<<<applicationId>>>>>>>", applicationId)
		if applicationStatus != "" || applicationId != "" {
			newApplicationStatus := applicationDomain.ApplicationStatus{applicationId, applicationStatus, email, comment, time.Now()}

			_, err := applicationIO.CreateApplicationStatus(newApplicationStatus)
			if err != nil {
				result = "fail"
				fmt.Println("error creating ApplicationStatus in ChangeApplicationStatusHandler")
			}

			/***We first remove the old data in Admin_message session***/
			app.Session.Remove(r.Context(), "Admin_message")
			app.Session.Put(r.Context(), "userId", email)
			app.Session.Put(r.Context(), "token", token)
			app.Session.Put(r.Context(), "Admin_message", result)
		}
		app.Session.Put(r.Context(), "tab", "document")
		http.Redirect(w, r, "/admin/applicant/document/"+applicationId+"/"+userId+"", 301)
	}
}

func AdminEmailHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clearCash(w)
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

		http.Redirect(w, r, "/admin/applicant", 301)
	}
}
func getDocument(docId string) documentDomain.Document {
	entity := documentDomain.Document{}
	document, err := documents.GetDocument(docId)
	if err != nil {
		return entity
	}
	return document

}

type ApplicantDetails struct {
	ApplicantionId  string
	Email           string
	Name            string
	LatName         string
	ApplicantType   string
	Institution     domain.Institution
	Course          string
	DateTime        time.Time
	ApplicationStat applicationDomain.ApplicationStatus
	Stat            string
	Title           domain3.Title
}

type ApplicationDetail struct {
	ApplicantDetails   ApplicantDetails
	ApplicationStatus  string
	ApplicationStatues []domain4.GenericStatus
	Modifier           userDomain.User
	ModificationDate   time.Time
	Comment            string
	Course             domain2.Course
}

type documentDetails struct {
	DocumentType   string
	Status         string
	Document       userDomain.UserDocument
	DocumentStatus []domain4.GenericStatus
	Doc            documentDomain.Document
	UserId         string
	ApplicationId  string
	DocStatus      []documenStatus
}
type documenStatus struct {
	DocStatus documentDomain.DocumentStatus
	Status    string
	Admin     userDomain.User
}

func getdocumenStatus(documentId string) []documenStatus {
	var docEntity []documenStatus
	//var entity []documentDomain.DocumentStatus
	documentStatues, err := documents.GetdocumentStatues(documentId)
	if err != nil {
		return docEntity
	}
	for _, docentities := range documentStatues {
		stat, err := util.GetStatus(docentities.StatusId)
		if err != nil {
			fmt.Println("error reading status in .....")
		}
		/***Getting the user who modified the document stat**/
		admin, err := users.GetUser(docentities.ModifiedBy)
		if err != nil {
			fmt.Println("error reading the user who modified the document stat in .....")
		}
		docEntity = append(docEntity, documenStatus{docentities, stat.Name, admin})

	}
	return docEntity
}

func getDocumentDetails(user, applicatioId string) []documentDetails {
	var docDetails []documentDetails
	fmt.Println(user, "<<<<userId")
	userDocuments, err := users.GetUserDocuments(user)
	if err != nil {
		fmt.Println(err.Error(), " error reading userDocuments")
	}
	fmt.Println(userDocuments, "<<<<<userDocuments")

	status, err := util.GetStatuses()
	if err != nil {
		fmt.Println(err.Error(), " error reading status")
	}
	for _, userdocument := range userDocuments {
		docDetails = append(docDetails, documentDetails{getDocumentType(userdocument.DocumentId).DocumentTypename, getCurrentDocumentStatus(userdocument.DocumentId).Name, userdocument, status, getDocument(userdocument.DocumentId), user, applicatioId, getdocumenStatus(userdocument.DocumentId)})
	}
	return docDetails
}

/***this Method returns all the status obj of a document**/
func getDocumentStatus(documentId string) []domain4.GenericStatus {
	var statues []domain4.GenericStatus
	documentStatus, err := documents.GetdocumentStatues(documentId)
	if err != nil {
		fmt.Println(err.Error(), " error reading document status")
	}
	for _, status := range documentStatus {
		stat, err := util.GetStatus(status.StatusId)
		if err != nil {
			fmt.Println(err.Error(), " error reading statu of a document>>>", status.DocumentId)
		}
		statues = append(statues, stat)
	}
	return statues
}

func getCurrentDocumentStatus(documentId string) domain4.GenericStatus {
	var statues = domain4.GenericStatus{}
	documentStatus, err := documents.GetDocumentStatus(documentId)
	if err != nil {
		fmt.Println(err.Error(), " error reading document status")
		return statues
	}
	stat, err := util.GetStatus(documentStatus.StatusId)
	if err != nil {
		fmt.Println(err.Error(), " error reading statu of a document>>>")
		return statues
	}

	return stat
}

func getUser(userId string) userDomain.User {
	var entity = userDomain.User{}
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

func getSearchResult(userId, applicationId string) ApplicationDetail {
	var entity = ApplicationDetail{}
	for _, applicantDetails := range getApplicants() {

		userCourse, err := users.GetUserApplicationCourse(userId, applicationId)
		if err != nil {
			fmt.Println("error reading GetUserApplicationCourse in getSearchResult")
		}
		course, err := academics.GetCourse(userCourse.CourseId)
		if err != nil {
			fmt.Println("error reading course in getSearchResult")
		}

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
			return ApplicationDetail{applicantDetails, getStatus(applicationStatus.StatusId).Name, Statues, getUser(applicationStatus.ModifiedBy), applicationStatus.DateTime, applicationStatus.Comment, course}
		}
	}
	return entity
}
func getUserIntitution(userId, applicationId string) domain.Institution {
	var entity domain.Institution
	userInstitution, err := users.GetUserApplicationInstitution(userId, applicationId)
	if err != nil {
		fmt.Println("error reading UserApplicationInstitution in getUserIntitution")
		return entity
	} else if userInstitution.InstitutionId != "" {
		institution, err := institutions.GetInstitution(userInstitution.InstitutionId)
		if err != nil {
			fmt.Println("error reading Institution in getUserIntitution")
			return entity
		}
		return institution
	}
	return entity
}

func getApplicants() []ApplicantDetails {
	var applicantD []ApplicantDetails
	var applicant ApplicantDetails

	applications, err := applicationIO.GetApplications()
	//fmt.Println(" applications>>>> ", applications)
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

		userApplicationCourse, err := users.GetUserApplicationCourse(user.Email, application.Id)
		if err != nil {
			fmt.Println("error reading userApplicationCourse in getApplicants")
		}

		course, err := academics.GetCourse(userApplicationCourse.CourseId)
		if err != nil {
			fmt.Println("error reading course in getApplicants")
		}
		applicationStat, err := applicationIO.GetApplicationStatus(application.Id)
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
		usetTitle, err := users.GetUserDemographic(user.Email)
		if err != nil {
			fmt.Println("error reading UserTitle in getApplicants")
		}
		title, err := demographics.GetTitle(usetTitle.TitleId)
		if err != nil {
			fmt.Println("error reading Title in getApplicants")
		}

		applicant = ApplicantDetails{userApplication.ApplicationId, user.Email, user.FirstName, user.LastName, applicantType.Name, getUserIntitution(user.Email, application.Id), course.CourseName, userApplication.DateTime, applicationStat, status.Name, title}

		applicantD = append(applicantD, applicant)
	}
	return applicantD
}

//this method takes UserId and applicationId, so that it can find user's documents
func AdminApplicationDocumentsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clearCash(w)
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
				documentType, err := documents.GetDocumentType(document.DocumentId)
				if err != nil {
					fmt.Println("error reading documentType in AdminApplicationDocumentsHandler")
				}
				myDocumentList = append(myDocumentList, documentDetails{documentType.DocumentTypename, doc.DocumentStatus, document, documentsStatus, doc, userId, applicationId, getdocumenStatus(document.DocumentId)})
			}
		}

		type PageData struct {
			Applicant   []ApplicantDetails
			Document    []documentDetails
			Application ApplicationDetail
			Tab         string
			SubMenu     string
		}
		Data := PageData{getApplicants(), myDocumentList, getSearchResult(userId, applicationId), "application", ""}
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
		type PageData struct {
			Tab     string
			SubMenu string
		}
		data := PageData{"application", ""}
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
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

type MyUserApplication struct {
	Application     applicationDomain.Application
	User            userDomain.User
	UserApplication userDomain.UserApplication
}

func AdminApplicantHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		clearCash(w)
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		fmt.Println(email, "<<<<<email session")
		tab := app.Session.GetString(r.Context(), "tab")
		fmt.Println(tab, "<<<<<Tab session")

		documents := []documentDetails{}
		application := ApplicationDetail{}
		applicationDetails := []ApplicationHolder{}
		type Pagedate struct {
			Applicant         []ApplicantDetails
			Document          []documentDetails
			Application       ApplicationDetail
			Tab               string
			SubMenu           string
			Accordion         string
			ApplicationStatus []ApplicationHolder
		}
		Data := Pagedate{getApplicants(), documents, application, "applicant", "", tab, applicationDetails}
		files := []string{
			app.Path + "content/admin/admin_application3.html",
			//app.Path + "content/admin/admin_applicant.html",
			app.Path + "content/admin/template/sidebar.template.html",
			app.Path + "content/admin/template/navbar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		app.Session.Remove(r.Context(), "tab")
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
		clearCash(w)
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := users.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		fmt.Println("User ", user)
		type PageData struct {
			Tab     string
			SubMenu string
		}
		data := PageData{"dashboard", ""}
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
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}

}

func getDocumentType(documentId string) documentDomain.DocumentType {
	var docType documentDomain.DocumentType
	document, err := documents.GetDocument(documentId)
	if err != nil {
		return docType
	}
	documentType, err := documents.GetDocumentType(document.DocumentTypeId)
	if err != nil {
		return docType
	}
	return documentType
}
func clearCash(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//w.Header().Set("Pragma", "no-cache")
	//w.Header().Set("Expires", "0")
}
