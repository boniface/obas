package controllers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"obas/config"
	addressControllers "obas/controllers/address"
	applicationControllers "obas/controllers/application"
	demographicsControllers "obas/controllers/demographics"
	documentsControllers "obas/controllers/documents"
	homeControllers "obas/controllers/home"
	institutionsControllers "obas/controllers/institutions"
	locationControllers "obas/controllers/location"
	logControllers "obas/controllers/log"
	loginControllers "obas/controllers/login"
	logoutControllers "obas/controllers/logout"
	registerControllers "obas/controllers/register"
	subjectsControllers "obas/controllers/subjects"
	usersControllers "obas/controllers/users"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	mux.Handle("/", homeControllers.Home(env))
	mux.Mount("/login", loginControllers.Login(env))
	mux.Mount("/logout", logoutControllers.Logout(env))
	mux.Mount("/register", registerControllers.Register(env))
	mux.Mount("/users", usersControllers.Users(env))
	mux.Mount("/subjects", subjectsControllers.Subjects(env))
	mux.Mount("/address", addressControllers.Addresses(env))
	mux.Mount("/demographics", demographicsControllers.Demographics(env))
	mux.Mount("/application", applicationControllers.Applications(env))
	mux.Mount("/documents", documentsControllers.Documents(env))
	mux.Mount("/institution", institutionsControllers.Institutions(env))
	mux.Mount("/location", locationControllers.Locations(env))
	mux.Mount("/log", logControllers.Logs(env))

	fileServer := http.FileServer(http.Dir("./views/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux

}
