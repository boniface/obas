package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"html/template"
	"net/http"
	"obas/config"
	"obas/io/demographics"
)

func DemographyManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", DemogrphyHandler(app))
	r.Post("/create/title", CreateTileHandler(app))
	r.Post("/create/gender", CreateGenderHandler(app))
	r.Post("/create/role", RoleCreateHandler(app))
	r.Post("/update/role", RoleUpdateHandler(app))
	r.Post("/update/title", UpdateTileHandler(app))
	r.Post("/update/gender", UpdateGenderHandler(app))
	r.Get("/title/delete/{titleId}", DeleteTitleHandler(app))
	r.Get("/title/deleteJ/{titleId}", DeleteJTitleHandler(app))
	r.Get("/gender/delete/{genderId}", DeleteGenderHandler(app))
	r.Get("/role/delete/{roleId}", DeleteRoleHandler(app))

	return r
}

func DeleteJTitleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		titleId := chi.URLParam(r, "titleId")
		var message = ""
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		_ = app.Session.Destroy(r.Context())
		if titleId != "" {
			title, err := demographics.GetTitle(titleId)
			if err != nil {
				fmt.Println("error reading title")
				app.ErrorLog.Println(err.Error())
			}
			if title.TitleId != "" {
				fmt.Println(title, "<<<<<title")
				result, err := demographics.DeleteTitle(title)
				if err != nil {
					fmt.Println("error updating loccation type")
					app.ErrorLog.Println(err.Error())
				}
				if result == true {
					message = "A new Title was created successfully (" + title.TitleName + ")"
				} else {
					message = " An error has occurred please try again"
				}
			}
		}

		app.Session.Put(r.Context(), "SupportMessage", message)
		app.Session.Put(r.Context(), "tab", "tab1")
		app.Session.Put(r.Context(), "userId", email)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/demography", 301)
	}
}

func DeleteRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleId := chi.URLParam(r, "roleId")
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message = ""
		_ = app.Session.Destroy(r.Context())
		roleeObject, err := demographics.GetRole(roleId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if roleeObject.Id != "" {
			result, err := demographics.DeleteRole(roleeObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
			if result == true {
				message = "A new Gender created successfully (" + roleeObject.RoleName + ")"
			} else {
				message = " An error has occurred please try again"
			}
			fmt.Println(message)
			app.Session.Put(r.Context(), "userId", userId)
			app.Session.Put(r.Context(), "token", token)
			render.JSON(w, r, result)
		}

	}
}

func RoleUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message = ""
		_ = app.Session.Destroy(r.Context())
		r.ParseForm()
		role := r.PostFormValue("role")

		if role != "" {
			//newRole:=domain.Role{"",role}
			newRole := demographics.Role{"", role}
			result, err := demographics.CreateRole(newRole)
			if err != nil {
				fmt.Println("error creating Role")
			}
			if result == true {
				message = "A new Gender created successfully (" + newRole.RoleName + ")"
			} else {
				message = " An error has occurred please try again"
			}
		}

		app.Session.Put(r.Context(), "SupportMessage", message)
		app.Session.Put(r.Context(), "tab", "tab3")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/demography", 301)
	}
}
func RoleCreateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		var message = ""
		_ = app.Session.Destroy(r.Context())
		r.ParseForm()
		role := r.PostFormValue("role")

		if role != "" {
			//newRole:=domain.Role{"",role}
			newRole := demographics.Role{"", role}
			result, err := demographics.CreateRole(newRole)
			if err != nil {
				fmt.Println("error creating Role")
			}
			if result == true {
				message = "A new Gender created successfully (" + newRole.RoleName + ")"
			} else {
				message = " An error has occurred please try again"
			}
		}

		app.Session.Put(r.Context(), "SupportMessage", message)
		app.Session.Put(r.Context(), "tab", "tab3")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/management/user", 301)
	}
}
func CreateGenderHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		_ = app.Session.Destroy(r.Context())

		r.ParseForm()

		gender := r.PostFormValue("genderName")

		var message = ""

		fmt.Println(gender, "<<<<<tile")
		if gender != "" {
			newgender := demographics.Gender{"", gender}
			fmt.Println(newgender, "<<<<<newgender")
			result, err := demographics.CreateGender(newgender)
			if err != nil {
				fmt.Println("error updating loccation type")
				app.ErrorLog.Println(err.Error())
			}
			if result == true {
				message = "A new Gender created successfully (" + newgender.GenderName + ")"
			} else {
				message = " An error has occurred please try again"
			}
		}
		app.Session.Put(r.Context(), "SupportMessage", message)
		app.Session.Put(r.Context(), "tab", "tab2")
		app.Session.Put(r.Context(), "userId", email)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/demography", 301)
	}
}

func UpdateGenderHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		_ = app.Session.Destroy(r.Context())

		r.ParseForm()

		gender := r.PostFormValue("genderName")
		id := r.PostFormValue("Id")
		var message = ""

		fmt.Println(gender, "<<<<<tile||id", id)
		if gender != "" || id != "" {
			newgender := demographics.Gender{id, gender}
			fmt.Println(newgender, "<<<<<newgender")
			result, err := demographics.UpdateGender(newgender, token)
			if err != nil {
				fmt.Println("error updating loccation type")
				app.ErrorLog.Println(err.Error())
			}
			if result == true {
				message = "A new Gender update was successful (" + newgender.GenderName + ")"
			} else {
				message = " An error has occurred please try again"
			}
		}
		app.Session.Put(r.Context(), "SupportMessage", message)
		app.Session.Put(r.Context(), "tab", "tab2")
		app.Session.Put(r.Context(), "userId", email)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/demography", 301)
	}
}

func DeleteGenderHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		genderId := chi.URLParam(r, "genderId")
		var message = ""
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		_ = app.Session.Destroy(r.Context())
		if genderId != "" {
			gender, err := demographics.GetGender(genderId)
			if err != nil {
				fmt.Println("error reading gender")
				app.ErrorLog.Println(err.Error())
			}
			if gender.GenderId != "" {
				fmt.Println(gender, "<<<<<title")
				result, err := demographics.DeleteGender(gender)
				if err != nil {
					fmt.Println("error updating loccation type")
					app.ErrorLog.Println(err.Error())
				}
				if result == true {
					message = "A new Title was created successfully (" + gender.GenderName + ")"
				} else {
					message = " An error has occurred please try again"
				}
			}
		}

		app.Session.Put(r.Context(), "SupportMessage", message)
		app.Session.Put(r.Context(), "tab", "tab2")
		app.Session.Put(r.Context(), "userId", email)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/demography", 301)
	}
}

func DeleteTitleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		titleId := chi.URLParam(r, "titleId")
		var message = ""
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		_ = app.Session.Destroy(r.Context())
		if titleId != "" {
			title, err := demographics.GetTitle(titleId)
			if err != nil {
				fmt.Println("error reading title")
				app.ErrorLog.Println(err.Error())
			}
			if title.TitleId != "" {
				fmt.Println(title, "<<<<<title")
				result, err := demographics.DeleteTitle(title)
				if err != nil {
					fmt.Println("error updating loccation type")
					app.ErrorLog.Println(err.Error())
				}
				if result == true {
					message = "A new Title was created successfully (" + title.TitleName + ")"
				} else {
					message = " An error has occurred please try again"
				}
			}
		}

		app.Session.Put(r.Context(), "SupportMessage", message)
		app.Session.Put(r.Context(), "tab", "tab1")
		app.Session.Put(r.Context(), "userId", email)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/demography", 301)
	}
}
func UpdateTileHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		_ = app.Session.Destroy(r.Context())

		r.ParseForm()

		title := r.PostFormValue("titleName")
		id := r.PostFormValue("Id")
		var message = ""

		fmt.Println(title, "<<<<<tile||")
		if title != "" || id != "" {
			newtitle := demographics.Title{id, title}
			fmt.Println(newtitle, "<<<<<newtitle")
			result, err := demographics.UpdateTitle(newtitle)
			if err != nil {
				fmt.Println("error updating loccation type")
				app.ErrorLog.Println(err.Error())
			}
			if result == true {
				message = "A new Title update was successful (" + newtitle.TitleName + ")"
			} else {
				message = " An error has occurred please try again"
			}
		}
		app.Session.Put(r.Context(), "SupportMessage", message)
		app.Session.Put(r.Context(), "tab", "tab1")
		app.Session.Put(r.Context(), "userId", email)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/demography", 301)
	}
}

func CreateTileHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}

		_ = app.Session.Destroy(r.Context())

		r.ParseForm()

		title := r.PostFormValue("tile")

		var message = ""

		fmt.Println(title, "<<<<<tile||")
		if title != "" {
			newtitle := demographics.Title{"", title}
			fmt.Println(newtitle, "<<<<<newtitle")
			result, err := demographics.CreateTitle(newtitle)
			if err != nil {
				fmt.Println("error updating loccation type")
				app.ErrorLog.Println(err.Error())
			}
			if result == true {
				message = "A new Title Creation was successful (" + newtitle.TitleName + ")"
			} else {
				message = " An error has occurred please try again"
			}
		}
		app.Session.Put(r.Context(), "SupportMessage", message)
		app.Session.Put(r.Context(), "tab", "tab1")
		app.Session.Put(r.Context(), "userId", email)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/demography", 301)
	}
}

type ActiveTab struct {
	Tab1 string
	Tab2 string
	Tab3 string
}

func DemographyTabs(tab string) ActiveTab {
	switch tab {
	case "tab1":
		return ActiveTab{"active show", "", ""}
	case "tab2":
		return ActiveTab{"", "active show", ""}
	case "tab3":
		return ActiveTab{"", "", "active show"}
	default:
		return ActiveTab{"active show", "", ""}
	}
}

func DemogrphyHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		tab := app.Session.GetString(r.Context(), "tab")
		message := app.Session.GetString(r.Context(), "SupportMessage")
		activeTab := DemographyTabs(tab)

		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		Titles, err := demographics.GetTitles()
		if err != nil {
			fmt.Println("error reading Titles")
		}
		Genders, err := demographics.GetGenders()
		if err != nil {
			fmt.Println("error reading Genders")
		}
		Roles, err := demographics.GetRoles()
		if err != nil {
			fmt.Println("error reading Roles")
		}

		type PageData struct {
			Title   []demographics.Title
			Gender  []demographics.Gender
			Roles   []demographics.Role
			Tabs    ActiveTab
			Message string
			Tab     string
			SubTab  string
		}

		Data := PageData{
			Titles,
			Genders,
			Roles,
			activeTab,
			message,
			"demography",
			"",
		}

		files := []string{
			app.Path + "content/tech/tech_admin_demography.html",
			app.Path + "content/tech/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		app.Session.Remove(r.Context(), "tab")
		app.Session.Remove(r.Context(), "SupportMessage")
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, Data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}

}
