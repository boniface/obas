package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	"obas/src/config"
	io "obas/src/io/documents"
	"obas/src/io/log"
)

func DocumentsType(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", documentsTypeHandler(app))
	return r
}

func documentsTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alldocst, err := io.GetDocumentType()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			documentType []io.DocumentType
			name         string
		}
		data := PageData{alldocst, ""}

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
