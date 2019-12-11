package management

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
)

func UserManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", UserManagementHandler(app))

	return r
}

func UserManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		files := []string{
			app.Path + "content/tech/tech_admin_users.html",
			app.Path + "content/tech/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
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
