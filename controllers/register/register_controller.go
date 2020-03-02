package register

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
			app.Path + "base/register/register.page.html",
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

		println(email, "register with this email")

		registered, err := login.DoRegister(email)

		println("registration is: ", registered)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Title string
			Info  string
		}
		var redirect, title, info string
		if registered {
			redirect = app.Path + "base/login/login.page.html"
			title = "Registration is successful"
			info = "A temporary password has been sent to your email. Please log in with the temporary password."
		} else {
			redirect = app.Path + "base/register/register.page.html"
			title = "Registration NOT successful"
			info = "An error occurred. Please try again or contact administrator."
		}
		data := PageData{title, info}
		files := []string{
			redirect,
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
