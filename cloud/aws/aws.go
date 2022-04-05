package aws

import (
	"last9/errors"
	"last9/schema"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Options struct {
	Region string
}

type AWS struct {
	region     string
	ec2Svc     *ec2.EC2
	eksSvc     *eks.EKS
	iamSvc     *iam.IAM
	s3Svc      *s3.S3
	sess       *session.Session
	cloudCreds *schema.CloudCred
	apiKey     string
	secretKey  string
}

// New - create a new AWS client
func NewOrchestrator(cloudCreds *schema.CloudCred) (*AWS, *errors.AppError) {
	primaryConfig := aws.NewConfig().
		WithRegion(cloudCreds.Region).
		WithCredentials(credentials.NewStaticCredentials(cloudCreds.APIKey, cloudCreds.SecretKey, ""))
	sess, err := session.NewSession(primaryConfig)
	if err != nil {
		return nil, errors.InternalServer("aws.NewOrchestrator failed to create session").AddDebug(err)
	}

	return &AWS{
		region:     cloudCreds.Region,
		ec2Svc:     ec2.New(sess),
		eksSvc:     eks.New(sess),
		iamSvc:     iam.New(sess),
		s3Svc:      s3.New(sess),
		sess:       sess,
		cloudCreds: cloudCreds,
		apiKey:     cloudCreds.APIKey,
		secretKey:  cloudCreds.SecretKey,
	}, nil
}

func (a *AWS) GetCloudType() string {
	return schema.CloudTypeAWS.String()
}

func (a *AWS) GetRegions() ([]*schema.Region, *errors.AppError) {
	regions := []*schema.Region{}
	input := &ec2.DescribeRegionsInput{
		AllRegions: aws.Bool(true),
	}
	result, err := a.ec2Svc.DescribeRegions(input)
	if err != nil {
		return nil, errors.InternalServer("aws.GetRegions failed to describe regions").AddDebug(err)
	}
	for _, region := range result.Regions {
		regions = append(regions, &schema.Region{
			OrganizationID: a.cloudCreds.OrganizationID,
			CloudCredsID:   a.cloudCreds.ID,
			Name:           *region.RegionName,
			Endpoint:       *region.Endpoint,
			OptInStatus:    *region.OptInStatus,
		})
	}
	return regions, nil
}

func (a *AWS) GetVPC() ([]string, *errors.AppError) {
	return nil, nil
}
