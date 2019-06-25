package application

import (
	"html/template"
	"net/http"
	"obas/src/config"
	io "obas/src/io/application"
)

func applicationStatues(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", applicationStatu(app))
	return r
}
func applicationStatu(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allapplications, err := io.GetApplicationStatuses()
		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			applicationRes []io.ApplicationStatus
			name           string
		}
		data := PageData{allapplications, ""}
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
