package publisher

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

func PublishMessage(snsClient *sns.SNS, topicArn, message string) (string, error) {
	result, err := snsClient.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(topicArn),
	})
	if err != nil {
		return "", err
	}

	return *result.MessageId, nil
}
