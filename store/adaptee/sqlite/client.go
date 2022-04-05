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
}

func (c *Client) Close() {
	c.db.Migrator().DropTable(&schema.Organization{})
	c.db.Migrator().DropTable(&schema.CloudCred{})
	c.db.Migrator().DropTable(&schema.Region{})
	// c.db.Migrator().DropTable(&schema.VPC{})

	db, _ := c.db.DB()
	db.Close()
}

func (c *Client) GetOrgByID(id uint) (*schema.Organization, *errors.AppError) {
	var org schema.Organization
	if err := c.db.First(&org).Where("id = ?", id); err != nil {
		return nil, errors.NotFound("invalid organization id")
	}
	return &org, nil
}

// func (c *Client) GetIntegrationByAPIKey(apiKey string) (*schema.Integration, *errors.AppError) {
// 	if c.integration.APIKey != apiKey {
// 		return nil, errors.NotFound("invalid integration api key")
// 	}
// 	return c.integration, nil
// }

// func (c *Client) GetIntegrationByID(id uint) (*schema.Integration, *errors.AppError) {
// 	if c.integration.ID != id {
// 		return nil, errors.NotFound("invalid integration id")
// 	}
// 	return c.integration, nil
// }

// Topic ...
func (c *Client) Regions() adapter.Regions {
	return c.RegionsConn
}

// Alerts ...
func (c *Client) VPC() adapter.VPC {
	return c.VpcConn
}
