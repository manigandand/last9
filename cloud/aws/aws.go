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
	region    string
	ec2Svc    *ec2.EC2
	eksSvc    *eks.EKS
	iamSvc    *iam.IAM
	s3Svc     *s3.S3
	sess      *session.Session
	apiKey    string
	secretKey string
}

// New - create a new AWS client
func NewOrchestrator(opt *Options) (*AWS, *errors.AppError) {
	primaryApiKey := ""
	primarySecretKey := "/"

	primaryConfig := aws.NewConfig().
		WithRegion(opt.Region).
		WithCredentials(credentials.NewStaticCredentials(primaryApiKey, primarySecretKey, ""))
	sess, err := session.NewSession(primaryConfig)
	if err != nil {
		return nil, errors.InternalServer("aws.NewOrchestrator failed to create session").AddDebug(err)
	}

	return &AWS{
		region:    opt.Region,
		ec2Svc:    ec2.New(sess),
		eksSvc:    eks.New(sess),
		iamSvc:    iam.New(sess),
		s3Svc:     s3.New(sess),
		sess:      sess,
		apiKey:    primaryApiKey,
		secretKey: primarySecretKey,
	}, nil
}

func (a *AWS) GetCloudType() string {
	return schema.CloudTypeAWS.String()
}

func (a *AWS) GetRegions() ([]string, *errors.AppError) {
	regions := []string{}
	input := &ec2.DescribeRegionsInput{
		AllRegions: aws.Bool(true),
	}
	result, err := a.ec2Svc.DescribeRegions(input)
	if err != nil {
		return nil, errors.InternalServer("aws.GetRegions failed to describe regions").AddDebug(err)
	}
	for _, region := range result.Regions {
		regions = append(regions, *region.RegionName)
	}
	return regions, nil
}

func (a *AWS) GetVPC() ([]string, *errors.AppError) {
	return nil, nil
}
