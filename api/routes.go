package api

import (
	"last9/cloud"
	"last9/errors"
	"last9/response"
	"last9/schema"
	"last9/task"
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
	r.Method(http.MethodGet, "/v1/{cloud_type}/regions", Handler(getAllAlerts))
	r.Method(http.MethodGet, "/v1/{cloud_type}/{region_name}/vpcs", Handler(discoverVPCByRegion))
	r.Method(http.MethodGet, "/v1/{cloud_type}/{region_name}/instances", Handler(discoverInstancesByRegion))
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

func discoverVPCByRegion(w http.ResponseWriter, r *http.Request) *errors.AppError {
	cloudType := chi.URLParam(r, "cloud_type")
	if ok := schema.ValidClouds[cloudType]; !ok {
		return errors.BadRequest("invalid cloud type")
	}
	regionName := chi.URLParam(r, "region_name")
	region, err := store.Regions().GetByName(regionName)
	if err != nil {
		return err
	}

	// get the first cloud creds
	cloudCreds, err := store.GetCloudCredByID(1)
	if err != nil {
		return err
	}
	cloudCreds.SetRegion(region.Name)
	ch, err := cloud.NewCloud(cloudCreds)
	if err != nil {
		return err
	}
	vpcs, err := ch.DiscoverVPC()
	if err != nil {
		return err
	}

	vpcRes, err := store.VPC().Save(vpcs)
	if err != nil {
		return err
	}

	response.OK(w, vpcRes)
	return nil
}

func discoverInstancesByRegion(w http.ResponseWriter, r *http.Request) *errors.AppError {
	cloudType := chi.URLParam(r, "cloud_type")
	if ok := schema.ValidClouds[cloudType]; !ok {
		return errors.BadRequest("invalid cloud type")
	}
	regionName := chi.URLParam(r, "region_name")
	region, err := store.Regions().GetByName(regionName)
	if err != nil {
		return err
	}

	// get the first cloud creds
	cloudCreds, err := store.GetCloudCredByID(1)
	if err != nil {
		return err
	}
	cloudCreds.SetRegion(region.Name)
	ch, err := cloud.NewCloud(cloudCreds)
	if err != nil {
		return err
	}
	ec2Insts, err := ch.DiscoverInstances()
	if err != nil {
		return err
	}
	if len(ec2Insts) == 0 {
		return errors.BadRequest("no ec2 instances discovered")
	}

	ec2InstsRes, err := store.EC2Instances().Save(ec2Insts)
	if err != nil {
		return err
	}

	go task.New(task.EventTypeNewInstance).Dispatch()

	response.OK(w, ec2InstsRes)
	return nil
}
