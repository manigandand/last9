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
