package logout

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
)

func Logout(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", logoutHandler(app))
	return r
}

func logoutHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = app.Session.Destroy(r.Context())
		files := []string{
			app.Path + "base/login/login.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
