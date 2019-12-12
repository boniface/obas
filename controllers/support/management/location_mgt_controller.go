package management

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	domain "obas/domain/location"
	locationIO "obas/io/location"
)

func LocationManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", LocationManagementHandler(app))

	r.Post("/type/add", AddLocationTypeHandler(app))
	r.Post("/add", AddLocationHandler(app))

	return r
}

func AddLocationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		locationTypeName := r.PostFormValue("locationTypeName")
		locationType := domain.LocationType{"", locationTypeName, ""}
		app.InfoLog.Println("Location type to save: ", locationType)
		savedLocationType, err := locationIO.CreateLocationType(locationType)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Create location type response is ", savedLocationType)
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

func AddLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//email := app.Session.GetString(r.Context(), "userId")
		//token := app.Session.GetString(r.Context(), "token")
		//if email == "" || token == "" {
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		r.ParseForm()
		locationName := r.PostFormValue("locationName")
		latitude := r.PostFormValue("latitude")
		longitude := r.PostFormValue("longitude")
		locationType := r.PostFormValue("locationType")
		locationParent := r.PostFormValue("locationParent")
		location := domain.Location{"", locationType, locationName, latitude, longitude, locationParent}
		app.InfoLog.Println("Location to save: ", location)
		savedLocation, err := locationIO.CreateLocation(location)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Create location response is ", savedLocation)
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

func LocationManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var locations []domain.Location
		locationTypes, err := locationIO.GetLocationTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else {
			locations, err = locationIO.GetLocations()
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}

		type PageData struct {
			LocationTypes []domain.LocationType
			Locations     []domain.Location
		}

		data := PageData{locationTypes, locations}

		files := []string{
			app.Path + "content/tech/tech_admin_loc.html",
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
