package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	locationHelper "obas/controllers/location"
	academicsDomain "obas/domain/academics"
	institutionDomain "obas/domain/institutions"
	domain "obas/domain/location"
	"obas/io/academics"
	"obas/io/address"
	institutionIO "obas/io/institutions"
)

func InstitutionManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", InstitutionManagementHandler(app))
	r.Get("/institution", GetInstitutionHandler(app))
	r.Get("/location", GetLocationHandler(app))
	r.Get("/address", GetAddressHandler(app))
	r.Get("/course", GetCourseHandler(app))
	r.Get("/type/delete/{formData}", DeleteTypeHandler(app))
	r.Get("/delete/institution/{institutionId}", DeleteInstitutionHandler(app))
	r.Get("/delete-institutionLocation/{institutionLocationId}", DeleteInstitutionLocationHandler(app))
	r.Get("/delete-institutionCourse/{institutionId}/{CourseId}", DeleteInstitutionCourseHandler(app))
	r.Get("/delete-institutionAddress/{InstitutionAddressId}/{AddressTypeId}", DeleteInstitutionAddressHandler(app))
	r.Post("/type/add", AddInstitutiontypeHandler(app))
	r.Post("/type/edit", EditTypeHandler(app))
	r.Post("/institution/update", UpdateInstitutionHandler(app))
	r.Post("/location/update", UpdateInstitutionHandler(app))
	r.Post("/add", AddInstitutionHandler(app))
	r.Post("/save/institution-location", SaveInstitutionLocationHandler(app))
	r.Post("/save/institution-course", SaveInstitutionCourseHandler(app))
	r.Post("/save/institution-address", SaveInstitutionAddressHandler(app))
	//
	return r
}

func InstitutionManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		type MYPageData struct {
			Tab    string
			SubTab string
		}
		data := MYPageData{"dashboard", "X"}
		fmt.Println(data, "<<<<<data")
		files := []string{
			app.Path + "content/tech/tech_dashboard.html",
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

func GetCourseHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		var institutionsCourseHolder []InstitutionCourseHolder

		/**Reading all the institution course and their type**/
		institutionsCourseHolder = GetInstitutionCourse()

		/**Reading all the course**/
		course, err := academics.GetAllCourses()
		if err != nil {
			app.InfoLog.Println(err.Error(), "error reading courses")
		}
		/**reading institution Type**/
		institutionType, err := institutionIO.GetInstitutionTypes()
		if err != nil {
			app.InfoLog.Println(err.Error(), "error reading institutionType")
		}

		type PageData struct {
			Tab               string
			SubTab            string
			InstitutionTypes  []institutionDomain.InstitutionTypes
			Courses           []academicsDomain.Course
			InstitutionCourse []InstitutionCourseHolder
		}
		data := PageData{"institution", "course", institutionType, course, institutionsCourseHolder}

		files := []string{
			app.Path + "content/tech/institution/course.html",
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

func GetAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		var institutionsAddressHolder []InstitutionAddressHolder

		/**Reading all the institution location and their type**/
		institutionsAddressHolder = GetInstitutionAddress()

		/**reading institution Type**/
		institutionType, err := institutionIO.GetInstitutionTypes()
		if err != nil {
			app.InfoLog.Println(err.Error(), "error reading institutionType")
		}

		/**reading Address Type**/
		addressType, err := address.GetAddressTypes()
		if err != nil {
			app.InfoLog.Println(err.Error(), "error reading institutionType")
		}

		type PageData struct {
			Tab                string
			SubTab             string
			InstitutionAddress []InstitutionAddressHolder
			AddressTypes       []address.AddressType
			InstitutionTypes   []institutionDomain.InstitutionTypes
		}

		data := PageData{"institution", "address", institutionsAddressHolder, addressType, institutionType}

		files := []string{
			app.Path + "content/tech/institution/address.html",
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

func GetLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		/**Reading all the institution location and their type**/
		institutionsLocationHolder := GetInstitutionLocation()

		/**reading institution Type**/
		institutionType, err := institutionIO.GetInstitutionTypes()
		if err != nil {
			app.InfoLog.Println(err.Error(), "error reading institutionType")
		}
		/**Getting all the provinces **/
		provinces, _ := locationHelper.GetProvinces(app)

		type PageData struct {
			Provinces           []domain.Location
			Tab                 string
			SubTab              string
			InstitutionLocation []InstitutionLocHolder
			InstitutionTypes    []institutionDomain.InstitutionTypes
		}

		data := PageData{provinces, "institution", "location", institutionsLocationHolder, institutionType}

		files := []string{
			app.Path + "content/tech/institution/location.html",
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

/***
this method will help to return every page under
institution tab depending on a variable called subTab
the page will be opened on what is specified in this variable
**/

/**
this method returns an open page on institutions
***/
func GetInstitutionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		var institutionsHolder []InstitutionHolder
		institutionTypes, err := institutionIO.GetInstitutionTypes()

		institutionsHolder = GetInstitutionHolder()

		type PageData struct {
			Tab                string
			SubTab             string
			InstitutionsHolder []InstitutionHolder
			InstitutionTypes   []institutionDomain.InstitutionTypes
		}

		data := PageData{"institution", "institution", institutionsHolder, institutionTypes}

		files := []string{
			app.Path + "content/tech/institution/institution.html",
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

func UpdateInstitutionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		institutionId := r.PostFormValue("Id")
		institutionType := r.PostFormValue("institutionType")
		name := r.PostFormValue("Name")

		if institutionType != "" || institutionId != "" || name != "" {
			institution := institutionDomain.Institution{institutionId, institutionType, name}
			fmt.Println("institution>>>>", institution)
			_, err := institutionIO.UpdateInstitution(institution, token)
			if err != nil {
				fmt.Println("error in Updating institution")
				app.ErrorLog.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/support/management/institution/institution", 301)
	}
}

func DeleteInstitutionAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		InstitutionAddressId := chi.URLParam(r, "InstitutionAddressId")
		AddressTypeId := chi.URLParam(r, "AddressTypeId")

		if InstitutionAddressId != "" {
			institutionAddressObject, err := institutionIO.ReadInstitutionAddress(InstitutionAddressId, AddressTypeId)
			if err != nil {
				fmt.Println("an error in DeleteInstitutionAddressHandler")
				app.ErrorLog.Println(err.Error())
			}
			if institutionAddressObject.InstitutionId != "" {
				_, err := institutionIO.DeleteInstitutionAddress(institutionAddressObject)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				}
			}
		}
		http.Redirect(w, r, "/support/management/institution/address", 301)
	}
}

func SaveInstitutionAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		addressType := r.PostFormValue("addressType")
		institutionId := r.PostFormValue("institutionAddress")
		postalCode := r.PostFormValue("postalCode")
		address := r.PostFormValue("address")

		if addressType != "" || institutionId != "" || postalCode != "" || address != "" {
			institutionAddress := institutionDomain.InstitutionAddress{institutionId, addressType, address, postalCode}
			_, err := institutionIO.CreateInstitutionAddress(institutionAddress)
			if err != nil {
				fmt.Println("error in creating institutionAddress")
				app.ErrorLog.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/support/management/institution/address", 301)
	}
}

func DeleteInstitutionCourseHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		institutionId := chi.URLParam(r, "institutionId")
		CourseId := chi.URLParam(r, "CourseId")

		if CourseId != "" || institutionId != "" {
			institutionCourserObject, err := institutionIO.ReadInstitutionCourse(institutionId, CourseId)
			if err != nil {
				fmt.Println("an error in DeleteInstitutionCourseHandler")
				app.ErrorLog.Println(err.Error())
			}
			if institutionCourserObject.InstitutionId != "" {
				_, err := institutionIO.DeleteInstitutionCourse(institutionCourserObject)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				}
			}
		}
		http.Redirect(w, r, "/support/management/institution/course", 301)
	}
}

func SaveInstitutionCourseHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		institutionId := r.PostFormValue("institutionCourseDrop")
		courseId := r.PostFormValue("courseId")

		if courseId != "" || institutionId != "" {
			institutionCourse := institutionDomain.InstitutionCourse{institutionId, courseId}
			_, err := institutionIO.CreateInstitutionCourse(institutionCourse)
			if err != nil {
				fmt.Println("error in creating institutionCourse")
				app.ErrorLog.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/support/management/institution/course", 301)
	}
}

type Tabs struct {
	Tab1 string
	Tab2 string
	Tab3 string
	Tab4 string
	Tab5 string
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
		http.Redirect(w, r, "/support/management/institution/location", 301)
	}
}

func SaveInstitutionLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		locationId := r.PostFormValue("town")
		institutionId := r.PostFormValue("institutionLocationDrop")
		longitude := r.PostFormValue("longitude")
		latitude := r.PostFormValue("latitude")

		if locationId != "" || institutionId != "" || longitude != "" || latitude != "" {
			institutionLocation := institutionDomain.InstitutionLocation{institutionId, locationId, longitude, latitude}

			fmt.Println(institutionLocation)
			_, err := institutionIO.CreateInstitutionLocation(institutionLocation)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/support/management/institution/location", 301)
	}
}

