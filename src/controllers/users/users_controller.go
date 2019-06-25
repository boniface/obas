package controllers

import (
	"github.com/go-chi/chi"

	"net/http"
	"obas/src/config"
	io "obas/src/io/users"
)

func Users(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", AdminHandler(app))
	r.Get("/", ProcessingStatusTypeHandler(app))
	r.Get("/", StudentApplicationStatusHandler(app))
	r.Get("/", StudentContactsHandler(app))
	r.Get("/", StudentDemographicsHandler(app))
	r.Get("/", StudentDocumentsHandler(app))
	r.Get("/", StudentProfileHandler(app))
	r.Get("/", StudentResultsHandler(app))
	return r
}

func AdminHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allAdmin, err := io.GetAdmins()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			courses []io.Admin
			name    string
		}
		data := PageData{allAdmin, ""}

		files := []string{
			app.Path + "", //maybe the admin base page should go here
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
		}
		ts, err := templates.ParseFiles(files...)
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
		allProcess, err := io.GetProcessingStatusTypes()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			subjects []io.ProcessingStatusType
			name     string
		}
		data := PageData{allProcess, ""}

		files := []string{
			app.Path + "", //maybe another base page should go here
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
		}
		ts, err := templates.ParseFiles(files...)
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
		allApplications, err := io.GetStudentApplicationStatuses()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			subjects []io.StudentApplicationStatus
			name     string
		}
		data := PageData{allApplications, ""}

		files := []string{
			app.Path + "",
		}
		ts, err := templates.ParseFiles(files...)
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

func StudentContactsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allStudentContacts, err := io.GetStudentContacts()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			subjects []io.StudentContacts
			name     string
		}
		data := PageData{allStudentContacts, ""}

		files := []string{
			app.Path + "",
		}
		ts, err := templates.ParseFiles(files...)
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
func StudentDemographicsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allStudentDemographics, err := io.GetStudentDemographics()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			subjects []io.StudentDemographics
			name     string
		}
		data := PageData{allStudentDemographics, ""}

		files := []string{
			app.Path + "",
		}
		ts, err := templates.ParseFiles(files...)
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
func StudentDocumentsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allStudentDocuments, err := io.GetStudentDocuments()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			subjects []io.StudentDocuments
			name     string
		}
		data := PageData{allStudentDocuments, ""}

		files := []string{
			app.Path + "",
		}
		ts, err := templates.ParseFiles(files...)
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
func StudentProfileHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allStudentProfiles, err := io.GetStudentProfiles()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			subjects []io.StudentProfiles
			name     string
		}
		data := PageData{allStudentProfiles, ""}

		files := []string{
			app.Path + "",
		}
		ts, err := templates.ParseFiles(files...)
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
		allStudentResults, err := io.GetStudentResults()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			subjects []io.StudentResults
			name     string
		}
		data := PageData{allStudentResults, ""}

		files := []string{
			app.Path + "",
		}
		ts, err := templates.ParseFiles(files...)
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
