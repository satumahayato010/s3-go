package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Session struct {
	S3Session *session.Session
}

func main() {
	paths := []string{"", ""}
	credential := credentials.NewStaticCredentials(
		os.Getenv("SECRET_ID"),
		os.Getenv("SECRET_KEY"),
		"",
	)

	awsConfig := aws.Config{
		Region:      aws.String(os.Getenv("REGION")),
		Credentials: credential,
	}

	s, err := session.NewSession(&awsConfig)
	if err != nil {
		log.Println("failed to create S3 session:", err.Error())
		return
	}

	se := Session{s}

	err = se.upload(paths)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
