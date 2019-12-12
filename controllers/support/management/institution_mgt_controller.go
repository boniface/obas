package management

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	institutionDomain "obas/domain/institutions"
	institutionIO "obas/io/institutions"
)

func InstitutionManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", InstitutionManagementHandler(app))
	r.Post("/add", AddInstitutionHandler(app))

	return r
}

func AddInstitutionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//email := app.Session.GetString(r.Context(), "userId")
		//token := app.Session.GetString(r.Context(), "token")
		//if email == "" || token == "" {
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		r.ParseForm()
		institutionName := r.PostFormValue("institutionName")
		institutionType := r.PostFormValue("institutionType")

		institution := institutionDomain.Institution{"", institutionType, institutionName}
		app.InfoLog.Println("Institution to save: ", institution)
		savedInstitution, err := institutionIO.CreateInstitution(institution)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Create location response is ", savedInstitution)
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func InstitutionManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type InstitutionHolder struct {
			InstitutionName string
			InstitutionType string
		}

		var institutions []institutionDomain.Institution
		var institutionsHolder []InstitutionHolder
		institutionTypes, err := institutionIO.GetInstitutionTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else {
			institutions, err = institutionIO.GetInstitutions()
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				for _, institution := range institutions {
					institutionTypeName := getInstitutionTypeName(institution, institutionTypes)
					institutionsHolder = append(institutionsHolder, InstitutionHolder{institution.Name, institutionTypeName})
				}
			}
		}

		type PageData struct {
			InstitutionTypes   []institutionDomain.InstitutionTypes
			InstitutionsHolder []InstitutionHolder
		}

		data := PageData{institutionTypes, institutionsHolder}

		files := []string{
			app.Path + "content/tech/tech_admin_institution.html",
			app.Path + "content/tech/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func getInstitutionTypeName(institution institutionDomain.Institution, institutionTypes []institutionDomain.InstitutionTypes) string {
	var institutionTypeName string
	for _, institutionType := range institutionTypes {
		if institution.InstitutionTypeId == institutionType.Id {
			institutionTypeName = institutionType.Name
		}
	}
	return institutionTypeName
}
