package main

import (
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

var favorites = map[string]string{
	"color":    "gray?",
	"food":     "pizza",
	"season":   "fall",
	"movie":    "Mullholand Drive",
	"beverage": "espresso",
}

func main() {
	nc, err := nats.Connect("nats://demo.nats.io:4222")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to NATS server", nc.ConnectedUrl())

	// create a microservice
	service, err := micro.AddService(nc, micro.Config{
		Name:        "jeremy",
		Version:     "0.0.1",
		Description: "Get jeremy's favorite things",
	})
	if err != nil {
		log.Fatal(err)
	}

	service.AddEndpoint("list_favorites", micro.HandlerFunc(func(r micro.Request) {
		//list keys in favorites map
		keys := make([]string, 0, len(favorites))
		for k := range favorites {
			keys = append(keys, k)
		}
		r.RespondJSON(keys)
	}),
		micro.WithEndpointSubject("jeremy.favorites"),
		micro.WithEndpointMetadata(map[string]string{
			"description": "List all of Jeremy's favorite things",
		}),
	)

	service.AddEndpoint("get_favorite", micro.HandlerFunc(func(r micro.Request) {
		input := string(r.Data())
		favorite, ok := favorites[input]
		if !ok {
			r.Error("not_found", "favorite not found: "+input, nil)
			return
		}
		r.Respond([]byte(favorite))
	}),
		micro.WithEndpointSubject("jeremy.favorite"),
		micro.WithEndpointMetadata(map[string]string{
			"description": "Get one of Jeremy's favorite things",
		}),
	)

	runtime.Goexit()
}
