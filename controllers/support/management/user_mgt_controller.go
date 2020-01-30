package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	domain "obas/domain/application"
	userDomain "obas/domain/users"
	"obas/io/applications"
	"obas/io/demographics"
	userIO "obas/io/users"
	"strings"
)

func UserManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", UserManagementHandler(app))
	r.Get("/role/delete/{roleId}", DeleteRoleManagementHandler(app))
	r.Get("/applicantType/delete/{Id}", DeleteApplicantTypeManagementHandler(app))
	r.Post("/role/create", RoleCreateManagementHandler(app))
	r.Post("/applicantType/create", ApplicantionCreateManagementHandler(app))
	r.Post("/role/update", RoleUpdateManagementHandler(app))
	r.Post("/applicantType/update", ApplicantUpdateManagementHandler(app))

	return r
}

func ApplicantUpdateManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		_ = app.Session.Destroy(r.Context())
		r.ParseForm()
		id := r.PostFormValue("id")
		name := r.PostFormValue("Name")
		Description := r.PostFormValue("Description")

		if id != "" || Description != "" || name != "" {
			newApplicantion := domain.ApplicantType{id, name, Description}
			_, err := applications.UpdateApplicantType(newApplicantion)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				fmt.Println("fail to update applicant type")
			}
		}
		app.Session.Put(r.Context(), "tab", "tab3")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/management/user", 301)
	}
}

func DeleteApplicantTypeManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleId := chi.URLParam(r, "Id")
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		_ = app.Session.Destroy(r.Context())
		applicantType, err := applications.GetApplicantType(roleId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if applicantType.Id != "" {
			_, err := applications.DeleteApplicantType(applicantType)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab3")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/management/user", 301)
	}
}

func ApplicantionCreateManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		_ = app.Session.Destroy(r.Context())
		r.ParseForm()
		applicationType := r.PostFormValue("applicationType")
		Description := r.PostFormValue("Description")

		if applicationType != "" || Description != "" {
			newApplicantion := domain.ApplicantType{"", applicationType, Description}
			_, err := applications.CreateApplicantType(newApplicantion)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab3")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/management/user", 301)
	}
}

func DeleteRoleManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleId := chi.URLParam(r, "roleId")
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		_ = app.Session.Destroy(r.Context())
		roleeObject, err := demographics.GetRole(roleId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if roleeObject.Id != "" {
			_, err := demographics.DeleteRole(roleeObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/management/user", 301)
	}
}

func RoleCreateManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		_ = app.Session.Destroy(r.Context())
		r.ParseForm()
		role := r.PostFormValue("role")

		if role != "" {
			//newRole:=domain.Role{"",role}
			newRole := demographics.Role{"", role}
			_, err := demographics.CreateRole(newRole)
			if err != nil {
				fmt.Println("error creating Role")
			}
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/management/user", 301)
	}
}

func RoleUpdateManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		_ = app.Session.Destroy(r.Context())
		r.ParseForm()
		id := strings.TrimSpace(r.PostFormValue("userId"))
		role := r.PostFormValue("roleId")
		fmt.Println(">>>>", id, "<<<<<<actualUser||actualRole>>>>>", role, "user ", userId)
		if id != "" || role != "" || token != "" {
			userRole := userDomain.UserRole{id, role}
			_, err := userIO.UpdateUserRole(userRole, token)
			//_, err := demographics.UpdateRole(userRole,token)
			if err != nil {
				fmt.Println("error update Role")
			}
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		app.Session.Put(r.Context(), "userId", userId)
		app.Session.Put(r.Context(), "token", token)
		http.Redirect(w, r, "/support/management/user", 301)
	}

}

type myUserRole struct {
	UserId       string
	RoleId       string
	RoleName     string
	UserName     string
	UserSurname  string
	UserIdNumber string
	Roles        []demographics.Role
}

func getUserRole(userId string) userDomain.UserRole {
	entity := userDomain.UserRole{}
	user, err := userIO.GetUserRole(userId)
	if err != nil {
		return entity
	}
	return user
}
func getRole(roleId string) demographics.Role {
	entity := demographics.Role{}
	role, err := demographics.GetRole(roleId)
	if err != nil {
		return entity
	}
	return role
}
func UserManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRoleList []myUserRole
		users, err := userIO.GetUsers()
		if err != nil {
			fmt.Println("error reading Users")
		}
		applicant, err := applications.GetApplicantTypes()
		if err != nil {
			fmt.Println("error reading applicant")
		}
		roles, err := demographics.GetRoles()
		if err != nil {
			fmt.Println("error reading Roles")
		}
		userRoles, err := userIO.GetUserRoles()
		if err != nil {
			fmt.Println("error reading userRole")
		} else {
			for _, user := range users {
				myUser := getUserRole(user.Email)
				role := getRole(myUser.RoleId)
				userObject := myUserRole{user.Email, role.Id, role.RoleName, user.FirstName, user.LastName, user.IdNumber, roles}
				userRoleList = append(userRoleList, userObject)
			}
		}
		tab := app.Session.GetString(r.Context(), "tab")
		activeTab := myTabs(tab)
		type PageData struct {
			Users         []userDomain.User
			Roles         []demographics.Role
			UserRole      []userDomain.UserRole
			UserRoleList  []myUserRole
			MyActiveTab   Roletabs
			ApplicantType []domain.ApplicantType
		}
		Data := PageData{users, roles, userRoles, userRoleList, activeTab, applicant}
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
		err = ts.Execute(w, Data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

type Roletabs struct {
	Tab1 string
	Tab2 string
	Tab3 string
}

func myTabs(tab string) Roletabs {
	switch tab {
	case "tab1":
		return Roletabs{"active show", "", ""}
	case "tab2":
		return Roletabs{"", "active show", ""}
	case "tab3":
		return Roletabs{"", "", "active show"}
	default:
		return Roletabs{"active show", "", ""}
	}
}
