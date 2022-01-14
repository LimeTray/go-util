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

	if err := awsutil.CreateNewDomain(
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
	if name, err := awsutil.GetGatewayDomainByName(&domainName); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(name)
		assert.NotNil(t, name)
		assert.NotEmpty(t, name)
		assert.Contains(t, name, "amazonaws.com")
	}
}

func TestCreateApiMapping(t *testing.T) {
	enableDotEnv()
	awsutil.CreateGlobalSession()

	domainName := os.Getenv("TEST_DOMAIN_NAME")
	apiId := os.Getenv("TEST_API_ID")
	stage := os.Getenv("TEST_STAGE")
	if err := awsutil.CreateApiMapping(&apiId, &domainName, &stage); err != nil {
		t.Fatal(err.Error())
	}
}
