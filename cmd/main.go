package main

import (
	"fmt"
	"log"
	"os"
	"playground/pub-sub-aws/internal/publisher"
	"playground/pub-sub-aws/internal/subscriber"
	"playground/pub-sub-aws/pkg/awsclient"
)

func main() {
	snsClient, sqsClient, err := awsclient.NewAWSClients()
	if err != nil {
		log.Fatalf("failed to create AWS Clients: %v", err)
	}

	topicArn := os.Getenv("TOPIC_ARN")
	queueURL := os.Getenv("QUEUE_URL")

	// publish
	messageID, err := publisher.PublishMessage(snsClient, topicArn, "This is just a test")
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	fmt.Printf("published message with ID: %s\n", messageID)

	// receive
	fmt.Println("waiting the messages")
	err = subscriber.ReceiveMessages(sqsClient, queueURL)
	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}
}
