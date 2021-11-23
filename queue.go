package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ListQueues(sess *session.Session) {
	svc := sqs.New(sess)

	if svc == nil {
		fmt.Errorf("SQS Session empty")
		return
	}

	result, err := svc.ListQueues(nil)
	if err != nil {
		fmt.Errorf("Error %v", err)
		return
	}

	for i, url := range result.QueueUrls {
		fmt.Printf("%d: %s\n", i+1, *url)
	}
}

func CreateQueue(sess *session.Session, queue *string) {
	svc := sqs.New(sess)

	if *queue == "" {
		fmt.Println("You must supply a queue name ")
		return
	}

	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: queue,
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("60"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})

	if err != nil {
		fmt.Errorf("Error %v", err)
		return
	}

	fmt.Printf("Queue created successfully %s !!", *result.QueueUrl)
}

func GetQueueURL(sess *session.Session, queue *string) *string {
	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		fmt.Errorf("Error %v", err)
		return nil
	}

	return result.QueueUrl
}

func DeleteQueue(sess *session.Session, queue *string) {
	svc := sqs.New(sess)

	url := GetQueueURL(sess, queue)
	if url == nil {
		fmt.Println("Queue does not exist")
		return
	}

	_, err := svc.DeleteQueue(&sqs.DeleteQueueInput{
		QueueUrl: url,
	})
	if err != nil {
		fmt.Errorf("Error %v", err)
		return
	}

	fmt.Println("Queue deleted successfully !!")
}
