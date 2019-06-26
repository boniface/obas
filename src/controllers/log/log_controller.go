package controllers

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"obas/src/config"
	"obas/src/io/log"
)

func Logs(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", logsHandler(app))
	return r
}

func logsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alllogs, err := io.GetLogEvents()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			logs []io.LogEvent
			name string
		}
		data := PageData{alllogs, ""}

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
