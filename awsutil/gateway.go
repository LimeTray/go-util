package awsutil

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
)

func getGatewayClient() *apigatewayv2.ApiGatewayV2 {
	svc := apigatewayv2.New(AWS_SESSION)
	return svc
}

func CreateNewDomain(domainname *string, certArn *string) error {
	input := apigatewayv2.CreateDomainNameInput{
		DomainName: domainname,
		DomainNameConfigurations: []*apigatewayv2.DomainNameConfiguration{
			{
				CertificateArn: certArn,
				SecurityPolicy: aws.String(apigatewayv2.SecurityPolicyTls12),  // Hard coded change later if required
				EndpointType:   aws.String(apigatewayv2.EndpointTypeRegional), // Hard coded change later if required
			},
		},
	}

	svc := getGatewayClient()
	if result, err := svc.CreateDomainName(&input); err != nil {
		Logger.Error(err.Error())
		return err
	} else {
		Logger.Info(*result.DomainNameConfigurations[0].ApiGatewayDomainName)
	}
	return nil
}

func CreateApiMapping(apiId *string, domainName *string, stage *string) error {
	// To map api gateway to gateway domain name
	svc := getGatewayClient()
	input := apigatewayv2.CreateApiMappingInput{
		ApiId:      apiId,
		DomainName: domainName,
		Stage:      stage,
	}
	if result, err := svc.CreateApiMapping(&input); err != nil {
		Logger.Error(err.Error())
		return err
	} else {
		Logger.Info(*result.ApiMappingId)
	}
	return nil
}

func GetGatewayDomainByName(domainName *string) (string, error) {
	// To get regional id of the domain name
	// This will be later used to map regional domain to route53
	svc := getGatewayClient()
	input := apigatewayv2.GetDomainNameInput{
		DomainName: domainName,
	}
	if result, err := svc.GetDomainName(&input); err != nil {
		Logger.Error(err.Error())
		return "", err
	} else {
		// If no config is found
		if len(result.DomainNameConfigurations) == 0 {
			return "", errors.New("DomainNameConfigurations not found")
		}
		// Return gateway domain
		return *result.DomainNameConfigurations[0].ApiGatewayDomainName, nil
	}
}
