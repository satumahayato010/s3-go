package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

type Session struct {
	S3Session *session.Session
}
