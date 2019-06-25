package location

import (
	"net/http"
	"obas/src/config"
	io "obas/src/io/location"
)

func LocationTypes(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", locationTypesHandler(app))
	return r
}

func locationTypesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alllocationtypes, err := io.GetLocationTypes()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			locationtypes []io.LocationType
			name          string
		}

		data := PageData{alllocationtypes, ""}

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
