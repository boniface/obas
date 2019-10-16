package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	"obas/config"
)

func Mails(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", MailsHandler(app))
	r.Get("/mailconfig", MailConfigHandler(app))
	r.Get("/smtpconfig", SmtpConfigHandler(app))
	r.Get("/mailapi", MailApiHandler(app))
	return r
}

func MailApiHandler(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func MailsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allmails, err :=io.Get
	}
}

func MailConfigHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allconfig, err := io.Get
	}
}

func SmtpConfigHandler(app *config.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//allSmtp, err := io.Get
	}
}
