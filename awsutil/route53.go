package awsutil

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

func GetHostedZone(id *string) (*route53.GetHostedZoneOutput, error) {
	// To get details of hosted zone
	svc := route53.New(AWS_SESSION)
	input := &route53.GetHostedZoneInput{
		Id: id,
	}
	if result, err := svc.GetHostedZone(input); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// To create/update a record using alias
func CreateARecordAlias(
	route53HostedZoneId *string,
	domainName *string,
	value *string,
	valueHostedZone *string,
) error {
	Logger.Info(fmt.Sprintf("%s %s %s %s", *route53HostedZoneId, *domainName, *value, *valueHostedZone))
	// To create/update a record
	input := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: route53HostedZoneId,
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String(route53.ChangeActionUpsert),
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              value,
							HostedZoneId:         valueHostedZone,
							EvaluateTargetHealth: aws.Bool(true),
						},
						Name: domainName,
						Type: aws.String(route53.RRTypeA),
					},
				},
			},
		},
	}

	svc := route53.New(AWS_SESSION)
	if result, err := svc.ChangeResourceRecordSets(input); err != nil {
		Logger.Error(err.Error())
		return err
	} else {
		Logger.Info(*result.ChangeInfo.Status)
	}
	return nil
}
