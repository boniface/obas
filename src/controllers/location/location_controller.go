package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	"obas/src/config"
	io "obas/src/io/location"
)

func Locations(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", locationsHandler(app))
	return r
}

func locationsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alllocations, err := io.GetLocations()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			locations []io.Location
			name      string
		}

		data := PageData{alllocations, ""}

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
