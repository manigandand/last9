package sqlite

import (
	"last9/cloud"
	"last9/config"
	"last9/schema"
	"last9/store/adapter"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewAdapter returns store sqlite adapter(*Client)
func NewAdapter() adapter.Store {
	db, err := gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Load Data
	c := &Client{
		db: db,
	}
	c.RegionsConn = NewRegionsStore(c)
	c.VpcConn = NewVPCStore(c)
	c.EC2InstConn = NewEC2InstStore(c)

	schemas := []struct {
		name   string
		schema interface{}
	}{
		{"organization", &schema.Organization{}},
		{"cloud_creds", &schema.CloudCred{}},
		{"regions", &schema.Region{}},
		{"vpcs", &schema.VPC{}},
		{"ec2_instances", &schema.EC2Instance{}},
	}
	for _, table := range schemas {
		if err := db.AutoMigrate(table.schema); err != nil {
			log.Printf("Failed to set up DB table(%s) with error: %s\n", table.name, err.Error())
		}
	}

	c.loadCloudConfig()

	return c
}

func (c *Client) loadCloudConfig() {
	if config.AWSAPIID == "" || config.AWSSecretKey == "" {
		panic("AWS credentials not set")
	}

	org := &schema.Organization{
		Name:         "Last9 Inc",
		Slug:         "last9",
		Subscription: "enterprise-monthly",
	}
	if err := c.db.Create(org).Error; err != nil {
		panic(err)
	}

	cloudCreds := &schema.CloudCred{
		OrganizationID: org.ID,
		Type:           schema.CloudTypeAWS,
		APIKey:         config.AWSAPIID,
		SecretKey:      config.AWSSecretKey,
	}
	if err := c.db.Create(cloudCreds).Error; err != nil {
		panic(err)
	}

	cloudCreds.SetRegion("eu-west-3")
	ch, err := cloud.NewCloud(cloudCreds)
	if err != nil {
		panic(err)
	}

	regions, err := ch.GetRegions()
	if err != nil {
		panic(err)
	}
	if err := c.db.Create(regions).Error; err != nil {
		panic(err)
	}
}
