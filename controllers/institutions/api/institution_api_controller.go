package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"obas/config"
	institutionDomain "obas/domain/institutions"
	institutionIO "obas/io/institutions"
)

func InstitutionAPI(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/getInstitutionsInLocation/{locationId}", GetInstitutionsInLocationHandler(app))

	return r
}

func GetInstitutionsInLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locationId := chi.URLParam(r, "locationId")
		var institutions []institutionDomain.Institution
		institutionLocations, err := institutionIO.GetInstitutionsInLocation(locationId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else {
			for _, institutionLocation := range institutionLocations {
				institution, err := institutionIO.GetInstitution(institutionLocation.InstitutionId)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				} else {
					institutions = append(institutions, institution)
				}
			}
		}
		render.JSON(w, r, institutions)
	}
}
