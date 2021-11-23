package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)


func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "devmfa"),
	})

	if err != nil {
		fmt.Errorf("%v\n", err)
	}

	queueName := "Test-Queue-Akshara"


	ListQueues(sess)
	CreateQueue(sess, &queueName)
	url := GetQueueURL(sess, &queueName)
	SendMessage(sess, url)
	ReceiveMessage(sess, url)
	DeleteQueue(sess, &queueName)
	ListQueues(sess)
}
