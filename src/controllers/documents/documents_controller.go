package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	"obas/src/config"
	io "obas/src/io/documents"
	"obas/src/io/log"
)

func Documents(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", documentsHandler(app))
	return r
}

func documentsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alldocs, err := io.GetDocuments()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			documents []io.Documents
			name      string
		}
		data := PageData{alldocs, ""}

		files := []string{
			app.Path + "",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, "", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
