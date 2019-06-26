package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
	domain "obas/src/domain/subjects"
	io "obas/src/io/subjects"
)

func Subjects(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", SubjectsHandler(app))
	r.Get("/matric", matricSubjectsHandler(app))
	r.Get("/university", universityCoursesHandler(app))
	return r
}

func SubjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allsubjects, err := io.GetSubjects()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			subjects []domain.Subjects
			name     string
		}
		data := PageData{allsubjects, ""}

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

func universityCoursesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allcourses, err := io.GetUniversityCourses()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			courses []domain.UniversityCourses
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

func matricSubjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allsubjects, err := io.GetMatricSubjects()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			subjects []io.Subjects
			name     string
		}
		data := PageData{allsubjects, ""}

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
