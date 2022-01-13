package awsutil

import (
	"os"

	"github.com/LimeTray/go-util/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var AWS_SESSION *session.Session
var Logger = logger.RegisterLogger("awsutil")

func CreateGlobalSession() {
	// WILL BE USING ENV FOR AWS CREDENTAILS
	// AWS_ACCESS_KEY_ID
	// AWS_SECRET_ACCESS_KEY

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)
	if err != nil {
		panic(err)
	}
	AWS_SESSION = sess
}
