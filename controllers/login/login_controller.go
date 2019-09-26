package login

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	"obas/io/login"
)

// Route Path
func Login(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", loginHome(app))
	r.Post("/login", loginHandler(app))
	r.Post("/accounts", getAccountsHandler(app))
	r.Get("/password", passwordHandler(app))
	r.Get("/verify", passwordHandler(app))
	return r
}

func loginHome(app *config.Env) http.HandlerFunc {
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

func loginHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		loginToken, err := login.DoLogin(email, password)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
		}
		app.InfoLog.Println("Login is successful. Result is ", loginToken)
		http.Redirect(w, r, "/user/student", 301)
	}
}

func logout(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))

	}
}

func forgotPassword(app *config.Env) http.HandlerFunc {

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

func getAccountsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
