package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
	io "obas/src/io/subjects"
)

func Subjects(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", MatricSubjectsHandler(app))
	r.Get("/", UniversityCoursesHandler(app))
	return r
}

func UniversityCoursesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allcourses, err := io.GetUniversityCourses()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			courses []io.MatricSubjects
			name    string
		}
		data := PageData{allcourses, ""}

		files := []string{
			app.Path + "/subjects/subjects.page.html",
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

func MatricSubjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allsubjects, err := io.GetMatricSubjects()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			subjects []io.MatricSubjects
			name     string
		}
		data := PageData{allsubjects, ""}

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
