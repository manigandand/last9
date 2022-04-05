package api

import (
	"last9/cloud"
	"last9/cloud/aws"
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
	r.Method(http.MethodGet, "/v1/regions", Handler(getAllAlerts))
}

func getAllAlerts(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ch, err := cloud.NewCloud(schema.CloudTypeAWS, &aws.Options{
		Region: "eu-west-3",
	})
	if err != nil {
		return err
	}
	regions, err := ch.GetRegions()
	if err != nil {
		return err
	}

	response.OK(w, regions)
	return nil
}
