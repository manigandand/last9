package schema

type VPC struct {
	BaseModel

	OrganizationID uint `json:"organization_id" gorm:"not null;uniqueIndex:idx_organization_id_cloud_creds_id_vpc_id,where:deleted_at IS NULL;index:idx_organization_id_cloud_creds_id"`
	CloudCredsID   uint `json:"cloud_creds_id" gorm:"not null;uniqueIndex:idx_organization_id_cloud_creds_id_vpc_id,where:deleted_at IS NULL;index:idx_organization_id_cloud_creds_id"`

	VPCID string `json:"vpc_id" gorm:"not null;uniqueIndex:idx_organization_id_cloud_creds_id_vpc_id,where:deleted_at IS NULL;index:idx_organization_id_cloud_creds_id"`
	CIDR  string `json:"cidr" gorm:"not null"`
	State string `json:"state" gorm:"not null"`
}
