package rabbit

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Noop struct{}

func (n *Noop) Consume(ctx context.Context, consumeFunction func(ctx context.Context, msg amqp.Delivery) error) error {
	return nil
}

func (n *Noop) Publish(ctx context.Context, body []byte) error {
	return nil
}
