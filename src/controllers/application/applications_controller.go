package controllers

import (
	"github.com/go-chi/chi"
	"html/template"

	"net/http"
	"obas/src/config"
	io "obas/src/io/application"
)

func Applications(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", ApplicationTypeHandler(app))
	r.Get("/", ApplicationResultHandler(app))
	r.Get("/", ApplicationStatusHandler(app))
	return r
}

func ApplicationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allApplicationsType, err := io.GetApplicationTypes()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			ApplicationsType []io.ApplicationType
			name             string
		}
		data := PageData{allApplicationsType, ""}

		files := []string{
			app.Path + "/html/application/application.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
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

func ApplicationResultHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allApplicationResult, err := io.GetApplicationResultes()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			ApplicationResults []io.ApplicationResult
			name               string
		}
		data := PageData{allApplicationResult, ""}

		files := []string{
			app.Path + "/html/application/application.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
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

func ApplicationStatusHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allApplicationResult, err := io.GetApplicationResultes()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			ApplicationResults []io.ApplicationResult
			name               string
		}
		data := PageData{allApplicationResult, ""}

		files := []string{
			app.Path + "/html/application/application.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebar.page.html",
			app.Path + "/base/footer.page.html",
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
