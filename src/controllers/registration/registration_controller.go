package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
	io "obas/src/io/registration"
)

func Registrations(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/registration", registrationsHandler(app))
	return r
}

func registrationsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allregistrations, err := io.GetRegisters()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			registrations []io.Register
			name          string
		}

		data := PageData{allregistrations, ""}

		files := []string{
			app.Path + "/registration/registration.page.html",
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
