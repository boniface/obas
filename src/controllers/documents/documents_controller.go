package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/src/config"
	io "obas/src/io/documents"
)

func Documents(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", DocumentsHandler(app))
	r.Get("/", DocumentsTypeHandler(app))
	return r
}

func DocumentsTypeHandler(app *config.Env) http.HandlerFunc {
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
			app.Path + "/html/documents/documents.page.html",
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

func DocumentsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alldocsType, err := io.GetDocumentTypes()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			documentsType []io.DocumentType
			name          string
		}
		data := PageData{alldocsType, ""}

		files := []string{
			app.Path + "/html/documents/documents.page.html",
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
