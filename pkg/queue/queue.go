package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Queue interface {
	Send(ctx context.Context, msg string) (msgId string, err error)
	Receive(ctx context.Context) ([]Message, error)
}

type queue struct {
	client *sqs.Client
	queue  *sqs.GetQueueUrlOutput
}

func New(queueName, queueUrl, queueRegion string) (Queue, error) {
	ctx := context.Background()

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           queueUrl,
			SigningRegion: queueRegion,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	q, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		return nil, err
	}

	return &queue{client: client, queue: q}, nil
}
