package schema

const (
	EC2InstanceStatePending = "pending"
	EC2InstanceStateRunning = "running"
)

type EC2Instance struct {
	BaseModel

	OrganizationID uint   `json:"organization_id" gorm:"not null;uniqueIndex:idx_organization_id_cloud_creds_id_instance_id,where:deleted_at IS NULL;index:idx_organization_id_cloud_creds_id_state"`
	CloudCredsID   uint   `json:"cloud_creds_id" gorm:"not null;uniqueIndex:idx_organization_id_cloud_creds_id_instance_id,where:deleted_at IS NULL;index:idx_organization_id_cloud_creds_id_state"`
	InstanceID     string `json:"instance_id" gorm:"not null;uniqueIndex:idx_organization_id_cloud_creds_id_instance_id,where:deleted_at IS NULL"`
	State          string `json:"state" gorm:"not null;index:idx_organization_id_cloud_creds_id_state"`

	VPCID            string `json:"vpc_id" gorm:"not null"`
	SubnetID         string `json:"subnet_id" gorm:"not null"`
	AvailabilityZone string `json:"availability_zone" gorm:"not null"`
	PrivateIpAddress string `json:"private_ip_address" gorm:"not null"`
	PublicIpAddress  string `json:"public_ip_address" gorm:"not null"`
}
