package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/slazutkin/go-yndx-message-queue/pkg/queue"
)

const (
	QUEUE_NAME   = "YNDX_QUEUE_NAME"
	QUEUE_URL    = "YNDX_URL"
	QUEUE_REGION = "YNDX_REGION"
)

func main() {
	queueName := os.Getenv(QUEUE_NAME)
	queueURL := os.Getenv(QUEUE_URL)
	queueRegion := os.Getenv(QUEUE_REGION)

	q, err := queue.New(queueName, queueURL, queueRegion)

	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	id, err := q.Send(ctx, "Hey, I am a message!")

	if err != nil {
		log.Fatalln("Failed to send a message")
	}

	log.Printf("Message is sent: %s", id)

	msg, err := q.Receive(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for _, m := range msg {
		fmt.Printf("Received: %s - %s\n", m.ID, m.Body)
	}
}
