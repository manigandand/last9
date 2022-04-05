package cloud

import (
	"last9/cloud/aws"
	"last9/errors"
	"last9/schema"
)

type Cloud interface {
	GetCloudType() string
	GetRegions() ([]string, *errors.AppError)
	GetVPC() ([]string, *errors.AppError)
}

func NewCloud(ctype schema.Cloudtype, opt *aws.Options) (Cloud, *errors.AppError) {
	switch ctype {
	case schema.CloudTypeAWS:
		return aws.NewOrchestrator(opt)
	case schema.CloudTypeGCP:
		return nil, errors.BadRequest("gcp not implemented")
	case schema.CloudTypeAzure:
		return nil, errors.BadRequest("gcp not implemented")
	default:
		return nil, errors.BadRequest("invalid cloud type")
	}
}
