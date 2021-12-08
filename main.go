package main

import (
	"fmt"

	"github.com/akshara-nigam/hello-SQS/consumer"
	"github.com/akshara-nigam/hello-SQS/producer"
	"github.com/akshara-nigam/hello-SQS/sqs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
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

	sqs.ListQueues(sess)
	sqs.CreateQueue(sess, &queueName)
	url := sqs.GetQueueURL(sess, &queueName)
	producer.SendMessage(sess, url, msgMap)

	s := consumer.NewSQS()
	s.Session = sess
	s.URL = url
	s.ReceiveMessage()

	sqs.DeleteQueue(sess, &queueName)
	sqs.ListQueues(sess)
}
