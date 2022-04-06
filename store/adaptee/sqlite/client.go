package sqlite

import (
	"last9/errors"
	"last9/schema"
	"last9/store/adapter"

	"gorm.io/gorm"
)

// Client struct implements the store adapter interface
type Client struct {
	db *gorm.DB
	// integration     *schema.Integration
	// alertConfigs    map[string]*schema.AlertConfig
	// alerts          []*schema.Alert
	RegionsConn adapter.Regions
	VpcConn     adapter.VPC
	EC2InstConn adapter.EC2Instances
}

func (c *Client) Close() {
	c.db.Migrator().DropTable(&schema.Organization{})
	c.db.Migrator().DropTable(&schema.CloudCred{})
	c.db.Migrator().DropTable(&schema.Region{})
	c.db.Migrator().DropTable(&schema.VPC{})
	c.db.Migrator().DropTable(&schema.EC2Instance{})

	db, _ := c.db.DB()
	db.Close()
}

func (c *Client) GetOrgByID(id uint) (*schema.Organization, *errors.AppError) {
	var org schema.Organization
	if err := c.db.First(&org).Where("id = ?", id).Error; err != nil {
		return nil, errors.NotFound("invalid organization id")
	}
	return &org, nil
}

func (c *Client) GetCloudCredByID(id uint) (*schema.CloudCred, *errors.AppError) {
	var cc schema.CloudCred
	if err := c.db.First(&cc).Where("id = ? AND organization_id = ?", id, 1).Error; err != nil {
		return nil, errors.NotFound("invalid cloud creds id")
	}
	return &cc, nil
}

// Topic ...
func (c *Client) Regions() adapter.Regions {
	return c.RegionsConn
}

// Alerts ...
func (c *Client) VPC() adapter.VPC {
	return c.VpcConn
}

func (c *Client) EC2Instances() adapter.EC2Instances {
	return c.EC2InstConn
}
