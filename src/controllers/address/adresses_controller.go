package controllers

import (
	"github.com/go-chi/chi"
	"html/template"

	"net/http"
	"obas/src/config"
)

//noinspection ALL
func Addresses(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/all", AddressTypeHandler(app))
	r.Get("/contact/all", ContactTypeTypeHandler(app))
	return r
}

//noinspection ALL
func AddressTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allAddresses, err := io.GetAddresses()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//addresses []io.AddressType
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/address/address.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
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
		//allContacts, err := io.GetContactTypes()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//contacts []io.ContactType
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/address/address.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
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
