package notification

import (
	"os"
)

var (
	DEFAIULT_SLACK_URL      string = "https://hooks.slack.com/services/xxx/xxx/xxx"
	DEFAIULT_SLACK_USERNAME string = "go-util"
	DEFAIULT_SLACK_CHANNEL  string = "#technology"
)

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
