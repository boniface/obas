package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	"obas/config"
)

func Storages(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/file", FileHanler(app))
	r.Get("/file/size", FileSizeHanler(app))
	r.Get("/file/upload", FileUploadHanler(app))
	r.Get("/file/delete", FileDeleteHanler(app))
	return r
}

func FileHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func FileSizeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func FileUploadHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func FileDeleteHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
