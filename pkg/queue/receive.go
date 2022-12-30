package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func (q *queue) Receive(ctx context.Context) ([]Message, error) {

	received, err := q.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl: q.queue.QueueUrl,
	})

	if err != nil {
		return nil, err
	}

	var msg []Message

	for _, m := range received.Messages {
		if _, err := q.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
			QueueUrl:      q.queue.QueueUrl,
			ReceiptHandle: m.ReceiptHandle,
		}); err != nil {
			return nil, err
		}

		msg = append(msg, Message{
			ID:   *m.MessageId,
			Body: *m.Body,
		})
	}

	return msg, nil
}
