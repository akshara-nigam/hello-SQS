package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "devmfa"),
	})

	if err != nil {
		fmt.Errorf("%v\n", err)
	}

	queueName := "Test-Queue-Akshara"
	msgMap := make(map[string]string)
	msgMap["Author"] = "John Grisham"
	msgMap["Title"] = "The Whistler"
	msgMap["WeeksOn"] = "6"

	ListQueues(sess)
	CreateQueue(sess, &queueName)
	url := GetQueueURL(sess, &queueName)
	SendMessage(sess, url, msgMap)
	ReceiveMessage(sess, url)
	DeleteQueue(sess, &queueName)
	ListQueues(sess)
}
