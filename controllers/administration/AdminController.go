package controller

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
)

func Admin(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", AdminHandler(app))
	return r
}

func AdminHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/**email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := users_io.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		fmt.Println("User ", user)*/
		files := []string{
			app.Path + "content/admin/admin_dashboard.page.html",
			app.Path + "content/admin/template/sidebar.template.html",
			app.Path + "content/admin/template/navbar.template.html",
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