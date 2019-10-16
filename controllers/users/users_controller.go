package controllers

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	addressIO "obas/io/address"
	usersIO "obas/io/users"
	"strings"
	"time"
)

const (
	layoutOBAS = "2006-01-02"
)

type AddressPlaceHolder struct {
	AddressName string
	Address     string
	PostalCode  string
}

func Users(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", UsersHandler(app))
	r.Get("/admin", AdminHandler(app))
	r.Get("/student", StudentHandler(app))

	r.Get("/processingStatus", ProcessingStatusTypeHandler(app))
	r.Get("/student/application", StudentApplicationStatusHandler(app))
	r.Get("/studentContact", StudentContactsHandler(app))
	r.Get("/studentDemographics", StudentDemographicsHandler(app))
	r.Get("/student/documents", StudentDocumentsHandler(app))
	r.Get("/studentResults", StudentResultsHandler(app))

	r.Get("/student/profile/personal", StudentProfilePersonalHandler(app))
	r.Get("/student/profile/address", StudentProfileAddressHandler(app))
	r.Get("/student/profile/guardian", StudentProfileGuardianHandler(app))
	r.Get("/student/profile/registration", StudentProfileRegistrationHandler(app))
	r.Get("/student/profile/courses", StudentProfileCourseHandler(app))
	r.Get("/student/profile/subjects", studentProfileSubjectHandler(app))
	r.Get("/student/profile/districts", studentProfileDistrictHandler(app))

	r.Post("/student/profile/personal/update", UpdateStudentProfilePersonalHandler(app))
	r.Post("/student/profile/address/addresstype", StudentProfileAddressTypeHandler(app))
	r.Post("/student/profile/address/update", StudentProfileAddressUpdateHandler(app))

	return r
}

func StudentProfileRegistrationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			//http.Redirect(w, r, "/login", 301)
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/registration.html", //***************
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

func StudentProfileCourseHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			//http.Redirect(w, r, "/login", 301)
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/courses.html", //***************
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

func studentProfileDistrictHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			//http.Redirect(w, r, "/login", 301)
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/distric_and_municipality.html", //***************
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

func studentProfileSubjectHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			//http.Redirect(w, r, "/login", 301)
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/subject.html ", //***************
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

func StudentProfileGuardianHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			//http.Redirect(w, r, "/login", 301)
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/guardian.html",
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

func StudentProfileAddressUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		addressTypeId := r.PostFormValue("addressTypeId")
		address := r.PostFormValue("address")
		postalCode := r.PostFormValue("postalCode")
		userAddress := usersIO.UserAddress{email, addressTypeId, address, postalCode}
		fmt.Println("UserAddress to update: ", userAddress)
		updated, err := usersIO.UpdateUserAddress(userAddress, token)
		fmt.Println("result of update: ", updated)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/address", 301)
	}
}

func StudentProfileAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			//http.Redirect(w, r, "/login", 301)
		}
		addressTypes, err := addressIO.GetAddressTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error(), addressTypes)
		}

		addresses := []AddressPlaceHolder{}

		for _, addressType := range addressTypes {
			userAddress, err := usersIO.GetUserAddress(email, addressType.AddressTypeID)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				addresses = append(addresses, AddressPlaceHolder{addressType.AddressName, userAddress.Address, userAddress.PostalCode})
			}
		}

		type PageData struct {
			Student       usersIO.User
			AddressTypes  []addressIO.AddressType
			Addresses     []AddressPlaceHolder
			Address       usersIO.UserAddress
			AddressTypeID string
			AddressName   string
		}

		data := PageData{user, addressTypes, addresses, usersIO.UserAddress{}, "", ""}
		files := []string{
			app.Path + "content/student/profile/address.html",
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

func StudentProfileAddressTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			//http.Redirect(w, r, "/login", 301)
		}
		r.ParseForm()
		addressTypeId := r.PostFormValue("addresstypes")
		userAddress, err := usersIO.GetUserAddress(email, addressTypeId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		fmt.Println(userAddress)

		addressTypes, err := addressIO.GetAddressTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error(), addressTypes)
		}

		addresses := []AddressPlaceHolder{}
		var addressName string

		for _, addressType := range addressTypes {
			if addressTypeId == addressType.AddressTypeID {
				addressName = addressType.AddressName
			}
			userAddress, err := usersIO.GetUserAddress(email, addressType.AddressTypeID)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				addresses = append(addresses, AddressPlaceHolder{addressType.AddressName, userAddress.Address, userAddress.PostalCode})
			}
		}

		type PageData struct {
			Student       usersIO.User
			AddressTypes  []addressIO.AddressType
			Addresses     []AddressPlaceHolder
			Address       usersIO.UserAddress
			AddressTypeID string
			AddressName   string
		}

		data := PageData{user, addressTypes, addresses, userAddress, addressTypeId, addressName}
		files := []string{
			app.Path + "content/student/profile/address.html",
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

func StudentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			//http.Redirect(w, r, "/login", 301)
		}
		type PageData struct {
			Student usersIO.User
		}
		data := PageData{user}
		files := []string{
			app.Path + "content/student/student_dashboard.page.html",
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

func StudentProfilePersonalHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			//http.Redirect(w, r, "/login", 301)
		}
		dobString := strings.Split(user.DateOfBirth.String(), " ")[0] // split date and get in format: yyy-mm-dd

		type PageData struct {
			Student     usersIO.User
			DateOfBirth string
		}

		data := PageData{user, dobString}
		files := []string{
			app.Path + "content/student/profile/personal.html",
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

func UpdateStudentProfilePersonalHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		idNumber := r.PostFormValue("id_number")
		firstName := r.PostFormValue("first_name")
		lastName := r.PostFormValue("last_name")
		dateOfBirthStr := r.PostFormValue("dateOfBirth")
		dateOfBirth, _ := time.Parse(layoutOBAS, dateOfBirthStr)
		user := usersIO.User{email, idNumber, firstName, "", lastName, dateOfBirth}
		fmt.Println("User to update: ", user)
		updated, err := usersIO.UpdateUser(user, token)
		fmt.Println("result of update: ", updated)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Update response is ", updated)
		http.Redirect(w, r, "/users/student/profile", 301)
	}
}

func UsersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			//courses []io.Users
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "base/register/register.page.html",
			/**app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.page.html",*/
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

func AdminHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allAdmin, err := io.GetAdmins()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//courses []io.Admin
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
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

func ProcessingStatusTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allProcess, err := io.GetProcessingStatusTypes()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//subjects []io.ProcessingStatusType
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
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

func StudentApplicationStatusHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "content/student/Student_Application.html",
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

func StudentContactsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allStudentContacts, err := io.GetStudentContacts()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//subjects []io.StudentContacts
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
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
func StudentDemographicsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allStudentDemographics, err := io.GetStudentDemographics()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//subjects []io.StudentDemographics
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
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
func StudentDocumentsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{

			app.Path + "content/student/Student_Documents.html",
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

func StudentResultsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allStudentResults, err := io.GetStudentResults()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//subjects []io.StudentResults
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
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
