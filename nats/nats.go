package nats

import (
	"context"
	"errors"

	"github.com/andiksetyawan/log"
	"github.com/andiksetyawan/messaging"
	natslib "github.com/nats-io/nats.go"
)

type nats struct {
	conn   *natslib.Conn
	logger log.Logger
}

func NewNATSMessaging(natsURL string, logger log.Logger) (messaging.Messaging, error) {
	conn, err := natslib.Connect(natsURL)
	if err != nil {
		return nil, err
	}

	return &nats{
		conn:   conn,
		logger: logger,
	}, nil
}

func (n *nats) Publish(ctx context.Context, topic string, message []byte) error {
	return n.conn.Publish(topic, message)
}

func (n *nats) Subscribe(ctx context.Context, topic string, handler func(ctx context.Context, message []byte) error) error {
	_, err := n.conn.Subscribe(topic, func(msg *natslib.Msg) {
		if err := handler(ctx, msg.Data); err == nil {
			if err := msg.Ack(); err != nil {
				n.logger.Error(ctx, "failed to Ack the message", "error", err)
				return
			}
		}
	})
	if err != nil {
		return err
	}

	return nil
}

func (n *nats) Unsubscribe(ctx context.Context, topic string) error {
	return errors.New("not implemented")
}

func (n *nats) Requeue(ctx context.Context, topic string, message []byte) error {
	return n.Publish(ctx, topic, message)
}

func (n *nats) ListTopics(ctx context.Context) ([]string, error) {
	return nil, errors.New("not implemented")
}

func (n *nats) HealthCheck(ctx context.Context) error {
	if n.conn.IsClosed() {
		return errors.New("connection is closed")
	}
	return nil
}

func (n *nats) Flush(ctx context.Context) error {
	return n.conn.Flush()
}

func (n *nats) Close(ctx context.Context) error {
	err := n.Flush(ctx)
	if err != nil {
		return err
	}
	n.conn.Close()
	return nil
}
