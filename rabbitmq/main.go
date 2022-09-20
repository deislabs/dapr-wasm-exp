package main

import (
	"fmt"
	"github.com/dapr/kit/logger"
	mdata "github.com/dragosv/dapr-wasm/metadata"
	"github.com/dragosv/dapr-wasm/pubsub"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

func newBroker() *rabbitMQInMemoryBroker {
	return &rabbitMQInMemoryBroker{
		buffer: make(chan amqp.Delivery, 2),
	}
}

func main() {
	broker := newBroker()
	pubsubRabbitMQ := newRabbitMQTest(broker)
	metadata := pubsub.Metadata{Base: mdata.Base{
		Properties: map[string]string{
			metadataHostKey:       "anyhost",
			metadataConsumerIDKey: "consumer",
		},
	}}
	err := pubsubRabbitMQ.Init(metadata)
	failIf(err, "unable init rabbitmq")
	failIfNotEqual(1, broker.connectCount)
	failIfNotEqual(0, broker.closeCount)

	topic := "mytopic"

	messageCount := 0
	lastMessage := ""
	processed := make(chan bool)
	handler := func(ctx context.Context, msg *pubsub.NewMessage) error {
		messageCount++
		lastMessage = string(msg.Data)
		processed <- true

		return nil
	}

	err = pubsubRabbitMQ.Subscribe(context.Background(), pubsub.SubscribeRequest{Topic: topic}, handler)
	failIf(err, "unable to subscribe to rabbitmq")

	err = pubsubRabbitMQ.Publish(&pubsub.PublishRequest{Topic: topic, Data: []byte("hello world")})
	failIf(err, "unable to publish to rabbitmq")
	<-processed
	failIfNotEqual(1, messageCount)
	failIfNotEqual("hello world", lastMessage)

	err = pubsubRabbitMQ.Publish(&pubsub.PublishRequest{Topic: topic, Data: []byte("foo bar")})
	//assert.Nil(t, err)
	<-processed
	failIfNotEqual(2, messageCount)
	failIfNotEqual("foo bar", lastMessage)

	fmt.Printf("round trip: %s\n", "test1")
}

func failIf(err error, msg string) {
	if err != nil {
		log.Fatalf("error "+msg+": %v", err)
	}
}

func failIfNotEqual(expected interface{}, actual interface{}) {
	if expected != actual {
		log.Fatalf("assert error")
	}
}

func newRabbitMQTest(broker *rabbitMQInMemoryBroker) pubsub.PubSub {
	return &rabbitMQ{
		declaredExchanges: make(map[string]bool),
		logger:            logger.NewLogger("test"),
		connectionDial: func(host string) (rabbitMQConnectionBroker, rabbitMQChannelBroker, error) {
			broker.connectCount++

			return broker, broker, nil
		},
	}
}

type rabbitMQInMemoryBroker struct {
	buffer chan amqp.Delivery

	connectCount int
	closeCount   int
}

func (r *rabbitMQInMemoryBroker) Qos(prefetchCount, prefetchSize int, global bool) error {
	return nil
}

func (r *rabbitMQInMemoryBroker) Publish(exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error {
	// This is actually how the SDK implements it
	_, err := r.PublishWithDeferredConfirm(exchange, key, mandatory, immediate, msg)
	return err
}

func (r *rabbitMQInMemoryBroker) PublishWithDeferredConfirm(exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) (*amqp.DeferredConfirmation, error) {
	if string(msg.Body) == errorChannelConnection {
		return nil, errors.New(errorChannelConnection)
	}

	r.buffer <- createAMQPMessage(msg.Body)

	return nil, nil
}

func (r *rabbitMQInMemoryBroker) QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (amqp.Queue, error) {
	return amqp.Queue{Name: name}, nil
}

func (r *rabbitMQInMemoryBroker) QueueBind(name string, key string, exchange string, noWait bool, args amqp.Table) error {
	return nil
}

func (r *rabbitMQInMemoryBroker) Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return r.buffer, nil
}

func (r *rabbitMQInMemoryBroker) Nack(tag uint64, multiple bool, requeue bool) error {
	return nil
}

func (r *rabbitMQInMemoryBroker) Ack(tag uint64, multiple bool) error {
	return nil
}

func (r *rabbitMQInMemoryBroker) ExchangeDeclare(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp.Table) error {
	return nil
}

func (r *rabbitMQInMemoryBroker) Confirm(noWait bool) error {
	return nil
}

func (r *rabbitMQInMemoryBroker) Close() error {
	r.closeCount++

	return nil
}

func (r *rabbitMQInMemoryBroker) IsClosed() bool {
	return r.connectCount <= r.closeCount
}
