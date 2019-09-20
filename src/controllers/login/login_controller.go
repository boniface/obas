package login

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
	"obas/src/io/login"
)

// Route Path
func Login(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", loginHandler(app))
	r.Post("/isUserRegistered", UserRegistered(app))
	r.Get("/password", passwordHandler(app))
	r.Get("/password/forgot", ForgotPassword(app))
	return r
}

func loginHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "base/login/login.page.html",
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

func logout(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))

	}
}

func ForgotPassword(app *config.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))

	}
}

func passwordHandler(app *config.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		files := []string{
			app.Path + "/login/password.page.html",
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

func UserRegistered(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var files []string
		r.ParseForm()
		email := r.PostFormValue("email")
		userReg, err := login.IsUserRegistered(email)
		if err != nil {
			fmt.Println(" The Error ", err)
		}
		switch 2 {
		case 0:
			{
				http.Redirect(w, r, "/login", 301)
			}
		case 1:
			{
				http.Redirect(w, r, "/login", 301)
			}
		default:
			{

				fmt.Println(" We got this user ", userReg)
				files = []string{
					app.Path + "Dashboard Here",
				}
				ts, err := template.ParseFiles(files...)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					return
				}
				err = ts.Execute(w, userReg)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				}
			}
		}
	}
}
