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
	r.Get("/error", loginError(app))
	r.Get("/forgetpassword", forgetPasswordHandler(app))
	r.Get("/passwordreset/{resetkey}", passwordResetHandler(app))
	r.Post("/login", loginHandler(app))
	r.Post("/processforgetpassword", processForgetPasswordHandler(app))
	r.Post("/accounts", getAccountsHandler(app))
	return r
}

func LogOut(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", logoutHandler(app))
	return r
}

func logoutHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = app.Session.Destroy(r.Context())
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

func loginError(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "base/login/login.page_Error.html",
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

func forgetPasswordHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "base/login/forgotpassword.page.html",
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

func passwordResetHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resetKey := chi.URLParam(r, "resetkey")
		result, err := login.DoReset(resetKey)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		type PageData struct {
			Email string
			Title string
			Info  string
		}
		var redirect, title, info string
		if result {
			redirect = app.Path + "base/login/login.page.html"
			title = "Password Reset Successful"
			info = "A temporary password has been sent to your email. Please log in with the temporary password."
		} else {
			redirect = app.Path + "base/login/forgotpassword.page.html"
			title = "Password Reset NOT Successful"
			info = "An error occurred. Please try again or contact administrator."
		}
		data := PageData{"do_not_reply@ict.cput.ac.za", title, info}
		files := []string{
			redirect,
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func processForgetPasswordHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("email")
		result, err := login.DoForgetPassword(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login/forgetpassword", 301)
		}
		app.InfoLog.Println("Login is successful. Result is ", result)
		http.Redirect(w, r, "/login", 301)
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
			app.Session.Put(r.Context(), "message", "Wrong Credentials!")
			http.Redirect(w, r, "/", 301)
		}

		app.Session.Put(r.Context(), "email", loginToken.Email)
		app.Session.Put(r.Context(), "token", loginToken.Token)
		app.InfoLog.Println("Login is successful. Result is ", loginToken)
		http.Redirect(w, r, "/users/student", 301)
		app.Session.Cookie.Name = "UserID"
		app.Session.Put(r.Context(), "UserID", loginToken.Email)
		app.Session.Put(r.Context(), "Token", loginToken.Token)
		http.Redirect(w, r, "/", 301)
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

func getAccountsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
