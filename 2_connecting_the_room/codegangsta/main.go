package main

import (
	"encoding/json"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

type GuestEntry struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

type Message struct {
	Message string `json:"message"`
}

type Error struct {
	Error string `json:"error"`
}

func main() {
	q := "guestbook"

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalln(err)
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalln(err)
	}

	kv, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket:      "gophercon_guestbook",
		Description: "A key value store for the gophercon guestbook app",
		History:     5,
	})
	if err != nil {
		log.Fatalln(err)
	}

	nc.QueueSubscribe("gophercon.guestbook", q, func(msg *nats.Msg) {
		msg.Respond([]byte(`
    Welcome to the gophercon guestbook!

    USAGE:

      # Sign the guestbook
      nats req gophercon.guestbook.sign.{guest_name} "{guest_message}"

      # List entries in the guestbook
      nats req gophercon.guestbook.guests ""

      # View all guestbook entries by a guest
      nats req gophercon.guestbook.guests.{guest_name} ""
    `))
	})

	nc.QueueSubscribe("gophercon.guestbook.sign.*", q, func(msg *nats.Msg) {
		name := strings.Split(msg.Subject, ".")[3]
		_, err := kv.Put(name, msg.Data)
		if err != nil {
			Respond(msg, Error{err.Error()})
			return
		}

		Respond(msg, Message{"Thank you for signing the guest book, " + name})
	})

	nc.QueueSubscribe("gophercon.guestbook.guests", q, func(msg *nats.Msg) {
		keys, err := kv.Keys()
		if err != nil {
			Respond(msg, Error{err.Error()})
			return
		}

		Respond(msg, struct {
			Guests []string `json:"guests"`
		}{
			Guests: keys,
		})
	})

	nc.QueueSubscribe("gophercon.guestbook.guests.*", q, func(msg *nats.Msg) {
		name := strings.Split(msg.Subject, ".")[3]
		history, err := kv.History(name)

		var entries []GuestEntry
		for _, h := range history {
			entries = append(entries, GuestEntry{
				Message: string(h.Value()),
				Time:    h.Created(),
			})
		}

		if err != nil {
			Respond(msg, Error{err.Error()})
			return
		}

		Respond(msg, struct {
			Name    string       `json:"name"`
			Entries []GuestEntry `json:"entries"`
		}{
			Name:    name,
			Entries: entries,
		})
	})

	runtime.Goexit()
}

func Respond(msg *nats.Msg, v any) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		data, _ = json.MarshalIndent(Error{err.Error()}, "", "  ")
	}

	msg.Respond(data)
}
