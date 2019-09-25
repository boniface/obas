package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	"obas/io/login"
)

// Route Path
func Register(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", RegisterHome(app))
	r.Post("/register", RegisterHandler(app))
	return r
}

func RegisterHome(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "/register/register.page.html",
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

func RegisterHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("email")
		registered, err := login.DoRegister(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Registration is ", registered)
		http.Redirect(w, r, "/login", 301)
	}
}
