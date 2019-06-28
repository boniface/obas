package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
)

func Subjects(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/matric", matricSubjectsHandler(app))
	r.Get("/university", universityCoursesHandler(app))
	return r
}

func universityCoursesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allcourses, err := io.GetUniversityCourses()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//courses []io.UniversityCourses
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/subjects/subjects.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/contents.page.html",
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
		//allmatrics, err := io.GetMatricSubjects()

		//if err != nil {
		//	fmt.Println(" IS this Error Called ", err)
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//matrics []io.MatricSubjects
			name string
		}
		data := PageData{""}

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
