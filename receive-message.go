package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ReceiveMessage(sess *session.Session, queueURL *string) {
	svc := sqs.New(sess)

	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(1),
	})
	if err != nil {
		fmt.Errorf("Error %v\n", err)
		return
	}

	if msgResult.Messages == nil {
		fmt.Printf("Queue %v empty\n", *queueURL)
		return
	}

	for _, msg := range msgResult.Messages {
		fmt.Printf("Body : %v\n", *msg.Body)
		fmt.Printf("Message : %v\n", msg.MessageAttributes)

		// Delete the message as soon as it is received from the queue
		DeleteMessage(sess, queueURL, msg.ReceiptHandle)
	}
}

func DeleteMessage(sess *session.Session, queueURL, receipt *string) {
	svc := sqs.New(sess)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: receipt,
	})
	if err != nil {
		fmt.Errorf("Error %v", err)
		return
	}

	fmt.Println("SQS Message deleted successfully !!")
}
