package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
	io "obas/src/io/demographics"
)

func demographics(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", GendersHandler(app))
	r.Get("/", RacesHandler(app))
	r.Get("/", RolesHandler(app))
	r.Get("/", TitlesHandler(app))
	return r
}

func GendersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allgenders, err := io.GetGenders()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			genders []io.Genders
			name    string
		}
		data := PageData{allgenders, ""}

		files := []string{
			app.Path + "/demographics/demographics.page.html",
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

func RacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allraces, err := io.GetRaces()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			Races []io.Races
			name  string
		}
		data := PageData{allraces, ""}

		files := []string{
			app.Path + "/demographics/demographics.page.html",
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

func RolesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allroles, err := io.GetRoles()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			Roles []io.Roles
			name  string
		}
		data := PageData{allroles, ""}

		files := []string{
			app.Path + "/demographics/demographics.page.html",
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
func TitlesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alltitles, err := io.GetTitles()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			Titles []io.Titles
			name   string
		}
		data := PageData{alltitles, ""}

		files := []string{
			app.Path + "",
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
