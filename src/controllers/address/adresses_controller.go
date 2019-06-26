package controllers

import (
	"github.com/go-chi/chi"
	"html/template"

	"net/http"
	"obas/src/config"
	io "obas/src/io/address"
)

func Addresses(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", AddressTypeHandler(app))
	r.Get("/", ContactTypeTypeHandler(app))
	return r
}

func AddressTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allAddresses, err := io.GetAddresses()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			addresses []io.AddressType
			name      string
		}
		data := PageData{allAddresses, ""}

		files := []string{
			app.Path + "/html/address/address.page.html",
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

func ContactTypeTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allContacts, err := io.GetContactTypes()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			contacts []io.ContactType
			name     string
		}
		data := PageData{allContacts, ""}

		files := []string{
			app.Path + "/html/address/address.page.html",
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
