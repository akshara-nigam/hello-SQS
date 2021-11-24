package consumer

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ReceiveMessage(sess *session.Session, queueURL *string) {
	svc := sqs.New(sess)

	chnMessages := make(chan *sqs.Message, 1)
	go pollMessages(svc, queueURL, chnMessages)

	for msg := range chnMessages {
		fmt.Println("\nRECEIVED MESSAGE >>> ")
		fmt.Printf("Body : %v\n", *msg.Body)
		fmt.Printf("Message : %v\n", msg.MessageAttributes)

		// Delete the message as soon as it is received from the queue
		deleteMessage(sess, queueURL, msg.ReceiptHandle)
	}
}

// pollMessages starts polling for messages in the queue
func pollMessages(svc *sqs.SQS, queueURL *string, chn chan<- *sqs.Message) {
	for {
		fmt.Println("\n\nPolling for Messages")

		msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:              queueURL,
			MaxNumberOfMessages:   aws.Int64(1),
			WaitTimeSeconds:       aws.Int64(15),
			AttributeNames:        []*string{aws.String(sqs.MessageSystemAttributeNameSentTimestamp)},
			MessageAttributeNames: []*string{aws.String(sqs.QueueAttributeNameAll)},
		})
		if err != nil {
			fmt.Errorf("Error %v\n", err)
			continue
		}

		if msgResult.Messages == nil {
			fmt.Printf("Queue %v empty\n", *queueURL)
			continue
		}

		for _, message := range msgResult.Messages {
			chn <- message
		}
	}
}

func deleteMessage(sess *session.Session, queueURL, receipt *string) {
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
