package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	"obas/src/config"
	io "obas/src/io/registration"
)

func Registrations(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", registrationsHandler(app))
	return r
}

func registrationsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allregistrations, err := io.GetRegister()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			registrations []io.Register
			name          string
		}

		data := PageData{allregistrations, ""}

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
