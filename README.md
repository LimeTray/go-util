# GO-UTIL

Please note this is public repositiry, do not commit any secret and organization related information here  

## Purpose

Every folder is supposed to be a packages, which can be used in go based service/scripts being used in limetray.
Example
 - AWS Credential Helper
 - Custom Logger
 - SQS Consumer Helper
 - Slack Notification

## Sample Usage
```
package main

import (
	logger "github.com/LimeTray/go-util/logger"
)

func main() {
	Logger := logger.RegisterLogger("api-gateway")
	Logger.Info("Application Started")
}
```


# How to user private repo in golang


Update GOPRIVATE variable
```
go env -w GOPRIVATE=github.com/{OrgNameHere}/*
```
If you use ssh to access git repo (locally hosted), you might want to add the following to your ~/.gitconfig
```
[url "ssh://git@git.local.intranet/"]
       insteadOf = https://git.local.intranet/
```
