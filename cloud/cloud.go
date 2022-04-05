package cloud

import (
	"last9/cloud/aws"
	"last9/errors"
	"last9/schema"
)

type Cloud interface {
	GetCloudType() string
	GetRegions() ([]*schema.Region, *errors.AppError)
	GetVPC() ([]*schema.VPC, *errors.AppError)
}

func NewCloud(cloudCreds *schema.CloudCred) (Cloud, *errors.AppError) {
	switch cloudCreds.Type {
	case schema.CloudTypeAWS:
		return aws.NewOrchestrator(cloudCreds)
	case schema.CloudTypeGCP:
		return nil, errors.BadRequest("gcp not implemented")
	case schema.CloudTypeAzure:
		return nil, errors.BadRequest("gcp not implemented")
	default:
		return nil, errors.BadRequest("invalid cloud type")
	}
}
