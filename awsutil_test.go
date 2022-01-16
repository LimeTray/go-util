package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/LimeTray/go-util/awsutil"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func enableDotEnv() {
	godotenv.Load()
}
func TestCreateGlobalSession(t *testing.T) {
	enableDotEnv()
	awsutil.CreateGlobalSession()
	assert.NotNil(t, awsutil.AWS_SESSION)
}

func TestGetEC2MetaByInstanceId(t *testing.T) {
	enableDotEnv()
	instanceId := os.Getenv("TEST_INSTANCE_ID")
	awsutil.CreateGlobalSession()
	instance, err := awsutil.GetEC2MetaByInstanceId(instanceId)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, instance)
	assert.Equal(t, *instance.InstanceId, instanceId)
}

func TestGetTagNameByInstance(t *testing.T) {
	enableDotEnv()
	instanceId := os.Getenv("TEST_INSTANCE_ID")
	awsutil.CreateGlobalSession()
	instance, err := awsutil.GetEC2MetaByInstanceId(instanceId)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, instance)
	name := awsutil.GetTagNameByInstance(instance)
	fmt.Println("Name: " + name)
	assert.NotEmpty(t, name)

}

func TestHostNameByInstanceId(t *testing.T) {
	enableDotEnv()
	instanceId := os.Getenv("TEST_INSTANCE_ID")
	awsutil.CreateGlobalSession()
	instance, err := awsutil.GetEC2MetaByInstanceId(instanceId)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, instance)
	privateDns := awsutil.GetHostNameByInstanceId(instance)
	fmt.Println("Private DNS: " + privateDns)
	assert.NotEmpty(t, privateDns)
	assert.Contains(t, privateDns, "compute.internal")
}

func TestGetCallerIdentity(t *testing.T) {
	enableDotEnv()
	awsutil.CreateGlobalSession()
	if caller, err := awsutil.GetCallerIdentity(); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(*caller.Arn)
		assert.NotNil(t, caller)
		assert.NotNil(t, *caller.Arn)
	}
}

func TestCreateDomainName(t *testing.T) {
	enableDotEnv()
	awsutil.CreateGlobalSession()

	domainName := os.Getenv("TEST_DOMAIN_NAME")
	arn := os.Getenv("TEST_CERT_ARN")

	if err := awsutil.CreateNewGatewayDomain(
		&domainName,
		&arn,
	); err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetGatewayDomainByName(t *testing.T) {
	enableDotEnv()
	awsutil.CreateGlobalSession()

	domainName := os.Getenv("TEST_DOMAIN_NAME")
	if hostedZoneId, domain, err := awsutil.GetGatewayDomainByName(&domainName); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(hostedZoneId)
		assert.NotNil(t, hostedZoneId)
		assert.NotEmpty(t, hostedZoneId)
		t.Log(domain)
		assert.NotNil(t, domain)
		assert.NotEmpty(t, domain)
		assert.Contains(t, domain, "amazonaws.com")
	}
}

func TestCreateApiMapping(t *testing.T) {
	enableDotEnv()
	awsutil.CreateGlobalSession()

	domainName := os.Getenv("TEST_DOMAIN_NAME")
	apiId := os.Getenv("TEST_API_ID")
	stage := os.Getenv("TEST_STAGE")
	if err := awsutil.CreateGatewayApiMapping(&apiId, &domainName, &stage); err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetHostedZone(t *testing.T) {
	enableDotEnv()
	awsutil.CreateGlobalSession()
	hostedZoneId := ""

	if hostedZone, err := awsutil.GetHostedZone(&hostedZoneId); err != nil {
		t.Fatal(err.Error())
	} else {
		assert.NotNil(t, hostedZone)
		assert.Equal(t, &hostedZone.HostedZone.Id, hostedZoneId)
		assert.NotNil(t, hostedZone.HostedZone.Name)
	}
}

func TestCreateARecordAlias(t *testing.T) {
	enableDotEnv()
	awsutil.CreateGlobalSession()
	hostedZoneId := os.Getenv("TEST_ROUTE53_HOSTEDZONE_ID")
	domainName := os.Getenv("TEST_DOMAIN_NAME")
	value := os.Getenv("TEST_GATEWAY_DOMAIN")
	valueHostedZoneID := os.Getenv("TEST_GATEWAY_HOSTEDZONE_ID")

	if err := awsutil.CreateARecordAlias(
		&hostedZoneId,
		&domainName,
		&value,
		&valueHostedZoneID,
	); err != nil {
		t.Fatal(err.Error())
	}
}
