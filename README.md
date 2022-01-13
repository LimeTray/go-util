# GO-UTIL

## Purpose

Every folder is supposed to be a packages, which can be used in go based service/scripts being used in limetray.
Example
 - AWS Credential Helper
 - Custom Logger
 - SQS Consumer Helper
 - Slack Notification


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
