package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	domain3 "obas/domain/application"
	domain2 "obas/domain/documents"
	domain "obas/domain/institutions"
	domain4 "obas/domain/location"
	domain5 "obas/domain/users"
	"obas/io/address"
	"obas/io/applications"
	"obas/io/documents"
	institutions2 "obas/io/institutions"
	"obas/io/location"
)

func TypesManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", TypeHomeHandler(app))
	r.Post("/institutionType/create", CreateInstitutionTypeHandler(app))
	r.Get("/institutiontype/delete/{institutionId}", DeleteInstitutionTypeHandler(app))

	r.Post("/documentType/create", CreateDocumentTypeHandler(app))
	r.Post("/documentType/update", UpdateDocumentTypeHandler(app))
	r.Get("/documentType/delete/{documentType}", DeleteDocumentTypeHandler(app))

	r.Post("/addressType/update", UpdateAddressTypeHandler(app))
	r.Post("/addressType/create", CreateAddressTypeHandler(app))
	r.Get("/addressType/delete/{addressTypeId}", DeleteAddressTypeHandler(app))

	r.Post("/applicationType/create", CreateApplicationTypeHandler(app))
	r.Post("/applicationType/update", UpdateApplicationTypeHandler(app))
	r.Get("/applicationType/delete/{applicationTypeId}", DeleteApplicationTypeHandler(app))

	r.Post("/applicantType/create", CreateApplicantTypeHandler(app))
	r.Post("/applicantType/update", UpdateApplicantTypeHandler(app))
	r.Get("/applicantType/delete/{applicantTypeId}", DeleteApplicantTypeHandler(app))

	r.Post("/locationType/create", CreateLocationTypeHandler(app))
	r.Get("/locationType/delete/{locationTypeId}", DeleteLocationTypeHandler(app))
	r.Post("/locationType/update", UpdateLocationTypeHandler(app))

	return r
}

func UpdateLocationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		locationType := r.PostFormValue("locationName")
		Code := r.PostFormValue("locationCode")
		id := r.PostFormValue("locationId")

		if locationType != "" || id != "" || Code != "" {

			locationType := domain4.LocationType{id, locationType, Code}
			result, err := location.UpdateLocationType(locationType)
			if err != nil {
				fmt.Println("error when updating Applicant Type ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			if result != true {
				fmt.Println("error when creating Applicant Type. ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "You have successfully updating a Applicant Type."
			messageType = "info"
		}
		app.Session.Put(r.Context(), "tab", "tab6")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func DeleteLocationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		locationTypeId := chi.URLParam(r, "locationTypeId")

		if locationTypeId != "" {

			locationType, err := location.GetLocationType(locationTypeId)
			if err != nil {
				fmt.Println("error when reading Location Type ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			} else {
				result, err := location.DeleteLocationType(locationType)
				if err != nil {
					fmt.Println("error when deleting Location Type ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				if result != true {
					fmt.Println("error when Deleting Location Type. ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				message = "You have successfully deleting an Location Type."
				messageType = "info"
			}
		}
		app.Session.Put(r.Context(), "tab", "tab6")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func CreateLocationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		locationType := r.PostFormValue("locationType")
		code := r.PostFormValue("code")

		if locationType != "" || code != "" {

			locationTypeObj := domain4.LocationType{"", locationType, code}
			result, err := location.CreateLocationType(locationTypeObj)
			if err != nil {
				fmt.Println("error when updating Location Type ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "An error has occurred please try again later."
			messageType = "warning"

			if result.LocationTypeId != "" {
				message = "You have successfully creating a Location Type."
				messageType = "info"
			}

		}
		app.Session.Put(r.Context(), "tab", "tab6")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func DeleteApplicantTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		applicantType := chi.URLParam(r, "applicantTypeId")

		if applicantType != "" {

			applicantType, err := applications.GetApplicantType(applicantType)
			if err != nil {
				fmt.Println("error when reading Applicant Type ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			} else {
				result, err := applications.DeleteApplicantType(applicantType)
				if err != nil {
					fmt.Println("error when deleting Applicant Type ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				if result != true {
					fmt.Println("error when Deleting Applicant Type. ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				message = "You have successfully deleting an Applicant Type."
				messageType = "info"
			}
		}
		app.Session.Put(r.Context(), "tab", "tab5")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func UpdateApplicantTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		applicantType := r.PostFormValue("applicantName")
		Description := r.PostFormValue("applicantDescription")
		id := r.PostFormValue("applicantId")

		if applicantType != "" || id != "" || Description != "" {

			applicationType := domain3.ApplicantType{id, applicantType, Description}
			result, err := applications.UpdateApplicantType(applicationType)
			if err != nil {
				fmt.Println("error when updating Applicant Type ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			if result.Id == "" {
				fmt.Println("error when creating Applicant Type. ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "You have successfully updating a Applicant Type."
			messageType = "info"
		}
		app.Session.Put(r.Context(), "tab", "tab5")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func CreateApplicantTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		applicantType := r.PostFormValue("applicantType")
		Description := r.PostFormValue("Description")

		if applicantType != "" || Description != "" {

			applicantObj := domain3.ApplicantType{"", applicantType, Description}
			result, err := applications.CreateApplicantType(applicantObj)
			if err != nil {
				fmt.Println("error when updating applicant type ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "An error has occurred please try again later."
			messageType = "warning"

			if result.Id != "" {
				message = "You have successfully creating a Applicant Type."
				messageType = "info"
			}

		}
		app.Session.Put(r.Context(), "tab", "tab4")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func DeleteApplicationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		applicationType := chi.URLParam(r, "applicationTypeId")

		if applicationType != "" {

			applicationType, err := applications.GetApplicationType(applicationType)
			if err != nil {
				fmt.Println("error when reading applicationType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			} else {
				result, err := applications.DeleteApplicationType(applicationType)
				if err != nil {
					fmt.Println("error when deleting applicationType ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				if result != true {
					fmt.Println("error when Deleting applicationType. ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				message = "You have successfully deleting an Application Type."
				messageType = "info"
			}
		}
		app.Session.Put(r.Context(), "tab", "tab4")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func UpdateApplicationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		applicationType := r.PostFormValue("applicationName")
		Description := r.PostFormValue("applicationDescription")
		id := r.PostFormValue("applicationId")

		if applicationType != "" || id != "" || Description != "" {

			applicationType := domain3.ApplicationType{id, applicationType, Description}
			result, err := applications.UpdateApplicationType(applicationType)
			if err != nil {
				fmt.Println("error when updating institutionType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			if result != true {
				fmt.Println("error when creating institutionType. ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "You have successfully updating a Application Type."
			messageType = "info"
		}
		app.Session.Put(r.Context(), "tab", "tab4")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func CreateApplicationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		applicantType := r.PostFormValue("Name")
		Description := r.PostFormValue("Description")

		if applicantType != "" || Description != "" {

			applicantObj := domain3.ApplicationType{"", applicantType, Description}
			result, err := applications.CreateApplicationType(applicantObj)
			if err != nil {
				fmt.Println("error when updating institutionType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}

			if result != true {
				fmt.Println("error when creating Application type. ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "You have successfully creating a Application Type."
			messageType = "info"

		}
		app.Session.Put(r.Context(), "tab", "tab4")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func DeleteAddressTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		addressTypeId := chi.URLParam(r, "addressTypeId")

		if addressTypeId != "" {

			addressType, err := address.GetAddressType(addressTypeId)
			if err != nil {
				fmt.Println("error when reading documentType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			} else {
				result, err := address.DeleteAddressType(addressType)
				if err != nil {
					fmt.Println("error when reading institutionType ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				if result != true {
					fmt.Println("error when creating institutionType. ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				message = "You have successfully deleting an Institution Type."
				messageType = "info"
			}
		}
		app.Session.Put(r.Context(), "tab", "tab3")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func CreateAddressTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		addressType := r.PostFormValue("Name")

		if addressType != "" {

			AddressType := address.AddressType{"", addressType}
			result, err := address.CreateAddressType(AddressType)
			if err != nil {
				fmt.Println("error when updating institutionType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			if result != true {
				fmt.Println("error when creating institutionType. ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "You have successfully creating a Address Type."
			messageType = "info"
		}
		app.Session.Put(r.Context(), "tab", "tab3")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func UpdateAddressTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		addressType := r.PostFormValue("addressName")
		id := r.PostFormValue("addressId")

		if addressType != "" || id != "" {

			AddressType := address.AddressType{id, addressType}
			result, err := address.UpdateAddressType(AddressType)
			if err != nil {
				fmt.Println("error when updating institutionType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			if result != true {
				fmt.Println("error when creating institutionType. ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "You have successfully updating a Address Type."
			messageType = "info"
		}
		app.Session.Put(r.Context(), "tab", "tab3")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func DeleteDocumentTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		documentTypeId := chi.URLParam(r, "documentType")

		if documentTypeId != "" {

			documentType, err := documents.GetDocumentType(documentTypeId)
			if err != nil {
				fmt.Println("error when reading institutionType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			} else {
				result, err := documents.DeleteDocumentType(documentType)
				if err != nil {
					fmt.Println("error when reading institutionType ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				if result != true {
					fmt.Println("error when creating institutionType. ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				message = "You have successfully created an Institution Type."
				messageType = "info"
			}
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func UpdateDocumentTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		documentType := r.PostFormValue("documetName")
		id := r.PostFormValue("documentTypeId")

		if documentType != "" || id != "" {

			institutionType := domain2.DocumentType{id, documentType}
			result, err := documents.UpdateDocumentType(institutionType)
			if err != nil {
				fmt.Println("error when updating institutionType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			if result != true {
				fmt.Println("error when creating institutionType. ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "You have successfully updating a Document Type."
			messageType = "info"
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func CreateDocumentTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		documentType := r.PostFormValue("documentType")

		if documentType != "" {
			institutionType := domain2.DocumentType{"", documentType}

			result, err := documents.CreateDocumentType(institutionType, token)
			if err != nil {
				fmt.Println("error when creating institutionType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			if result != true {
				fmt.Println("error when creating institutionType. ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "You have successfully created a Document Type."
			messageType = "info"
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func DeleteInstitutionTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		institutionTypeId := chi.URLParam(r, "institutionId")
		_ = app.Session.Destroy(r.Context())

		if institutionTypeId != "" {
			institutionType, err := institutions2.GetInstitutionType(institutionTypeId)
			if err != nil {
				fmt.Println("error when reading institutionType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			} else {
				result, err := institutions2.DeleteInstitutionType(institutionType)
				if err != nil {
					fmt.Println("error when reading institutionType ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				if result != true {
					fmt.Println("error when creating institutionType. ")
					message = "An error has occurred please try again later."
					messageType = "warning"
				}
				message = "You have successfully created an Institution Type."
				messageType = "info"
			}

		}
		app.Session.Put(r.Context(), "tab", "tab1")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func CreateInstitutionTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message string
		var messageType string

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		institutionType := r.PostFormValue("institutionType")
		Description := r.PostFormValue("Description")

		if institutionType != "" || Description != "" {
			institutionType := domain.InstitutionTypes{"", institutionType, Description}

			result, err := institutions2.CreateInstitutionType(institutionType)
			if err != nil {
				fmt.Println("error when creating institutionType ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			if result.Id != "" {
				fmt.Println("error when creating institutionType. ")
				message = "An error has occurred please try again later."
				messageType = "warning"
			}
			message = "You have successfully created an Institution Type."
			messageType = "info"
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		app.Session.Put(r.Context(), "message", message)
		app.Session.Put(r.Context(), "messageType", messageType)
		http.Redirect(w, r, "/support/types/", 301)
	}
}

func TypeHomeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		tab := app.Session.GetString(r.Context(), "tab")
		message := app.Session.GetString(r.Context(), "message")
		messageType := app.Session.GetString(r.Context(), "messageType")

		if userId == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		institutionType, err := institutions2.GetInstitutionTypes()
		if err != nil {
			fmt.Println("error reading institutionTypes")
		}
		documentType, err := documents.GetDocumentTypes()
		if err != nil {
			fmt.Println("error reading documentType")
		}
		addressType, err := address.GetAddressTypes()
		if err != nil {
			fmt.Println("error reading AddressType")
		}
		applicationType, err := applications.GetApplicationTypes()
		if err != nil {
			fmt.Println("error reading applicationType")
		}
		applicantType, err := applications.GetApplicantTypes()
		if err != nil {
			fmt.Println("error reading applicantType")
		}
		locationType, err := location.GetLocationTypes()
		if err != nil {
			fmt.Println("error reading applicantType")
		}

		tabPosition := getTab(tab)

		type PageData struct {
			InstitutionType []domain.InstitutionTypes
			DocumentType    []domain2.DocumentType
			AddressType     []address.AddressType
			ApplicationType []domain3.ApplicationType
			ApplicantType   []domain3.ApplicantType
			LocationType    []domain4.LocationType
			MyActiveTab     TypeTab
			Message         string
			MessageType     string
			Tab             string
			SubTab          string
			ProfileUser     domain5.User
		}
		data := PageData{
			InstitutionType: institutionType,
			DocumentType:    documentType,
			AddressType:     addressType,
			ApplicationType: applicationType,
			ApplicantType:   applicantType,
			LocationType:    locationType,
			MyActiveTab:     tabPosition,
			Message:         message,
			MessageType:     messageType,
			Tab:             "types",
			SubTab:          "",
			ProfileUser:     getUser(userId),
		}

		files := []string{
			app.Path + "content/tech/tech_admin_types.html",
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

type TypeTab struct {
	Tab1 string
	Tab2 string
	Tab3 string
	Tab4 string
	Tab5 string
	Tab6 string
}

func getTab(tab string) TypeTab {
	switch tab {
	case "tab1":
		return TypeTab{"active show", "", "", "", "", ""}
	case "tab2":
		return TypeTab{"", "active show", "", "", "", ""}
	case "tab3":
		return TypeTab{"", "", "active show", "", "", ""}
	case "tab4":
		return TypeTab{"", "", "", "active show", "", ""}
	case "tab5":
		return TypeTab{"", "", "", "", "active show", ""}
	case "tab6":
		return TypeTab{"", "", "", "", "", "active show"}
	default:
		return TypeTab{"active show", "", "", "", "", ""}
	}
}
