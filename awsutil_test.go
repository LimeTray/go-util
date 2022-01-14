package main

import (
	"fmt"
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
	instanceId := "i-048a665ab5955f5b8"
	awsutil.CreateGlobalSession()
	instance, err := awsutil.GetEC2MetaByInstanceId(instanceId)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, instance)
	assert.Equal(t, *instance.InstanceId, instanceId)
}

func TestGetTagNameByInstance(t *testing.T) {
	enableDotEnv()
	instanceId := "i-048a665ab5955f5b8"
	awsutil.CreateGlobalSession()
	instance, err := awsutil.GetEC2MetaByInstanceId(instanceId)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, instance)
	name := awsutil.GetTagNameByInstance(instance)
	fmt.Println("Name: " + name)
	assert.NotEmpty(t, name)

}

func TestHostNameByInstanceId(t *testing.T) {
	enableDotEnv()
	instanceId := "i-048a665ab5955f5b8"
	awsutil.CreateGlobalSession()
	instance, err := awsutil.GetEC2MetaByInstanceId(instanceId)
	if err != nil {
		t.Error(err)
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
		t.Log(err.Error())
	} else {
		t.Log(*caller.Arn)
		assert.NotNil(t, caller)
		assert.NotNil(t, *caller.Arn)
	}
}
