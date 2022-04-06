### last9

> Directory structure

```
- main.go
--- api: region, vpc related apis.
--- config: app level config pkg
--- errors: custom app error pkg
--- response: custom http response pkg
--- schema: db schema
--- store: data store layer
--- task: task worker pool maintainer & dispatcher.
    --- worker: sub workers which subscribed to events and does the specfic jobs.
```

> How to run

`make run`

> darwin build

or

```
Just build and run.

go build && ./last9 config.json

```

> API's

> Discover ec2 instances
> `GET localhost:8080/v1/aws/eu-west-3/instances`

```json
{
  "data": [
    {
      "id": 1,
      "created_at": "2022-04-06T13:53:43.850586+05:30",
      "updated_at": "2022-04-06T13:53:43.850586+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "instance_id": "i-014087f64fe2a4a23",
      "state": "running",
      "vpc_id": "vpc-07ad51c832db95af4",
      "subnet_id": "subnet-0a8920b561648eccc",
      "availability_zone": "eu-west-3a",
      "private_ip_address": "172.31.13.97",
      "public_ip_address": "35.180.187.164"
    }
  ],
  "meta": {
    "status_code": 200
  }
}
```

> Discover vpcs in region
> `GET localhost:8080/v1/{aws}/{us-west-2}/vpcs`

```json
{
  "data": [
    {
      "id": 0,
      "created_at": "2022-04-06T03:35:21.73962+05:30",
      "updated_at": "2022-04-06T03:35:21.73962+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "vpc_id": "vpc-0b85bd18b306f9d19",
      "cidr": "172.31.0.0/16",
      "state": "available"
    }
  ],
  "meta": {
    "status_code": 200
  }
}
```

> GET all regions
> `POST localhost:8080/v1/{aws}/regions`

```json
{
  "data": [
    {
      "id": 1,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "af-south-1",
      "endpoint": "ec2.af-south-1.amazonaws.com",
      "opt_in_status": "not-opted-in"
    },
    {
      "id": 2,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "eu-north-1",
      "endpoint": "ec2.eu-north-1.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 3,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "ap-south-1",
      "endpoint": "ec2.ap-south-1.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 4,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "eu-west-3",
      "endpoint": "ec2.eu-west-3.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 5,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "eu-west-2",
      "endpoint": "ec2.eu-west-2.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 6,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "eu-south-1",
      "endpoint": "ec2.eu-south-1.amazonaws.com",
      "opt_in_status": "not-opted-in"
    },
    {
      "id": 7,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "eu-west-1",
      "endpoint": "ec2.eu-west-1.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 8,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "ap-northeast-3",
      "endpoint": "ec2.ap-northeast-3.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 9,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "ap-northeast-2",
      "endpoint": "ec2.ap-northeast-2.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 10,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "me-south-1",
      "endpoint": "ec2.me-south-1.amazonaws.com",
      "opt_in_status": "not-opted-in"
    },
    {
      "id": 11,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "ap-northeast-1",
      "endpoint": "ec2.ap-northeast-1.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 12,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "sa-east-1",
      "endpoint": "ec2.sa-east-1.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 13,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "ca-central-1",
      "endpoint": "ec2.ca-central-1.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 14,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "ap-east-1",
      "endpoint": "ec2.ap-east-1.amazonaws.com",
      "opt_in_status": "not-opted-in"
    },
    {
      "id": 15,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "ap-southeast-1",
      "endpoint": "ec2.ap-southeast-1.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 16,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "ap-southeast-2",
      "endpoint": "ec2.ap-southeast-2.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 17,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "eu-central-1",
      "endpoint": "ec2.eu-central-1.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 18,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "ap-southeast-3",
      "endpoint": "ec2.ap-southeast-3.amazonaws.com",
      "opt_in_status": "not-opted-in"
    },
    {
      "id": 19,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "us-east-1",
      "endpoint": "ec2.us-east-1.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 20,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "us-east-2",
      "endpoint": "ec2.us-east-2.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 21,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "us-west-1",
      "endpoint": "ec2.us-west-1.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    },
    {
      "id": 22,
      "created_at": "2022-04-06T02:21:09.770882+05:30",
      "updated_at": "2022-04-06T02:21:09.770882+05:30",
      "deleted_at": null,
      "organization_id": 1,
      "cloud_creds_id": 1,
      "name": "us-west-2",
      "endpoint": "ec2.us-west-2.amazonaws.com",
      "opt_in_status": "opt-in-not-required"
    }
  ],
  "meta": {
    "status_code": 200
  }
}
```

```
aws ec2 create-vpc \
    --cidr-block 10.0.0.0/16 \
    --instance-tenancy dedicated \
    --tag-specifications ResourceType=vpc,Tags='[{Key=Owner,Value="manigandand"}]'


aws ec2 run-instances \
    --image-id ami-000037ee86985ff3b \
    --instance-type t2.micro \
    --subnet-id subnet-0a8920b561648eccc \
    --count 1 \
    --tag-specifications 'ResourceType=instance,Tags=[{Key=owner,Value=manigandand}]' 'ResourceType=volume,Tags=[{Key=cost-center,Value=cc123}]'
```
