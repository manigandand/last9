package schema

import "strings"

type Cloudtype string

const (
	CloudTypeAWS   Cloudtype = "aws"
	CloudTypeGCP   Cloudtype = "gcp"
	CloudTypeAzure Cloudtype = "azure"
)

func (c Cloudtype) String() string {
	return string(c)
}

func (c Cloudtype) ToUpper() string {
	return strings.ToUpper(c.String())
}

var ValidClouds = map[string]bool{
	CloudTypeAWS.String(): true,
	// CloudTypeGCP.String():   true,
	// CloudTypeAzure.String(): true,
}

type CloudCred struct {
	BaseModel
	OrganizationID uint      `json:"organization_id" gorm:"not null"`
	Type           Cloudtype `json:"cloud_type" gorm:"not null"`
	APIKey         string    `json:"api_key" gorm:"not null"`
	SecretKey      string    `json:"secret_key" gorm:"not null"`

	Region string `json:"-" gorm:"-"`
}

func (c *CloudCred) SetRegion(name string) {
	c.Region = name
}

func (c *CloudCred) GetRegion() string {
	return c.Region
}
