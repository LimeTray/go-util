package awsutil

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/sts"
)

// To get current login details
func GetCallerIdentity() (*sts.GetCallerIdentityOutput, error) {
	svc := sts.New(AWS_SESSION)
	input := &sts.GetCallerIdentityInput{}
	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				Logger.Error(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			Logger.Error(err.Error())
		}
		return nil, err
	}
	return result, nil
}
