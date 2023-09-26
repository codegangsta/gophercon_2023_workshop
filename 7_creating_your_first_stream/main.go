package main

import (
	"context"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	ctx := context.Background()
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	stream, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:        "my_stream",
		Description: "My stream",
		Subjects:    []string{"my_stream.>"},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Stream: %+v", stream)

	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:    "my_consumer",
		Durable: "my_consumer",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = consumer.Consume(func(msg jetstream.Msg) {
		meta, err := msg.Metadata()
		if err != nil {
			log.Print(err)
			msg.Nak()
			return
		}
		log.Printf("Message: %d %s %s", meta.Sequence.Consumer, msg.Subject(), string(msg.Data()))
		msg.Ack()
	})
	if err != nil {
		log.Fatal(err)
	}

	runtime.Goexit()
}
