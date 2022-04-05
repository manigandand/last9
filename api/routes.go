package api

import (
	"last9/errors"
	"last9/response"
	"last9/schema"
	"net/http"

	"github.com/go-chi/chi"
)

// Routes - all the registered routes
func Routes(router chi.Router) {
	router.Get("/", IndexHandeler)
	router.Get("/top", HealthHandeler)
	router.Route("/", InitV1Routes)
}

func InitV1Routes(r chi.Router) {
	// r.Method(http.MethodPost, "/{version}/alert/{slug}/{apiKey}", Handler(handleEvent))
	// r.Method(http.MethodGet, "/v1/alert-sources", Handler(getAllAlertSources))
	r.Method(http.MethodGet, "/v1/{cloud_type}/regions", Handler(getAllAlerts))
}

func getAllAlerts(w http.ResponseWriter, r *http.Request) *errors.AppError {
	cloudType := chi.URLParam(r, "cloud_type")
	if ok := schema.ValidClouds[cloudType]; !ok {
		return errors.BadRequest("invalid cloud type")
	}

	regions, err := store.Regions().All()
	if err != nil {
		return err
	}

	response.OK(w, regions)
	return nil
}
