package main

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	ctx := context.Background()
	nc, err := nats.Connect("nats://demo.nats.io:4222")
	if err != nil {
		log.Fatal(err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	bucket, err := js.KeyValue(ctx, "my_bucket")
	if err != nil {
		log.Fatal(err)
	}

	w, err := bucket.WatchAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for entry := range w.Updates() {
		if entry != nil {
			log.Printf("[%s] %s: %s", entry.Operation().String(), entry.Key(), string(entry.Value()))
		}
	}
}
