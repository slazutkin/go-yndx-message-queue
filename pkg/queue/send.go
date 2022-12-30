package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func (q *queue) Send(ctx context.Context, msg string) (msgId string, err error) {
	sent, err := q.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    q.queue.QueueUrl,
		MessageBody: &msg,
	})

	if err != nil {
		return "", err
	}

	return *sent.MessageId, nil
}
