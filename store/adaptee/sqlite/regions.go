package sqlite

import (
	"last9/errors"
	"last9/schema"
)

// Regions implements Regions adapter interface
type Regions struct {
	*Client
}

// NewRegionsStore ...
func NewRegionsStore(client *Client) *Regions {
	return &Regions{client}
}

func (r *Regions) GetByName(name string) (*schema.Region, *errors.AppError) {
	var region schema.Region
	if err := r.db.First(&region, "organization_id = ? AND cloud_creds_id = ? AND name = ?", 1, 1, name).Error; err != nil {
		return nil, errors.NotFound("invalid region name")
	}
	return &region, nil
}

func (r *Regions) All() ([]*schema.Region, *errors.AppError) {
	var regions []*schema.Region
	if err := r.db.Find(&regions, "organization_id = ? AND cloud_creds_id = ?", 1, 1).Error; err != nil {
		return nil, errors.InternalServer("failed to get regions").AddDebug(err)
	}
	return regions, nil
}
