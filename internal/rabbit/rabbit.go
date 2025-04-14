package rabbit

import (
	"context"
	cfg "ekko/pkg/config"
	"ekko/pkg/logger/pkg/logging"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit interface {
	Consume(ctx context.Context, consumeFunction func(ctx context.Context, msg amqp.Delivery) error) error
	Publish(ctx context.Context, body []byte) error
}

type rabbit struct {
	connectionUrl string
	comsumeQueue  string
	publicQueue   string
	max_comsumer  int32
	expireTime    int32
}

func New(cfg *cfg.RabbitMQ) Rabbit {
	if cfg == nil {
		return &Noop{}
	}

	connectionUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.Username, cfg.Password, cfg.Address, cfg.Port)
	return &rabbit{
		connectionUrl: connectionUrl,
		comsumeQueue:  cfg.ConsumeQueue,
		publicQueue:   cfg.PublicQueue,
		max_comsumer:  cfg.MaxConsumer,
		expireTime:    cfg.ExpireTime,
	}
}

func (r *rabbit) processMessage(ctx context.Context, msg amqp.Delivery, sem chan struct{}, consumeFunction func(ctx context.Context, msg amqp.Delivery) error) {
	logging.Logger(ctx).Info(fmt.Sprintf("Received: %s", msg.Body))
	defer func() { <-sem }()

	if err := consumeFunction(ctx, msg); err != nil {
		logging.Logger(ctx).Error(fmt.Sprintf("Error: %s", err.Error()))
		msg.Nack(false, true)
	} else {
		logging.Logger(ctx).Info(fmt.Sprintf("Acknowledge: %s", msg.Body))
		msg.Ack(false)
	}
}

func (r *rabbit) Consume(ctx context.Context, consumeFunction func(ctx context.Context, msg amqp.Delivery) error) error {

	conn, err := amqp.Dial(r.connectionUrl)
	if err != nil {
		return err
	}
	defer conn.Close()

	logging.Logger(ctx).Info("Connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(r.comsumeQueue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	sem := make(chan struct{}, r.max_comsumer)

	for msg := range msgs {
		sem <- struct{}{}
		go r.processMessage(ctx, msg, sem, consumeFunction)
	}

	return nil
}

func (r *rabbit) Publish(ctx context.Context, body []byte) error {
	conn, err := amqp.Dial(r.connectionUrl)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(r.publicQueue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
		Expiration:  fmt.Sprintf("%d", r.expireTime),
	})
	if err != nil {
		return err
	}

	logging.Logger(ctx).Info(fmt.Sprintf("Sent: %s", string(body)))
	return nil
}
