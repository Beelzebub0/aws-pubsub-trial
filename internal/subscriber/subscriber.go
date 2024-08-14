package subscriber

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ReceiveMessages(sqsClient *sqs.SQS, queueURL string) error {
	for {
		result, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: aws.Int64(1),
			WaitTimeSeconds:     aws.Int64(20),
		})
		if err != nil {
			return err
		}

		for _, message := range result.Messages {
			fmt.Printf("Received message: %s\n", *message.Body)

			_, err := sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: message.ReceiptHandle,
			})
			if err != nil {
				return err
			}
		}
	}
}
