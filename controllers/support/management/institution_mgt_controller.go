package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	institutionDomain "obas/domain/institutions"
	locationDomain "obas/domain/location"
	institutionIO "obas/io/institutions"
	location "obas/io/location"
	"obas/util"
)

func InstitutionManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", InstitutionManagementHandler(app))
	r.Get("/type/delete/{formData}", DeleteTypeHandler(app))
	r.Get("/delete/institution/{institutionId}", DeleteInstitutionTypeHandler(app))
	r.Post("/type/add", AddInstitutiontypeHandler(app))
	r.Post("/type/edit", EditTypeHandler(app))
	r.Post("/add", AddTypeHandler(app))
	r.Post("/save/institution-location", SaveInstitutionLocationHandler(app))
	r.Get("/delete-institutionLocation/{institutionLocationId}", DeleteInstitutionLocationHandler(app))
	return r
}

func DeleteInstitutionLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resetKey := chi.URLParam(r, "institutionLocationId")

		if resetKey != "" {
			institutionLocationObject, err := institutionIO.ReadInstitutionLocation(resetKey)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
			if institutionLocationObject.InstitutionId != "" {
				_, err := institutionIO.DeleteInstitutionLocation(institutionLocationObject)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				}
			}
		}
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func SaveInstitutionLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		locationId := r.PostFormValue("townId")
		institutionId := r.PostFormValue("institution")
		longitude := r.PostFormValue("longitude")
		latitude := r.PostFormValue("latitude")

		if locationId != "" || institutionId != "" || longitude != "" || latitude != "" {
			institutionLocation := institutionDomain.InstitutionLocation{institutionId, locationId, longitude, latitude}
			_, err := institutionIO.CreateInstitutionLocation(institutionLocation)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func DeleteInstitutionTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resetKey := chi.URLParam(r, "institutionId")
		institutionObject, err := institutionIO.GetInstitution(resetKey)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if institutionObject.Id != "" {
			_, err := institutionIO.DeleteInstitution(institutionObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func AddTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		institutionName := r.PostFormValue("institutionName")
		institutionType := r.PostFormValue("institutionType")
		fmt.Println(institutionName, "  <<<<institutionName    institutionType", institutionType)
		if institutionName != "" || institutionType != "" {
			institu := institutionDomain.Institution{"", institutionType, institutionName}
			_, err := institutionIO.CreateInstitution(institu)
			if err != nil {
				fmt.Println("error creating institution")
				app.ErrorLog.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func EditTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//email := app.Session.GetString(r.Context(), "userId")
		//token := app.Session.GetString(r.Context(), "token")
		//if email == "" || token == "" {
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		r.ParseForm()
		institutionTypeId := r.PostFormValue("Id")
		institutionTypeName := r.PostFormValue("Name")
		institutionTypeDescription := r.PostFormValue("Description")

		if institutionTypeId != "" || institutionTypeName != "" || institutionTypeDescription != "" {

			institutionTypeObject := institutionDomain.InstitutionTypes{institutionTypeId, institutionTypeName, institutionTypeDescription}
			fmt.Print(institutionTypeObject)
			_, err := institutionIO.DeleteInstitutionType(institutionTypeObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
			_, erro := institutionIO.CreateInstitutionType(institutionTypeObject)
			if erro != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func DeleteTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resetKey := chi.URLParam(r, "formData")
		institutionObject, err := institutionIO.GetInstitutionType(resetKey)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if institutionObject.Id != "" {
			_, err := institutionIO.DeleteInstitutionType(institutionObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())

			}
		}
		http.Redirect(w, r, "/support/management/institution", 301)
	}

}

func AddInstitutiontypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//email := app.Session.GetString(r.Context(), "userId")
		//token := app.Session.GetString(r.Context(), "token")
		//if email == "" || token == "" {
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		r.ParseForm()
		institutionName := r.PostFormValue("Name")
		institutionDescription := r.PostFormValue("Description")

		institution := institutionDomain.InstitutionTypes{"", institutionName, institutionDescription}
		app.InfoLog.Println("Institution to save: ", institution)
		savedInstitutionTypes, err := institutionIO.CreateInstitutionType(institution)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Create location response is ", savedInstitutionTypes)
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func InstitutionManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type InstitutionHolder struct {
			Id              string
			InstitutionName string
			InstitutionType string
		}
		type InstitutionLocHolder struct {
			Id          string
			Institution string
			Town        string
			Longitude   string
			Latitude    string
		}

		var institutionLocation []InstitutionLocHolder
		var institutions []institutionDomain.Institution
		var institutionsHolder []InstitutionHolder

		institutsLocation, err := institutionIO.ReadInstitutionLocations()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else {
			for _, institutionLoca := range institutsLocation {
				fmt.Println("error in reading townNamw in InstitutionManagementHandler method", institutionLoca.LocationId)

				institutionName, errr := institutionIO.GetInstitution(institutionLoca.InstitutionId)
				if errr != nil {
					fmt.Println("error in reading institutionName in InstitutionManagementHandler method")
					app.ErrorLog.Println(errr.Error())
				}
				townNamw, err := location.GetLocation(institutionLoca.LocationId)
				fmt.Println("error in reading townNamw in InstitutionManagementHandler method", townNamw)
				if err != nil {
					fmt.Println("error in reading townNamw in InstitutionManagementHandler method")
					app.ErrorLog.Println(err.Error())
				}
				institutionLocation = append(institutionLocation, InstitutionLocHolder{institutionLoca.InstitutionId, institutionName.Name, townNamw.Name, institutionLoca.Longitude, institutionLoca.Latitude})
			}
		}

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
					institutionsHolder = append(institutionsHolder, InstitutionHolder{institution.Id, institution.Name, institutionTypeName})
				}
			}
		}
		provinces, _ := util.GetProvinces()

		type PageData struct {
			InstitutionTypes    []institutionDomain.InstitutionTypes
			InstitutionsHolder  []InstitutionHolder
			Provinces           []locationDomain.Location
			InstitutionLocation []InstitutionLocHolder
		}

		data := PageData{institutionTypes, institutionsHolder, provinces, institutionLocation}

		files := []string{
			app.Path + "content/tech/tech_admin_institution.html",
			app.Path + "content/tech/template/sidebar.template.html",
			app.Path + "base/template/form/location-form.template.html",
			app.Path + "base/template/form/institution-form.template.html",
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
