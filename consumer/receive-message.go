package consumer

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQS struct {
	Session *session.Session
	URL *string
}

// ConsumeMessage contains the message body
type ConsumerMessage struct {
	Body    string            `json:"body"`
	Message map[string]string `json:"message"`
}

func NewSQS() *SQS {
	return &SQS{
		Session: nil,
		URL: nil,
	}
}

func (s *SQS) ReceiveMessage() {
	chnMessages := make(chan *sqs.Message, 1)
	go s.pollMessages(chnMessages)

	for msg := range chnMessages {
		fmt.Println("\nRECEIVED MESSAGE >>> ")

		var message ConsumerMessage
		message.Body = *msg.Body
		message.Message = make(map[string]string)

		for key, atr := range msg.MessageAttributes {
			message.Message[key] = *atr.StringValue
		}
		fmt.Println(message)

		// Delete the message as soon as it is received from the queue
		s.deleteMessage(msg.ReceiptHandle)
	}
}

// pollMessages starts polling for messages in the queue
func (s *SQS) pollMessages(chn chan<- *sqs.Message) {
	svc := sqs.New(s.Session)

	for {
		fmt.Println("\n\nPolling for Messages")

		msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:              s.URL,
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
			fmt.Printf("Queue %v empty\n", *s.URL)
			continue
		}

		for _, message := range msgResult.Messages {
			chn <- message
		}
	}
}

func (s *SQS) deleteMessage(receipt *string) {
	svc := sqs.New(s.Session)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      s.URL,
		ReceiptHandle: receipt,
	})
	if err != nil {
		fmt.Errorf("Error %v", err)
		return
	}

	fmt.Println("SQS Message deleted successfully !!")
}
