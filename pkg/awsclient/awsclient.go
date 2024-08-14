package awsclient

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func NewAWSClients() (*sns.SNS, *sqs.SQS, error) {

	sess, err := session.NewSession(&aws.Config{
		Region:                        aws.String("ap-southeast-2"),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})
	if err != nil {
		return nil, nil, err
	}

	snsClient := sns.New(sess)
	sqsClient := sqs.New(sess)

	return snsClient, sqsClient, nil
}
