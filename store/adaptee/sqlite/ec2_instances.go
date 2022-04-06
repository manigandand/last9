package sqlite

import (
	"last9/errors"
	"last9/schema"

	"gorm.io/gorm/clause"
)

// EC2Inst implements EC2Inst adapter interface
type EC2Inst struct {
	*Client
}

// NewEC2InstStore ...
func NewEC2InstStore(client *Client) *EC2Inst {
	return &EC2Inst{client}
}

func (v *EC2Inst) All() ([]*schema.EC2Instance, *errors.AppError) {
	var ec2Insts []*schema.EC2Instance
	if err := v.db.Find(&ec2Insts, "organization_id = ? AND cloud_creds_id = ?", 1, 1).Error; err != nil {
		return nil, errors.InternalServer("failed to get ec2Insts").AddDebug(err)
	}
	return ec2Insts, nil
}

func (v *EC2Inst) Save(ec2Insts []*schema.EC2Instance) ([]*schema.EC2Instance, *errors.AppError) {
	if len(ec2Insts) == 0 {
		return nil, errors.BadRequest("no ec2Insts to save")
	}
	if err := v.db.Clauses(clause.OnConflict{DoNothing: true}).
		Create(ec2Insts).Error; err != nil {
		return nil, errors.InternalServer("failed to save ec2 instances").AddDebug(err)
	}
	return ec2Insts, nil
}
