package sqlite

import (
	"last9/errors"
	"last9/schema"

	"gorm.io/gorm/clause"
)

// VPC implements VPC adapter interface
type VPC struct {
	*Client
}

// NewVPCStore ...
func NewVPCStore(client *Client) *VPC {
	return &VPC{client}
}

func (v *VPC) Save(vpcs []*schema.VPC) ([]*schema.VPC, *errors.AppError) {
	if err := v.db.Clauses(clause.OnConflict{DoNothing: true}).
		Create(vpcs).Error; err != nil {
		return nil, errors.InternalServer("failed to save vpcs").AddDebug(err)
	}
	return vpcs, nil
}
