package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	home := os.Getenv("HOME")
	s, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
		Endpoint: aws.String("http://localhost:4576"),
		Credentials: credentials.NewSharedCredentials(home + "/.aws/credentials", "localstack"),
	})

	service := sqs.New(s)

	params := &sqs.SendMessageInput{
		QueueUrl: aws.String("http://localhost:4576/queue/TEST"),
		MessageBody: aws.String("testtesttest"),
	}

	req, err := service.SendMessage(params)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(req)
}

