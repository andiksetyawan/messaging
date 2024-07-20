package messaging

import (
	"context"
)

type Messaging interface {
	Publish(ctx context.Context, topic string, message []byte) error
	Subscribe(ctx context.Context, topic string, handler func(ctx context.Context, message []byte) error) error
	Unsubscribe(ctx context.Context, topic string) error
	Requeue(ctx context.Context, topic string, message []byte) error
	ListTopics(ctx context.Context) ([]string, error)
	HealthCheck(ctx context.Context) error
	Flush(ctx context.Context) error
	Close(ctx context.Context) error
}
