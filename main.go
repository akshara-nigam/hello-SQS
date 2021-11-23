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
		fmt.Errorf("%v", err)
	}

	queueName := "Test-Queue-Akshara"

	DeleteQueue(sess, &queueName)
	//ListQueues(sess)
	//CreateQueue(&queueName)
	//url := GetQueueURL(&queueName)

	//SendMessage(url)
	//ReceiveMessage(url)
}
