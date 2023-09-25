package main

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/nats-io/nats.go"
)

var favorites = map[string]string{
	"color":   "gray?",
	"food":    "pizza",
	"season":  "fall",
	"movie":   "The Shining",
	"alcohol": "Gin",
}

func main() {
	natsUrl := os.Getenv("NATS_URL")
	if natsUrl == "" {
		natsUrl = "nats://demo.nats.io:4222"
	}

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatalln(err)
	}

	nc.QueueSubscribe("gophercon.services", "jeremy", func(msg *nats.Msg) {
		msg.Respond([]byte("gophercon.services.jeremy.favorites"))
	})

	nc.QueueSubscribe("gophercon.services.jeremy.favorites", "jeremy", func(msg *nats.Msg) {
		val, ok := favorites[string(msg.Data)]
		if !ok {
			keys := []string{}
			for k := range favorites {
				keys = append(keys, k)
			}
			msg.Respond([]byte("Hm... I don't recognize that. Try asking me about one of these: " + strings.Join(keys, ",")))
			return
		}

		msg.Respond([]byte(val))
	})

	runtime.Goexit()
}
