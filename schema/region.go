package schema

type Region struct {
	BaseModel
	OrganizationID uint `json:"organization_id" gorm:"not null;uniqueIndex:idx_organization_id_cloud_creds_id_name,where:deleted_at IS NULL;index:idx_organization_id_cloud_creds_id"`
	CloudCredsID   uint `json:"cloud_creds_id" gorm:"not null;uniqueIndex:idx_organization_id_cloud_creds_id_name,where:deleted_at IS NULL;index:idx_organization_id_cloud_creds_id"`

	Name        string `json:"name" gorm:"not null;uniqueIndex:idx_organization_id_cloud_creds_id_name,where:deleted_at IS NULL"`
	Endpoint    string `json:"endpoint" gorm:"not null"`
	OptInStatus string `json:"opt_in_status" gorm:"not null"`
}
