package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
	io "obas/src/io/address"
)

func Contacts(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", contactHandler(app))
	return r
}

func contactHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allcontact, err := io.GetContactTypes()
		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			contact []io.ContactType
			name    string
		}
		data := PageData{allcontact, ""}
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
