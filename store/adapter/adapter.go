package adapter

import (
	"last9/errors"
	"last9/schema"
)

type Store interface {
	Close()
	GetOrgByID(id uint) (*schema.Organization, *errors.AppError)
	GetCloudCredByID(id uint) (*schema.CloudCred, *errors.AppError)
	Regions() Regions
	VPC() VPC
}

type Regions interface {
	GetByName(name string) (*schema.Region, *errors.AppError)
	All() ([]*schema.Region, *errors.AppError)
}

type VPC interface {
	// All() ([]*schema.Alert, *errors.AppError)
	Save(alert []*schema.VPC) ([]*schema.VPC, *errors.AppError)
}
