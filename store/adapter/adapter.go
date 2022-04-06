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
	EC2Instances() EC2Instances
}

type Regions interface {
	GetByName(name string) (*schema.Region, *errors.AppError)
	All() ([]*schema.Region, *errors.AppError)
}

type VPC interface {
	Save(alert []*schema.VPC) ([]*schema.VPC, *errors.AppError)
}

type EC2Instances interface {
	All() ([]*schema.EC2Instance, *errors.AppError)
	Save(alert []*schema.EC2Instance) ([]*schema.EC2Instance, *errors.AppError)
}
