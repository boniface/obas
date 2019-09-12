package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
)

func Locations(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", LocationsHandler(app))
	r.Get("/type", LocationTypesHandler(app))
	return r
}

func LocationsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//alllocations, err := io.GetLocations()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//locations []io.Location
			name string
		}

		data := PageData{""}

		files := []string{
			app.Path + "/location/location.page.html",
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

func LocationTypesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//alllocationTypes, err := io.GetLocationTypes()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//locationsType []io.LocationType
			name string
		}

		data := PageData{""}

		files := []string{
			app.Path + "/location/location.page.html",
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
