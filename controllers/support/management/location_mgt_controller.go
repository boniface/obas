package management

import (
	"fmt"
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
	r.Post("/type/edit", EditLocationTypeHandler(app))
	r.Post("/location/update", EditLocationHandler(app))
	r.Get("/delete/location/{resetkey}", DeleteLocationHandler(app))
	return r
}

func EditLocationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/***
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		*/
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()

		locationName := r.PostFormValue("Name")
		locationId := r.PostFormValue("Id")
		code := r.PostFormValue("Code")

		fmt.Println(locationName, "<<<<<locationName||locationId", locationId)
		if locationName != "" || locationId != "" {
			locationType := domain.LocationType{locationId, locationName, code}
			fmt.Println(locationType, "<<<<<locationType type")
			_, err := locationIO.UpdateLocationType(locationType)
			if err != nil {
				fmt.Println("error updating loccation type")
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

func EditLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		locationId := r.PostFormValue("LocationId")
		locationName := r.PostFormValue("Name")
		latitude := r.PostFormValue("Latitude")
		longitude := r.PostFormValue("Longitude")
		locationType := r.PostFormValue("locationType")
		locationParent := r.PostFormValue("locationParent")
		location := domain.Location{locationId, locationType, locationName, latitude, longitude, locationParent}

		fmt.Println(location)
		_, err := locationIO.UpdateLocation(location)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

func DeleteLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locationId := chi.URLParam(r, "resetkey")
		_ = app.Session.Destroy(r.Context())
		fmt.Print(locationId)
		locationObject, err := locationIO.GetLocation(locationId)
		fmt.Print(locationObject)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if locationObject.LocationId != "" {
			_, err := locationIO.DeleteLocation(locationObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

func AddLocationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = app.Session.Destroy(r.Context())
		/***
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		*/

		r.ParseForm()
		locationTypeName := r.PostFormValue("locationTypeName")
		if locationTypeName != "" {
			locationType := domain.LocationType{"", locationTypeName, ""}
			app.InfoLog.Println("Location type to save: ", locationType)
			_, err := locationIO.CreateLocationType(locationType)

			if err != nil {
				app.ErrorLog.Println(err.Error())

			}
		}

		app.Session.Put(r.Context(), "tab", "tab1")

		//app.InfoLog.Println("Create location type response is ", savedLocationType)
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

func AddLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = app.Session.Destroy(r.Context())
		//email := app.Session.GetString(r.Context(), "userId")
		//token := app.Session.GetString(r.Context(), "token")
		//if email == "" || token == "" {
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		r.ParseForm()
		locationName := r.PostFormValue("locationName")
		latitude := r.PostFormValue("Latitude")
		longitude := r.PostFormValue("Longitude")
		locationType := r.PostFormValue("LocationType")
		locationParent := r.PostFormValue("LocationParent")
		location := domain.Location{"", locationType, locationName, latitude, longitude, locationParent}
		app.InfoLog.Println("Location to save: ", location)
		savedLocation, err := locationIO.CreateLocation(location)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		app.InfoLog.Println("Create location response is ", savedLocation)
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

type MyLocations struct {
	Location       domain.Location
	LocationType   domain.LocationType
	ParentLocation domain.Location
}

func ReadlocationType(locationTypeId string) domain.LocationType {
	entity := domain.LocationType{}
	locationType, erro := locationIO.GetLocationType(locationTypeId)
	if erro != nil {
		return entity
	}
	return locationType
}
func ReadParentlocation(locationParentId string) domain.Location {
	entity := domain.Location{}
	locationType, erro := locationIO.GetLocation(locationParentId)
	if erro != nil {
		return entity
	}
	return locationType
}
func LocationManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var myLocations []MyLocations

		var locations []domain.Location
		locationTypes, err := locationIO.GetLocationTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else {
			locations, err = locationIO.GetLocations()
			for _, value := range locations {
				myLocations = append(myLocations, MyLocations{value, ReadlocationType(value.LocationTypeId), ReadParentlocation(value.LocationParentId)})
			}
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		type PageData struct {
			LocationTypes []domain.LocationType
			Locations     []MyLocations
			Location      []domain.Location
			Tab           string
			SubTab        string
		}
		data := PageData{locationTypes, myLocations, locations, "location", ""}
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
