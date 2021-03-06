package producer

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendMessage(ch chan map[string]string, sess *session.Session, queueURL *string) {
	for i:=0;i<10;i++{
		svc := sqs.New(sess)
		msg := transformToSQSMap(<-ch)

		_, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds:      aws.Int64(10),
			MessageAttributes: msg,
			MessageBody:       aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
			QueueUrl:          queueURL,
		})
		if err != nil {
			fmt.Errorf("Error %v\n", err)
			return
		}

		fmt.Println("Message sent successfully !!")
	}
	close(ch)
}

func transformToSQSMap(msgMap map[string]string) map[string]*sqs.MessageAttributeValue {
	// Create msg to the type accepted by SQS
	msg := make(map[string]*sqs.MessageAttributeValue)
	for k, v := range msgMap {
		msg[k] = &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(v),
		}
	}
	return msg
}
