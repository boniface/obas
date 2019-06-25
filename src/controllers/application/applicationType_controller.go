package application

import (
	"html/template"
	"net/http"
	"obas/src/config"
	io "obas/src/io/application"
)

func applicationTypes(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", applicationTypeHandler(app))
	return r
}
func applicationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allApplicationType, err := io.GetApplicationTypes()

		if err != nil {
			app.ServerError(w, err)
		}
		type PageData struct {
			applicationT []io.ApplicationType
			name         string
		}
		data := PageData{allApplicationType, ""}
		filees := []string{
			app.Path + "",
		}
		ts, err := template.ParseFiles(filees...)
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
