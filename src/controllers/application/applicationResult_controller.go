package application

import (
	"html/template"
	"net/http"
	"obas/src/config"
	io "obas/src/io/application"
)

func applicationResults(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", applicationHandler(app))
	return r
}
func applicationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		allapplicationResults, err := io.GetApplicationResultes()
		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			applicationR []io.ApplicationResult
			name         string
		}
		data := PageData{allapplicationResults, "/"}
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
