package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
	io "obas/src/io/address"
)

func Address(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", addressHandler(app))
	return r
}
func addressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alladdresses, err := io.GetAddresses()
		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			address []io.AddressType
			name    string
		}
		data := PageData{alladdresses, ""}
		files := []string{
			app.Path + "",
		}
		ts, err := template.ParseFiles(files...) //need to inspect here**template OR templates
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