func DeleteInstitutionHandler(app *config.Env) http.HandlerFunc {
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
		http.Redirect(w, r, "/support/management/institution/institution", 301)
	}
}

func AddInstitutionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		institutionName := r.PostFormValue("institutionName")
		institutionType := r.PostFormValue("institutionType")
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		fmt.Println(institutionName, "  <<<<institutionName    institutionType>>>>>", institutionType)
		if institutionName != "" || institutionType != "" {
			institu := institutionDomain.Institution{"", institutionType, institutionName}
			_, err := institutionIO.CreateInstitution(institu)
			if err != nil {
				fmt.Println("error creating institution")
				app.ErrorLog.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/support/management/institution/institution", 301)
	}
}

func EditTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		institutionTypeId := r.PostFormValue("Id")
		institutionTypeName := r.PostFormValue("Name")
		institutionTypeDescription := r.PostFormValue("Description")

		_ = app.Session.Destroy(r.Context())

		if institutionTypeId != "" || institutionTypeName != "" || institutionTypeDescription != "" {

			institutionTypeObject := institutionDomain.InstitutionTypes{institutionTypeId, institutionTypeName, institutionTypeDescription}
			fmt.Print(institutionTypeObject)

			_, erro := institutionIO.UpdateInstitutionType(institutionTypeObject, token)
			if erro != nil {
				app.ErrorLog.Println(erro.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func DeleteTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resetKey := chi.URLParam(r, "formData")

		_ = app.Session.Destroy(r.Context())

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
		app.Session.Put(r.Context(), "tab", "tab1")
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

		_ = app.Session.Destroy(r.Context())

		institution := institutionDomain.InstitutionTypes{"", institutionName, institutionDescription}
		app.InfoLog.Println("Institution to save: ", institution)
		_, err := institutionIO.CreateInstitutionType(institution)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}
func GetTabs(tab string) Tabs {

	switch tab {
	case "tab1":
		return Tabs{"active show", "", "", "", ""}
	case "tab2":
		return Tabs{"", "active show", "", "", ""}
	case "tab3":
		return Tabs{"", "", "active show", "", ""}
	case "tab4":
		return Tabs{"", "", "", "active show", ""}
	case "tab5":
		return Tabs{"", "", "", "", "active show"}
	default:
		return Tabs{"active show", "", "", "", ""}
	}
}

//
//func InstitutionManagementHandler(app *config.Env) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		tab := app.Session.GetString(r.Context(), "tab")
//
//		activeTab := GetTabs(tab)
//		type InstitutionHolder struct {
//			Id              string
//			InstitutionName string
//			InstitutionType string
//		}
//
//
//		var institutionLocation []InstitutionLocHolder
//		var institutions []institutionDomain.Institution
//		var institutionsHolder []InstitutionHolder
//		var institutionsCourseHolder []InstitutionCourseHolder
//		var institutionsAddressHolder []InstitutionAddressHolder
//
//		institutionAddresses, err := institutionIO.GetAllInstitutionAddresses()
//		if err != nil {
//			app.ErrorLog.Println(err.Error())
//		} else if institutionAddresses != nil {
//			for _, institutionAddrres := range institutionAddresses {
//				myInstitution, err := institutionIO.GetInstitution(institutionAddrres.InstitutionId)
//				if err != nil {
//					fmt.Println("An error in InstitutionManagementHandler when reading myInstitution")
//					app.ErrorLog.Println(err.Error())
//				} else if myInstitution.Name != "" {
//					institutionAddress := InstitutionAddressHolder{institutionAddrres.AddressTypeId, institutionAddrres.InstitutionId, myInstitution.Name, institutionAddrres.Address, institutionAddrres.PostalCode}
//					institutionsAddressHolder = append(institutionsAddressHolder, institutionAddress)
//				}
//
//			}
//		}
//
//		allInstitutionCourse, err := institutionIO.GetAllInstitutionCourses()
//		if err != nil {
//			app.ErrorLog.Println(err.Error())
//		} else {
//			for _, institutionCourse := range allInstitutionCourse {
//				institution, err := institutionIO.GetInstitution(institutionCourse.InstitutionId)
//				if err != nil {
//					fmt.Println("error reading institution in InstitutionManagementHandler method")
//					app.ErrorLog.Println(err.Error())
//				}
//				couse, err := academics.GetCourse(institutionCourse.CourseId)
//				if err != nil {
//					fmt.Println("error reading institution in InstitutionManagementHandler method")
//					app.ErrorLog.Println(err.Error())
//				}
//				if institution.Name != "" || couse.CourseName != "" {
//					myInstitutionCours := InstitutionCourseHolder{institutionCourse.InstitutionId, institutionCourse.CourseId, institution.Name, couse.CourseName, couse.CourseDesc}
//					institutionsCourseHolder = append(institutionsCourseHolder, myInstitutionCours)
//				}
//			}
//		}
//
//		institutsLocation, err := institutionIO.ReadInstitutionLocations()
//		if err != nil {
//			app.ErrorLog.Println(err.Error())
//		} else {
//			for _, institutionLoca := range institutsLocation {
//				//fmt.Println("error in reading townNamw in InstitutionManagementHandler method", institutionLoca.LocationId)
//
//				institutionName, errr := institutionIO.GetInstitution(institutionLoca.InstitutionId)
//				if errr != nil {
//					fmt.Println("error in reading institutionName in InstitutionManagementHandler method")
//					app.ErrorLog.Println(errr.Error())
//				}
//				townNamw, err := location.GetLocation(institutionLoca.LocationId)
//				//fmt.Println("error in reading townNamw in InstitutionManagementHandler method", townNamw)
//				if err != nil {
//					fmt.Println("error in reading townNamw in InstitutionManagementHandler method")
//					app.ErrorLog.Println(err.Error())
//				}
//				institutionLocation = append(institutionLocation, InstitutionLocHolder{institutionLoca.InstitutionId, institutionName.Name, townNamw.Name, institutionLoca.Longitude, institutionLoca.Latitude})
//			}
//		}
//
//		institutionTypes, err := institutionIO.GetInstitutionTypes()
//		if err != nil {
//			app.ErrorLog.Println(err.Error())
//		} else {
//			institutions, err = institutionIO.GetInstitutions()
//			if err != nil {
//				app.ErrorLog.Println(err.Error())
//			} else {
//				for _, institution := range institutions {
//					institutionTypeName := getInstitutionTypeName(institution, institutionTypes)
//					institutionsHolder = append(institutionsHolder, InstitutionHolder{institution.Id, institution.Name, institutionTypeName})
//				}
//			}
//		}
//		provinces, _ := locationHelper.GetProvinces(app)
//		fmt.Println("provinces>>>", provinces)
//		courses, err := academics.GetAllCourses()
//		if err != nil {
//			fmt.Println("error in InstitutionManagementHandler when reading courses")
//			app.ErrorLog.Println(err.Error())
//		}
//		addressTypes, err := address.GetAddressTypes()
//		if err != nil {
//			fmt.Println("error in InstitutionManagementHandler when reading addressTypes")
//			app.ErrorLog.Println(err.Error())
//		}
//
//		type PageData struct {
//			InstitutionTypes    []institutionDomain.InstitutionTypes
//			InstitutionsHolder  []InstitutionHolder
//			Provinces           []locationDomain.Location
//			InstitutionLocation []InstitutionLocHolder
//			MyActiveTab         Tabs
//			Courses             []academicsDomain.Course
//			InstitutionCourse   []InstitutionCourseHolder
//			AddressTypes        []address.AddressType
//			InstitutionAddress  []InstitutionAddressHolder
//		}
//		data := PageData{institutionTypes, institutionsHolder, provinces, institutionLocation, activeTab, courses, institutionsCourseHolder, addressTypes, institutionsAddressHolder}
//
//		files := []string{
//			app.Path + "content/tech/tech_admin_institution.html",
//			app.Path + "content/tech/template/sidebar.template.html",
//			app.Path + "base/template/form/location-form.template.html",
//			app.Path + "base/template/form/institution-form.template.html",
//			app.Path + "base/template/footer.template.html",
//		}
//		ts, err := template.ParseFiles(files...)
//		if err != nil {
//			app.ErrorLog.Println(err.Error())
//			return
//		}
//		err = ts.Execute(w, data)
//		if err != nil {
//			app.ErrorLog.Println(err.Error())
//		}
//	}
//}

func getInstitutionTypeName(institution institutionDomain.Institution, institutionTypes []institutionDomain.InstitutionTypes) string {
	var institutionTypeName string
	for _, institutionType := range institutionTypes {
		if institution.InstitutionTypeId == institutionType.Id {
			institutionTypeName = institutionType.Name
		}
	}
	return institutionTypeName
}
