package awsutil

import (
	"errors"
	
	"strings"

	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetEC2MetaByInstanceId(instanceId string) (*ec2.Instance, error) {
	svc := ec2.New(AWS_SESSION)
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
	}
	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				Logger.Error(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			Logger.Error(aerr.Error())
		}
		return nil, err
	}
	if len(result.Reservations) == 0 || len(result.Reservations[0].Instances) == 0 {
		return nil, errors.New("invalid instance")
	}
	return result.Reservations[0].Instances[0], nil
}

// To get private dns name as host name
func GetHostNameByInstanceId(instance *ec2.Instance) string {
	return *instance.PrivateDnsName
}

// To extract specific tag from instance tag
func GetTagValueByInstace(instance *ec2.Instance, key string) string {
	val := ""
	tags := instance.Tags
	if len(tags) > 0 {
		for _, t := range tags {
			if strings.EqualFold(strings.ToLower(*t.Key), strings.ToLower(key)) {
				val = *t.Value
				break
			}
		}
	}
	return val
}

// To get name tag from instance
func GetTagNameByInstance(instance *ec2.Instance) string {
	return GetTagValueByInstace(instance, "name")
}

// To get ASG tag from instance
func GetASGByInstance(instance *ec2.Instance) string {
	return GetTagValueByInstace(instance, "aws:autoscaling:groupName")
}
