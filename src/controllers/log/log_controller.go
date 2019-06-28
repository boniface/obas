package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
)

func Logs(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/log", logsHandler(app))
	return r
}

func logsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//alllogs, err := io.GetLogEvents()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//logs []io.LogEvent
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/log/log.page.html",
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
